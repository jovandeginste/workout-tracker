package importers

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func importFitotrack(c echo.Context, body io.ReadCloser) (*Content, error) {
	headers := c.Request().Header

	if t := headers.Get("FitoTrack-Type"); t != "workout-gpx" {
		return nil, fmt.Errorf("unsupported FitoTrack-Type: %s", t)
	}

	wt := headers.Get("FitoTrack-Workout-Type")
	wn := headers.Get("FitoTrack-Comment")

	b, err := importGeneric(c, body)
	if err != nil {
		return nil, err
	}

	b.Type = wt
	b.Notes = wn
	b.Filename = ""

	return b, nil
}
