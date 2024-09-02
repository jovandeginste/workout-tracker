package importers

import (
	"errors"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

var ErrUnsupportedProgram = errors.New("unsupported program")

type Content struct {
	Content  []byte
	Filename string
	Notes    string
	Type     string
}

func Import(program string, c echo.Context, body io.ReadCloser) (*Content, error) {
	defer body.Close()

	switch program {
	case "generic":
		return importGeneric(c, body)
	case "fitotrack":
		return importFitotrack(c, body)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedProgram, program)
	}
}
