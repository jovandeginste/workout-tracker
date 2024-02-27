package database

import (
	"testing/fstest"
)

var gpxFS fstest.MapFS

func populateGPXFS() {
	gpxFS = fstest.MapFS{}

	gpxFS["sample1.gpx"] = &fstest.MapFile{Data: []byte(GpxSample1)}
}
