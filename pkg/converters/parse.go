package converters

import (
	"errors"
	"fmt"
	"path"

	"github.com/tkrajina/gpxgo/gpx"
)

var ErrUnsupportedFile = errors.New("unsupported file")

type parserFunc func(content []byte) (*gpx.GPX, error)

func Parse(filename string, content []byte) (*gpx.GPX, error) {
	c, err := ParseCollection(filename, content)
	if err != nil {
		return nil, err
	}

	if len(c) == 0 {
		return nil, nil
	}

	g := c[0]

	return g, nil
}

func ParseCollection(filename string, content []byte) ([]*gpx.GPX, error) {
	if filename == "" {
		// Assume GPX when filename is empty
		return parseSingle(ParseGPX, content)
	}

	basename := path.Base(filename)

	c, err := parseContent(basename, content)
	if err != nil {
		return nil, err
	}

	for _, g := range c {
		fixName(g, basename)
	}

	return c, nil
}

func fixName(g *gpx.GPX, basename string) {
	if g.Name != "" {
		// We have a name
		return
	}

	if len(g.Tracks) > 0 && g.Tracks[0].Name != "" {
		// Copy the name of the first track
		g.Name = g.Tracks[0].Name
		return
	}

	// Use the filename
	g.Name = basename
}

func parseContent(filename string, content []byte) ([]*gpx.GPX, error) {
	suffix := path.Ext(filename)

	switch suffix {
	case ".gpx":
		return parseSingle(ParseGPX, content)
	case ".fit":
		return parseSingle(ParseFit, content)
	case ".tcx":
		return parseSingle(ParseTCX, content)
	case ".ftb":
		return ParseFTB(content)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedFile, filename)
	}
}

func parseSingle(f parserFunc, content []byte) ([]*gpx.GPX, error) {
	g, err := f(content)
	if err != nil {
		return nil, err
	}

	if g == nil {
		return nil, nil
	}

	return []*gpx.GPX{g}, nil
}
