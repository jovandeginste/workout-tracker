package slognil

import (
	"context"
	"log/slog"
)

type Handler struct{}

var _ slog.Handler = Handler{}

func (Handler) Enabled(context.Context, slog.Level) bool {
	return false
}

func (Handler) Handle(context.Context, slog.Record) error {
	return nil
}

func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h Handler) WithGroup(name string) slog.Handler {
	return h
}

// NewLogger returns a logger configured to use the nil handler.
func NewLogger() *slog.Logger {
	return slog.New(Handler{})
}
