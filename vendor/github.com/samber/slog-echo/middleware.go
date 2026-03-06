package slogecho

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/trace"
)

const (
	customAttributesCtxKey = "slog-echo.custom-attributes"
)

var (
	TraceIDKey   = "trace_id"
	SpanIDKey    = "span_id"
	RequestIDKey = "id"

	RequestBodyMaxSize  = 64 * 1024 // 64KB
	ResponseBodyMaxSize = 64 * 1024 // 64KB

	HiddenRequestHeaders = map[string]struct{}{
		"authorization": {},
		"cookie":        {},
		"set-cookie":    {},
		"x-auth-token":  {},
		"x-csrf-token":  {},
		"x-xsrf-token":  {},
	}
	HiddenResponseHeaders = map[string]struct{}{
		"set-cookie": {},
	}
)

type Config struct {
	DefaultLevel     slog.Level
	ClientErrorLevel slog.Level
	ServerErrorLevel slog.Level

	WithUserAgent      bool
	WithRequestID      bool
	WithRequestBody    bool
	WithRequestHeader  bool
	WithResponseBody   bool
	WithResponseHeader bool
	WithSpanID         bool
	WithTraceID        bool
	WithClientIP       bool
	WithCustomMessage  func(c echo.Context, err error) string

	Filters []Filter
}

// New returns a echo.MiddlewareFunc (middleware) that logs requests using slog.
//
// Requests with errors are logged using slog.Error().
// Requests without errors are logged using slog.Info().
func New(logger *slog.Logger) echo.MiddlewareFunc {
	return NewWithConfig(logger, DefaultConfig())
}

// NewWithFilters returns a echo.MiddlewareFunc (middleware) that logs requests using slog.
//
// Requests with errors are logged using slog.Error().
// Requests without errors are logged using slog.Info().
func NewWithFilters(logger *slog.Logger, filters ...Filter) echo.MiddlewareFunc {
	config := DefaultConfig()
	config.Filters = filters
	return NewWithConfig(logger, config)
}

// DefaultConfig returns the default configuration for the request logger.
func DefaultConfig() Config {
	return Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		WithUserAgent:      false,
		WithRequestID:      true,
		WithRequestBody:    false,
		WithRequestHeader:  false,
		WithResponseBody:   false,
		WithResponseHeader: false,
		WithSpanID:         false,
		WithTraceID:        false,
		WithClientIP:       true,
		WithCustomMessage:  nil,

		Filters: []Filter{},
	}
}

