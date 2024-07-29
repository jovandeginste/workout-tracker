package importers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrUnsupportedProgram = errors.New("unsupported program")

type Content struct {
	Content  []byte
	Filename string
	Notes    string
	Type     string
}

func Import(program string, headers http.Header, body io.ReadCloser) (*Content, error) {
	defer body.Close()

	switch program {
	case "generic":
		return importGeneric(headers, body)
	case "fitotrack":
		return importFitotrack(headers, body)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedProgram, program)
	}
}
