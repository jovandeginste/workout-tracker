package slogecho

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
)

var _ http.ResponseWriter = (*bodyWriter)(nil)
var _ http.Flusher = (*bodyWriter)(nil)
var _ http.Hijacker = (*bodyWriter)(nil)

type bodyWriter struct {
	http.ResponseWriter
	body    *bytes.Buffer
	maxSize int
	bytes   int
}

// implements http.ResponseWriter
func (w *bodyWriter) Write(b []byte) (int, error) {
	length := len(b)

	if w.body != nil {
		if w.body.Len()+length > w.maxSize {
			w.body.Truncate(min(w.maxSize, length, w.body.Len()))
			w.body.Write(b[:min(w.maxSize-w.body.Len(), length)])
		} else {
			w.body.Write(b)
		}
	}

	w.bytes += length //nolint:staticcheck
	return w.ResponseWriter.Write(b)
}

// implements http.Flusher
func (w *bodyWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// implements http.Hijacker
func (w *bodyWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hi, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hi.Hijack()
	}

	return nil, nil, errors.New("Hijack not supported")
}

func newBodyWriter(writer http.ResponseWriter, maxSize int, recordBody bool) *bodyWriter {
	var body *bytes.Buffer
	if recordBody {
		body = bytes.NewBufferString("")
	}

	return &bodyWriter{
		ResponseWriter: writer,
		body:           body,
		maxSize:        maxSize,
		bytes:          0,
	}
}

type bodyReader struct {
	io.ReadCloser
	body    *bytes.Buffer
	maxSize int
	bytes   int
}

// implements io.Reader
func (r *bodyReader) Read(b []byte) (int, error) {
	n, err := r.ReadCloser.Read(b)
	if r.body != nil && r.body.Len() < r.maxSize {
		if r.body.Len()+n > r.maxSize {
			r.body.Write(b[:min(r.maxSize-r.body.Len(), n)])
		} else {
			r.body.Write(b[:n])
		}
	}
	r.bytes += n
	return n, err
}

func newBodyReader(reader io.ReadCloser, maxSize int, recordBody bool) *bodyReader {
	var body *bytes.Buffer
	if recordBody {
		body = bytes.NewBufferString("")
	}

	return &bodyReader{
		ReadCloser: reader,
		body:       body,
		maxSize:    maxSize,
		bytes:      0,
	}
}
