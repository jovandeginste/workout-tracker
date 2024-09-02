package importers

import (
	"cmp"
	"io"

	"github.com/labstack/echo/v4"
)

func importGeneric(c echo.Context, body io.ReadCloser) (*Content, error) {
	name := cmp.Or(c.QueryParam("name"), "no-name")
	t := cmp.Or(c.QueryParam("type"), "auto")

	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	g := &Content{
		Filename: name,
		Content:  b,
		Type:     t,
	}

	return g, nil
}
