package converters

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/tkrajina/gpxgo/gpx"
)

var (
	ErrUnsupportedFile = errors.New("unsupported file")
	SupportedFileTypes = []string{".fit", ".ftb", ".gpx", ".tcx", ".zip"}
)

type (
	parserFunc func(content []byte) (*gpx.GPX, error)
)

func Parse(filename string, content []byte) (*Workout, error) {
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

func ParseCollection(filename string, content []byte) ([]*Workout, error) {
	if filename == "" {
		// Assume GPX when filename is empty
		return parseSingle(ParseGPX, "gpx", content)
	}

	basename := path.Base(filename)

	c, err := parseContent(basename, content)
	if err != nil {
		return nil, err
	}

	for _, g := range c {
		g.FixName(basename)
	}

	return c, nil
}

func parseContent(filename string, content []byte) ([]*Workout, error) {
	suffix := strings.ToLower(path.Ext(filename))

	switch suffix {
	case ".gpx":
		return parseSingle(ParseGPX, "gpx", content)
	case ".fit":
		return parseSingle(ParseFit, "fit", content)
	case ".tcx":
		return parseSingle(ParseTCX, "tcx", content)
	case ".zip":
		return ParseZip(content)
	case ".ftb":
		return ParseFTB(content)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedFile, filename)
	}
}

func parseSingle(f parserFunc, t string, content []byte) ([]*Workout, error) {
	g, err := f(content)
	if err != nil {
		return nil, err
	}

	if g == nil {
		return nil, nil
	}

	w := &Workout{
		GPX:      g,
		FileType: t,
		Content:  content,
	}

	return []*Workout{w}, nil
}
