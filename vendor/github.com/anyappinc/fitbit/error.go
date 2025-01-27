package fitbit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Error is the interface that has ability to return raw error returned from Fitbit APIs.
//
// This also implements the builtin error interface.
type Error interface {
	Raw() []byte
	Error() string
}

// RawError represents an error that only contains raw error returned from Fitbit APIs.
type RawError struct {
	raw []byte
}

// Raw implements ability to return raw error.
func (e *RawError) Raw() []byte {
	return e.raw
}

// Error implements the error interface.
func (e *RawError) Error() string {
	return string(e.raw)
}

// MessageError represents an error returned from Fitbit APIs
// that contains error type and message fields.
type MessageError struct {
	*RawError
	Type    string
	Message string
}

// Error implements the error interface.
func (e *MessageError) Error() string {
	return e.Message
}

// FieldNameMessageError represents an error returned from Fitbit APIs
// that contains MessageError plus fieldName.
type FieldNameMessageError struct {
	*MessageError
	FieldName string
}

// DetailError represents an error returned from Fitbit APIs
// that contains error title and detail fields.
type DetailError struct {
	*RawError
	Title  string
	Detail string
}

// Error implements the error interface.
func (e *DetailError) Error() string {
	return e.Detail
}

// DetailSourceError represents an error returned from Fitbit APIs
// that contains DetailError plus source parameter.
type DetailSourceError struct {
	*DetailError
	Source struct {
		Parameter string
	}
}

type (
	rawErrorResponse struct {
		Success bool                     `json:"success"`
		Errors  []map[string]interface{} `json:"errors"`
	}

	// ErrorResponse represents an error response from Fitbit APIs.
	ErrorResponse struct {
		Success bool    `json:"success"`
		Errors  []Error `json:"errors"`
	}
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *ErrorResponse) UnmarshalJSON(b []byte) error {
	raw := rawErrorResponse{
		Success: true,
		Errors:  nil,
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	e.Success = raw.Success
	if len(raw.Errors) == 0 {
		return nil
	}

	e.Errors = make([]Error, len(raw.Errors))
	for i := range raw.Errors {
		_err, err := extractError(raw.Errors[i])
		if err != nil {
			return err
		}
		e.Errors[i] = _err
	}

	return nil
}

func extractError(e map[string]interface{}) (Error, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	rawError := &RawError{b}

	// Check if it is MessageError or FieldNameMessageError
	errorType, ok1 := e["errorType"].(string)
	message, ok2 := e["message"].(string)
	if ok1 && ok2 {
		messageError := &MessageError{
			RawError: rawError,
			Type:     errorType,
			Message:  message,
		}
		if fieldName, ok := e["fieldName"].(string); ok {
			return &FieldNameMessageError{
				MessageError: messageError,
				FieldName:    fieldName,
			}, nil
		}
		return messageError, nil
	}

	// Check if it is DetailError or DetailSourceError
	title, ok3 := e["title"].(string)
	detail, ok4 := e["detail"].(string)
	if ok3 && ok4 {
		detailError := &DetailError{
			RawError: rawError,
			Title:    title,
			Detail:   detail,
		}
		if source, ok := e["source"].(map[string]interface{}); ok {
			if parameter, ok := source["parameter"].(string); ok {
				return &DetailSourceError{
					DetailError: detailError,
					Source: struct {
						Parameter string
					}{
						Parameter: parameter,
					},
				}, nil
			}
		}
		return detailError, nil
	}

	return rawError, nil
}

func parseErrorResponse(b []byte) (*ErrorResponse, error) {
	if !json.Valid(b) {
		return nil, nil
	}
	var errResp ErrorResponse
	if err := json.Unmarshal(b, &errResp); err != nil {
		return nil, err
	}
	return &errResp, nil
}

// APIError represents an error that occurred on a request to Fitbit APIs.
type APIError struct {
	ErrResp  *ErrorResponse
	HTTPResp *http.Response
	Body     []byte
}

// Error implements the error interface.
func (ae *APIError) Error() string {
	if ae.ErrResp == nil || len(ae.ErrResp.Errors) == 0 {
		return ae.HTTPResp.Status
	}
	errMsgs := make([]string, len(ae.ErrResp.Errors))
	for i, e := range ae.ErrResp.Errors {
		errMsgs[i] = e.Error()
	}
	return strings.Join(errMsgs, "\n")
}

// RequestError represents an error that occurred in a request process.
type RequestError struct {
	Op  string
	URL string
	Err error
}

func wrapAsRequestError(op, url string, err error) error {
	if err == nil {
		return nil
	}
	return &RequestError{
		Op:  op,
		URL: url,
		Err: err,
	}
}

// Unwrap adds support for `errors` error wrapping.
func (e *RequestError) Unwrap() error {
	return e.Err
}

// Error implements the error interface.
func (e *RequestError) Error() string {
	return fmt.Sprintf("%s %q: %s", e.Op, e.URL, e.Err)
}

func parseError(r *http.Response, b []byte) error {
	errResp, err := parseErrorResponse(b)
	if err != nil {
		return err
	}
	if errResp == nil || !errResp.Success {
		return &APIError{
			ErrResp:  errResp,
			HTTPResp: r,
			Body:     b,
		}
	}
	return nil
}
