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
	if program == "fitotrack" {
		return importFitotrack(headers, body)
	}

	return nil, fmt.Errorf("%w: %s", ErrUnsupportedProgram, program)
}
