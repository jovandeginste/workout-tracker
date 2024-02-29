package fit

import (
	"log"
	"os"
)

type decodeOptions struct {
	logger          Logger
	unknownFields   bool
	unknownMessages bool
}

// DecodeOption configures a decoder.
type DecodeOption func(*decodeOptions)

// WithLogger configures the decoder to enable debug logging using the provided
// logger.
func WithLogger(logger Logger) DecodeOption {
	return func(o *decodeOptions) {
		o.logger = logger
	}
}

// WithStdLogger configures the decoder to enable debug logging using the
// standard library's logger.
func WithStdLogger() DecodeOption {
	return func(o *decodeOptions) {
		o.logger = log.New(os.Stderr, "", 0)
	}
}

// WithUnknownFields configures the decoder to record information about unknown
// fields encountered when decoding a known message type. Currently message
// number, field number and number of occurrences are recorded.
func WithUnknownFields() DecodeOption {
	return func(o *decodeOptions) {
		o.unknownFields = true
	}
}

// WithUnknownMessages configures the decoder to record information about unknown
// messages encountered during decoding of a FIT file. Currently message
// number and number of occurrences are recorded.
func WithUnknownMessages() DecodeOption {
	return func(o *decodeOptions) {
		o.unknownMessages = true
	}
}
