package converters

import (
	"archive/zip"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseZip_SkipsFolderEntries(t *testing.T) {
	const track = `<?xml version="1.0" encoding="UTF-8"?>
<gpx version="1.1" creator="test" xmlns="http://www.topografix.com/GPX/1/1"><trk><trkseg>
<trkpt lat="51.0000" lon="4.0000"><time>2026-10-05T07:00:00Z</time></trkpt>
<trkpt lat="51.0100" lon="4.0000"><time>2026-10-05T07:05:00Z</time></trkpt>
</trkseg></trk></gpx>`

	// What zipping a folder on macOS produces
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for name, content := range map[string]string{
		"rides/":                 "",
		"rides/a.gpx":            track,
		"__MACOSX/rides/._a.gpx": "\x00\x05\x16\x07 AppleDouble",
		"rides/._a.gpx":          "\x00\x05\x16\x07 AppleDouble",
		"rides/.DS_Store":        "\x00\x00\x00\x01Bud1",
	} {
		f, err := zw.Create(name)
		require.NoError(t, err)
		_, err = f.Write([]byte(content))
		require.NoError(t, err)
	}
	require.NoError(t, zw.Close())

	workouts, err := ParseZip(buf.Bytes())
	require.NoError(t, err)
	assert.Len(t, workouts, 1)
}
