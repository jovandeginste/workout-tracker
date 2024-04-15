package importers

import (
	"fmt"
	"io"
	"net/http"
)

func importFitotrack(headers http.Header, body io.ReadCloser) (*Content, error) {
	if t := headers.Get("FitoTrack-Type"); t != "workout-gpx" {
		return nil, fmt.Errorf("unsupported FitoTrack-Type: %s", t)
	}

	wt := headers.Get("FitoTrack-Workout-Type")
	wn := headers.Get("FitoTrack-Comment")

	defer body.Close()

	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	g := &Content{
		Content: b,
		Type:    wt,
		Notes:   wn,
	}

	return g, nil
}
