package converters

import (
	"errors"
	"fmt"
	"path"

	"github.com/tkrajina/gpxgo/gpx"
)

var ErrUnsupportedFile = errors.New("unsupported file")

func Parse(filename string, content []byte) (*gpx.GPX, error) {
	if filename == "" {
		// Assume GPX when filename is empty
		return ParseGPX(content)
	}

	suffix := path.Ext(filename)

	switch suffix {
	case ".gpx":
		return ParseGPX(content)
	case ".fit":
		return ParseFit(content)
	case ".tcx":
		return ParseTCX(content)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedFile, filename)
	}
}
