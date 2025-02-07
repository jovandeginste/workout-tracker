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

	basename := path.Base(filename)

	c, err := parseContent(basename, content)
	if err != nil {
		return nil, err
	}

	switch {
	case c.Name != "":
		// We have a name
	case len(c.Tracks) > 0 && c.Tracks[0].Name != "":
		// Copy the name of the first track
		c.Name = c.Tracks[0].Name
	default:
		// Use the filename
		c.Name = basename
	}

	return c, nil
}

func parseContent(filename string, content []byte) (*gpx.GPX, error) {
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
