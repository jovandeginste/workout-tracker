package converters

import "github.com/tkrajina/gpxgo/gpx"

func ParseFTB(ftbFile []byte) ([]*gpx.GPX, error) {
	g, err := parseFTB(ftbFile)
	if err != nil {
		return nil, err
	}

	if g == nil {
		return nil, nil
	}

	return []*gpx.GPX{g}, nil
}

func parseFTB(ftbFile []byte) (*gpx.GPX, error) {
	// TODO
	return nil, nil
}
