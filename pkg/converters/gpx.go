package converters

import "github.com/tkrajina/gpxgo/gpx"

func ParseGPX(content []byte) (*gpx.GPX, error) {
	gpxContent, err := gpx.ParseBytes(content)
	if err != nil {
		return nil, err
	}

	return gpxContent, nil
}
