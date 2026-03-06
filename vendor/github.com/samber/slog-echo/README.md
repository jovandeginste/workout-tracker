
# slog: Echo middleware

[![tag](https://img.shields.io/github/tag/samber/slog-echo.svg)](https://github.com/samber/slog-echo/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/slog-echo?status.svg)](https://pkg.go.dev/github.com/samber/slog-echo)
![Build Status](https://github.com/samber/slog-echo/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/slog-echo)](https://goreportcard.com/report/github.com/samber/slog-echo)
[![Coverage](https://img.shields.io/codecov/c/github/samber/slog-echo)](https://codecov.io/gh/samber/slog-echo)
[![Contributors](https://img.shields.io/github/contributors/samber/slog-echo)](https://github.com/samber/slog-echo/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/slog-echo)](./LICENSE)

[Echo](https://github.com/labstack/echo) middleware to log http requests using [slog](https://pkg.go.dev/log/slog).

<div align="center">
  <hr>
  <sup><b>Sponsored by:</b></sup>
  <br>
  <a href="https://cast.ai/samuel">
    <div>
      <img src="https://github.com/user-attachments/assets/502f8fa8-e7e8-4754-a51f-036d0443e694" width="200" alt="Cast AI">
    </div>
    <div>
      Cut Kubernetes & AI costs, boost application stability
    </div>
  </a>
  <br>
  <a href="https://www.dash0.com?utm_campaign=148395251-samber%20github%20sponsorship&utm_source=github&utm_medium=sponsorship&utm_content=samber">
    <div>
      <img src="https://github.com/user-attachments/assets/b1f2e876-0954-4dc3-824d-935d29ba8f3f" width="200" alt="Dash0">
    </div>
    <div>
      100% OpenTelemetry-native observability platform<br>Simple to use, built on open standards, and designed for full cost control
    </div>
  </a>
  <hr>
</div>

**See also:**

- [slog-multi](https://github.com/samber/slog-multi): `slog.Handler` chaining, fanout, routing, failover, load balancing...
- [slog-formatter](https://github.com/samber/slog-formatter): `slog` attribute formatting
- [slog-sampling](https://github.com/samber/slog-sampling): `slog` sampling policy
- [slog-mock](https://github.com/samber/slog-mock): `slog.Handler` for test purposes

**HTTP middlewares:**

- [slog-gin](https://github.com/samber/slog-gin): Gin middleware for `slog` logger
- [slog-echo](https://github.com/samber/slog-echo): Echo middleware for `slog` logger
- [slog-fiber](https://github.com/samber/slog-fiber): Fiber middleware for `slog` logger
- [slog-chi](https://github.com/samber/slog-chi): Chi middleware for `slog` logger
- [slog-http](https://github.com/samber/slog-http): `net/http` middleware for `slog` logger

**Loggers:**

- [slog-zap](https://github.com/samber/slog-zap): A `slog` handler for `Zap`
- [slog-zerolog](https://github.com/samber/slog-zerolog): A `slog` handler for `Zerolog`
- [slog-logrus](https://github.com/samber/slog-logrus): A `slog` handler for `Logrus`

**Log sinks:**

- [slog-datadog](https://github.com/samber/slog-datadog): A `slog` handler for `Datadog`
- [slog-betterstack](https://github.com/samber/slog-betterstack): A `slog` handler for `Betterstack`
- [slog-rollbar](https://github.com/samber/slog-rollbar): A `slog` handler for `Rollbar`
- [slog-loki](https://github.com/samber/slog-loki): A `slog` handler for `Loki`
- [slog-sentry](https://github.com/samber/slog-sentry): A `slog` handler for `Sentry`
- [slog-syslog](https://github.com/samber/slog-syslog): A `slog` handler for `Syslog`
- [slog-logstash](https://github.com/samber/slog-logstash): A `slog` handler for `Logstash`
- [slog-fluentd](https://github.com/samber/slog-fluentd): A `slog` handler for `Fluentd`
- [slog-graylog](https://github.com/samber/slog-graylog): A `slog` handler for `Graylog`
- [slog-quickwit](https://github.com/samber/slog-quickwit): A `slog` handler for `Quickwit`
- [slog-slack](https://github.com/samber/slog-slack): A `slog` handler for `Slack`
- [slog-telegram](https://github.com/samber/slog-telegram): A `slog` handler for `Telegram`
- [slog-mattermost](https://github.com/samber/slog-mattermost): A `slog` handler for `Mattermost`
- [slog-microsoft-teams](https://github.com/samber/slog-microsoft-teams): A `slog` handler for `Microsoft Teams`
- [slog-webhook](https://github.com/samber/slog-webhook): A `slog` handler for `Webhook`
- [slog-kafka](https://github.com/samber/slog-kafka): A `slog` handler for `Kafka`
- [slog-nats](https://github.com/samber/slog-nats): A `slog` handler for `NATS`
- [slog-parquet](https://github.com/samber/slog-parquet): A `slog` handler for `Parquet` + `Object Storage`
- [slog-channel](https://github.com/samber/slog-channel): A `slog` handler for Go channels

## 🚀 Install

```sh
# echo v4 (current)
go get github.com/samber/slog-echo

# echo v5 (alpha)
go get github.com/samber/slog-echo@echo-v5
```

**Compatibility**: go >= 1.21

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Usage

### Handler options

```go
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
```

Attributes will be injected in log payload.

Other global parameters:

```go
slogecho.TraceIDKey = "trace_id"
slogecho.SpanIDKey = "span_id"
slogecho.RequestBodyMaxSize  = 64 * 1024 // 64KB
slogecho.ResponseBodyMaxSize = 64 * 1024 // 64KB
slogecho.HiddenRequestHeaders = map[string]struct{}{ ... }
slogecho.HiddenResponseHeaders = map[string]struct{}{ ... }
```

### Minimal

```go
import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"log/slog"
)

// Create a slog logger, which:
//   - Logs to stdout.
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// Echo instance
e := echo.New()

// Middleware
e.Use(slogecho.New(logger))
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
})
e.GET("/error", func(c echo.Context) error {
	return echo.
		NewHTTPError(http.StatusInternalServerError, "I'm angry").
		WithInternal(errors.New("I'm angry internally"))
})

// Start server
e.Logger.Fatal(e.Start(":4242"))

// output:
// time=2023-10-15T20:32:58.926+02:00 level=INFO msg="Success" env=production request.time=2023-10-15T20:32:58.626+02:00 request.method=GET request.path=/ request.route="" request.ip=127.0.0.1:63932 request.length=0 response.time=2023-10-15T20:32:58.926+02:00 response.latency=100ms response.status=200 response.length=7 id=229c7fc8-64f5-4467-bc4a-940700503b0d  http.error="map[code:500 internal:I'm angry internally message:I'm angry]" http.internal="I'm angry internally"
```

### OTEL

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

config := slogecho.Config{
	WithSpanID:  true,
	WithTraceID: true,
}

e := echo.New()
e.Use(slogecho.NewWithConfig(logger, config))
e.Use(middleware.Recover())
```

### Custom log levels

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

config := slogecho.Config{
	DefaultLevel:     slog.LevelInfo,
	ClientErrorLevel: slog.LevelWarn,
	ServerErrorLevel: slog.LevelError,
}

e := echo.New()
e.Use(slogecho.NewWithConfig(logger, config))
e.Use(middleware.Recover())
```

### Verbose

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

config := slogecho.Config{
	WithRequestBody: true,
	WithResponseBody: true,
	WithRequestHeader: true,
	WithResponseHeader: true,
}

e := echo.New()
e.Use(slogecho.NewWithConfig(logger, config))
e.Use(middleware.Recover())
```

### Filters

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

e := echo.New()
e.Use(
	slogecho.NewWithFilters(
		logger,
		slogecho.Accept(func (c echo.Context) bool {
			return xxx
		}),
		slogecho.IgnoreStatus(401, 404),
	),
)
e.Use(middleware.Recover())
```

Available filters:

### Filters

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

router := gin.New()
router.Use(
	sloggin.NewWithFilters(
		logger,
		sloggin.Accept(func (c *gin.Context) bool {
			return xxx
		}),
		sloggin.IgnoreStatus(401, 404),
	),
)
```

Available filters:
- Accept / Ignore
- AcceptMethod / IgnoreMethod
- AcceptStatus / IgnoreStatus
- AcceptStatusGreaterThan / IgnoreStatusGreaterThan
- AcceptStatusLessThan / IgnoreStatusLessThan
- AcceptStatusGreaterThanOrEqual / IgnoreStatusGreaterThanOrEqual
- AcceptStatusLessThanOrEqual / IgnoreStatusLessThanOrEqual
- AcceptPath / IgnorePath
- AcceptPathContains / IgnorePathContains
- AcceptPathPrefix / IgnorePathPrefix
- AcceptPathSuffix / IgnorePathSuffix
- AcceptPathMatch / IgnorePathMatch
- AcceptHost / IgnoreHost
- AcceptHostContains / IgnoreHostContains
- AcceptHostPrefix / IgnoreHostPrefix
- AcceptHostSuffix / IgnoreHostSuffix
- AcceptHostMatch / IgnoreHostMatch

### Using custom time formatters

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	slogformatter "github.com/samber/slog-formatter"
	"log/slog"
)

// Create a slog logger, which:
//   - Logs to stdout.
//   - RFC3339 with UTC time format.
logger := slog.New(
	slogformatter.NewFormatterHandler(
		slogformatter.TimezoneConverter(time.UTC),
		slogformatter.TimeFormatter(time.DateTime, nil),
	)(
		slog.NewTextHandler(os.Stdout, nil),
	),
)

// Echo instance
e := echo.New()

// Middleware
e.Use(slogecho.New(logger))
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
})
e.GET("/error", func(c echo.Context) error {
	return echo.
		NewHTTPError(http.StatusInternalServerError, "I'm angry").
		WithInternal(errors.New("I'm angry internally"))
})

// Start server
e.Logger.Fatal(e.Start(":4242"))

// output:
// time=2023-10-15T20:32:58.926+02:00 level=INFO msg="Success" env=production request.time=2023-10-15T20:32:58Z request.method=GET request.path=/ request.route="" request.ip=127.0.0.1:63932 request.length=0 response.time=2023-10-15T20:32:58Z response.latency=100ms response.status=200 response.length=7 id=229c7fc8-64f5-4467-bc4a-940700503b0d error="map[code:500 internal:I'm angry internally message:I'm angry]" internal="I'm angry internally"
```

### Using custom logger sub-group

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// Echo instance
e := echo.New()

// Middleware
e.Use(slogecho.New(logger.WithGroup("http")))
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
})
e.GET("/error", func(c echo.Context) error {
	return echo.
		NewHTTPError(http.StatusInternalServerError, "I'm angry").
		WithInternal(errors.New("I'm angry internally"))
})

// Start server
e.Logger.Fatal(e.Start(":4242"))

// output:
// time=2023-10-15T20:32:58.926+02:00 level=INFO msg="Success" env=production http.request.time=2023-10-15T20:32:58.626+02:00 http.request.method=GET http.request.path=/ http.request.route="" http.request.ip=127.0.0.1:63932 http.request.length=0 http.response.time=2023-10-15T20:32:58.926+02:00 http.response.latency=100ms http.response.status=200 http.response.length=7 http.id=229c7fc8-64f5-4467-bc4a-940700503b0d http.error="map[code:500 internal:I'm angry internally message:I'm angry]" http.internal="I'm angry internally"
```

### Add logger to a single route

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// Echo instance
e := echo.New()

// Middleware
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}, slogecho.New(logger))

// Start server
e.Logger.Fatal(e.Start(":4242"))
```

### Adding custom attributes

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// Add an attribute to all log entries made through this logger.
logger = logger.With("env", "production")

// Echo instance
e := echo.New()

// Middleware
e.Use(slogecho.New(logger))
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	// Add an attribute to a single log entry.
	slogecho.AddCustomAttributes(c, slog.String("foo", "bar"))
	return c.String(http.StatusOK, "Hello, World!")
})

// Start server
e.Logger.Fatal(e.Start(":4242"))

// output:
// time=2023-10-15T20:32:58.926+02:00 level=INFO msg="Success" env=production request.time=2023-10-15T20:32:58.626+02:00 request.method=GET request.path=/ request.route="" request.ip=127.0.0.1:63932 request.length=0 response.time=2023-10-15T20:32:58.926+02:00 response.latency=100ms response.status=200 response.length=7 id=229c7fc8-64f5-4467-bc4a-940700503b0d foo=bar error="map[code:500 internal:I'm angry internally message:I'm angry]" internal="I'm angry internally"
```

### JSON output

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

// Echo instance
e := echo.New()

// Middleware
e.Use(slogecho.New(logger))
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
})

// Start server
e.Logger.Fatal(e.Start(":4242"))

// output:
// {"time":"2023-10-15T20:32:58.926+02:00","level":"INFO","msg":"Success","env":"production","http":{"request":{"time":"2023-10-15T20:32:58.626+02:00","method":"GET","path":"/","route":"","ip":"127.0.0.1:55296","length":0},"response":{"time":"2023-10-15T20:32:58.926+02:00","latency":100000,"status":200,"length":7},"id":"04201917-d7ba-4b20-a3bb-2fffba5f2bd9"}, "error": {"code":500, "internal":"I'm angry internally", "message":"I'm angry"}, "internal": "I'm angry internally"}
```

## 🤝 Contributing

- Ping me on Twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/slog-echo)
- Fix [open issues](https://github.com/samber/slog-echo/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/slog-echo)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2023 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
