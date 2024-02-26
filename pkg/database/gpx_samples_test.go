package database

import (
	"testing"
	"testing/fstest"
)

var gpxFS fstest.MapFS

func populateGPXFS(t *testing.T) {
	gpxFS = fstest.MapFS{}

	gpxFS["sample1.gpx"] = &fstest.MapFile{Data: []byte(GpxSample1)}
}
