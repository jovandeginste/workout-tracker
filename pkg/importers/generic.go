package importers

import (
	"io"
	"net/http"
)

func importGeneric(_ http.Header, body io.ReadCloser) (*Content, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	g := &Content{
		Content: b,
		Type:    "auto",
	}

	return g, nil
}
