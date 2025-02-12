package logger

import (
	"io"
	"log"
	"os"
)

var (
	Debug   = log.New(os.Stdout, "[Debug] Fitbit API Client: ", log.Lshortfile) // Debug is a logger for debug level messages.
	Info    = log.New(os.Stdout, "[Info] Fitbit API Client: ", 0)               // Info is a logger for infomation level messages.
	Warn    = log.New(os.Stderr, "[Warning] Fitbit API Client: ", 0)            // Warn is a logger for warning level messages.
	Err     = log.New(os.Stderr, "[Error] Fitbit API Client: ", 0)              // Err is a logger for error level messages.
	loggers = []*log.Logger{Debug, Info, Warn, Err}
)

// SetLogsFlags sets the output flags for all types of loggers.
func SetLogsFlags(flags int) {
	for _, logger := range loggers {
		logger.SetFlags(flags)
	}
}

// SetLogsOutput sets the output destination for all types of loggers.
func SetLogsOutput(w io.Writer) {
	for _, logger := range loggers {
		logger.SetOutput(w)
	}
}

// SetLogsPrefix sets the output prefix for all types of loggers.
func SetLogsPrefix(prefix string) {
	for _, logger := range loggers {
		logger.SetPrefix(prefix)
	}
}
