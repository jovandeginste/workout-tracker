package fit

import "fmt"

// An IntegrityError reports that a header or file CRC check failed.
type IntegrityError string

func (e IntegrityError) Error() string {
	return ("integrity error: " + string(e))
}

// A FormatError reports that the input is not valid FIT.
type FormatError string

func (e FormatError) Error() string {
	return "invalid format: " + string(e)
}

// A NotSupportedError reports that the input uses a valid but unimplemented
// FIT feature.
type NotSupportedError string

func (e NotSupportedError) Error() string {
	return "not supported: " + string(e)
}

type ioError struct {
	op  string
	err error
}

func (e ioError) Error() string {
	return fmt.Sprintf("%s: %v", e.op, e.err)
}
