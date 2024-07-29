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

	b, err := importGeneric(headers, body)
	if err != nil {
		return nil, err
	}

	b.Type = wt
	b.Notes = wn

	return b, nil
}