// NewWithConfig returns a echo.HandlerFunc (middleware) that logs requests using slog.
func NewWithConfig(logger *slog.Logger, config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			path := req.URL.Path
			query := req.URL.RawQuery

			params := map[string]string{}
			for i, k := range c.ParamNames() {
				params[k] = c.ParamValues()[i]
			}

			// dump request body
			br := newBodyReader(req.Body, RequestBodyMaxSize, config.WithRequestBody)
			req.Body = br

			// dump response body
			bw := newBodyWriter(res.Writer, ResponseBodyMaxSize, config.WithResponseBody)
			res.Writer = bw

			err = next(c)

			if err != nil {
				if _, ok := err.(*echo.HTTPError); !ok {
					err = echo.
						NewHTTPError(http.StatusInternalServerError).
						WithInternal(err)
					c.Error(err)
				}
			}

			// Pass thru filters and skip early the code below, to prevent unnecessary processing.
			for _, filter := range config.Filters {
				if !filter(c) {
					return
				}
			}

			status := res.Status
			method := req.Method
			host := req.Host
			route := c.Path()
			end := time.Now()
			latency := end.Sub(start)
			userAgent := req.UserAgent()
			ip := c.RealIP()
			referer := c.Request().Referer()

			errMsg := err

			var httpErr *echo.HTTPError
			if err != nil && errors.As(err, &httpErr) {
				status = httpErr.Code
				if msg, ok := httpErr.Message.(string); ok {
					errMsg = errors.New(msg)
				}
			}

			baseAttributes := make([]slog.Attr, 0, 3)
			requestAttributes := make([]slog.Attr, 0, 14)
			responseAttributes := make([]slog.Attr, 0, 6)

			requestAttributes = append(requestAttributes,
				slog.Time("time", start.UTC()),
				slog.String("method", method),
				slog.String("host", host),
				slog.String("path", path),
				slog.String("query", query),
				slog.Any("params", params),
				slog.String("route", route),
				slog.String("referer", referer),
			)

			if config.WithClientIP {
				requestAttributes = append(requestAttributes,
					slog.String("ip", ip),
				)
			}

			responseAttributes = append(responseAttributes,
				slog.Time("time", end.UTC()),
				slog.Duration("latency", latency),
				slog.Int("status", status),
			)

			if config.WithRequestID {
				requestID := req.Header.Get(echo.HeaderXRequestID)
				if requestID == "" {
					requestID = res.Header().Get(echo.HeaderXRequestID)
				}
				if requestID != "" {
					baseAttributes = append(baseAttributes, slog.String(RequestIDKey, requestID))
				}
			}

			// otel
			baseAttributes = append(baseAttributes, extractTraceSpanID(c.Request().Context(), config.WithTraceID, config.WithSpanID)...)

			// request body
			requestAttributes = append(requestAttributes, slog.Int("length", br.bytes))
			if config.WithRequestBody {
				requestAttributes = append(requestAttributes, slog.String("body", br.body.String()))
			}

			// request headers
			if config.WithRequestHeader {
				kv := []any{}

				for k, v := range c.Request().Header {
					if _, found := HiddenRequestHeaders[strings.ToLower(k)]; found {
						continue
					}
					kv = append(kv, slog.Any(k, v))
				}

				requestAttributes = append(requestAttributes, slog.Group("header", kv...))
			}

			if config.WithUserAgent {
				requestAttributes = append(requestAttributes, slog.String("user-agent", userAgent))
			}

			xForwardedFor, ok := c.Get(echo.HeaderXForwardedFor).(string)
			if ok && len(xForwardedFor) > 0 {
				ips := lo.Map(strings.Split(xForwardedFor, ","), func(ip string, _ int) string {
					return strings.TrimSpace(ip)
				})
				requestAttributes = append(requestAttributes, slog.Any("x-forwarded-for", ips))
			}

			// response body
			responseAttributes = append(responseAttributes, slog.Int("length", bw.bytes))
			if config.WithResponseBody {
				responseAttributes = append(responseAttributes, slog.String("body", bw.body.String()))
			}

			// response headers
			if config.WithResponseHeader {
				kv := []any{}

				for k, v := range c.Response().Header() {
					if _, found := HiddenResponseHeaders[strings.ToLower(k)]; found {
						continue
					}
					kv = append(kv, slog.Any(k, v))
				}

				responseAttributes = append(responseAttributes, slog.Group("header", kv...))
			}

			attributes := append(
				[]slog.Attr{
					{
						Key:   "request",
						Value: slog.GroupValue(requestAttributes...),
					},
					{
						Key:   "response",
						Value: slog.GroupValue(responseAttributes...),
					},
				},
				baseAttributes...,
			)

			// custom context values
			if v := c.Get(customAttributesCtxKey); v != nil {
				switch attrs := v.(type) {
				case []slog.Attr:
					attributes = append(attributes, attrs...)
				}
			}

			level := config.DefaultLevel
			msg := "Incoming request"

			if status >= http.StatusInternalServerError {
				level = config.ServerErrorLevel
				if err != nil {
					msg = errMsg.Error()
				} else {
					msg = http.StatusText(status)
				}
			} else if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
				level = config.ClientErrorLevel
				if err != nil {
					msg = errMsg.Error()
				} else {
					msg = http.StatusText(status)
				}
			}

			if httpErr != nil {
				attributes = append(
					attributes,
					slog.Any("error", map[string]any{
						"code":     httpErr.Code,
						"message":  httpErr.Message,
						"internal": httpErr.Internal,
					}),
				)

				if httpErr.Internal != nil {
					attributes = append(attributes, slog.String("internal", httpErr.Internal.Error()))
				}
			}

			if config.WithCustomMessage != nil {
				msg = config.WithCustomMessage(c, errMsg)
			}

			logger.LogAttrs(c.Request().Context(), level, msg, attributes...)

			return
		}
	}
}

// AddCustomAttributes adds custom attributes to the request context.
func AddCustomAttributes(c echo.Context, attrs ...slog.Attr) {
	v := c.Get(customAttributesCtxKey)
	if v == nil {
		c.Set(customAttributesCtxKey, attrs)
		return
	}

	switch vAttrs := v.(type) {
	case []slog.Attr:
		c.Set(customAttributesCtxKey, append(vAttrs, attrs...))
	}
}

func extractTraceSpanID(ctx context.Context, withTraceID bool, withSpanID bool) []slog.Attr {
	if !withTraceID && !withSpanID {
		return []slog.Attr{}
	}

	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return []slog.Attr{}
	}

	attrs := make([]slog.Attr, 0, 2)
	spanCtx := span.SpanContext()

	if withTraceID && spanCtx.HasTraceID() {
		traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		attrs = append(attrs, slog.String(TraceIDKey, traceID))
	}

	if withSpanID && spanCtx.HasSpanID() {
		spanID := spanCtx.SpanID().String()
		attrs = append(attrs, slog.String(SpanIDKey, spanID))
	}

	return attrs
}
