package converters

import (
	"archive/zip"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTrack = `<?xml version="1.0" encoding="UTF-8"?>
<gpx version="1.1" creator="test" xmlns="http://www.topografix.com/GPX/1/1"><trk><trkseg>
<trkpt lat="51.0000" lon="4.0000"><time>2026-10-05T07:00:00Z</time></trkpt>
<trkpt lat="51.0100" lon="4.0000"><time>2026-10-05T07:05:00Z</time></trkpt>
</trkseg></trk></gpx>`

func makeTestZip(t *testing.T, files map[string]string) []byte {
	t.Helper()

	var buf bytes.Buffer

	zw := zip.NewWriter(&buf)
	for name, content := range files {
		f, err := zw.Create(name)
		require.NoError(t, err)

		_, err = f.Write([]byte(content))
		require.NoError(t, err)
	}

	require.NoError(t, zw.Close())

	return buf.Bytes()
}

func TestParseZip_SkipsFolderAndJunkEntries(t *testing.T) {
	content := makeTestZip(t, map[string]string{
		"rides/":                 "",
		"rides/a.gpx":            testTrack,
		"__MACOSX/rides/._a.gpx": "\x00\x05\x16\x07 AppleDouble",
		"rides/._a.gpx":          "\x00\x05\x16\x07 AppleDouble",
		"rides/.DS_Store":        "\x00\x00\x00\x01Bud1",
		"rides/.git/config":      "[core]",
		"rides/.git/track.gpx":   testTrack,
		"rides/Thumbs.db":        "\x00\x00\x00\x01Thumbs",
		"rides/desktop.ini":      "[.ShellClassInfo]",
		"rides/Icon\r":           "custom icon",
		"rides/README.txt":       "not a workout",
		"rides/photo.jpg":        "\xff\xd8\xff",
	})

	workouts, err := ParseZip(content)
	require.NoError(t, err)
	assert.Len(t, workouts, 1)
}

func TestParseZip_NormalizesEntryPaths(t *testing.T) {
	content := makeTestZip(t, map[string]string{
		"./rides/a.gpx":          testTrack,
		"rides/../rides/b.gpx":   testTrack,
		"/rides/c.gpx":           testTrack,
		"../outside.gpx":         testTrack,
		"rides/../../escape.gpx": testTrack,
	})

	workouts, err := ParseZip(content)
	require.NoError(t, err)
	assert.Len(t, workouts, 3)
}

func TestParseZip_CollectsErrors(t *testing.T) {
	content := makeTestZip(t, map[string]string{
		"good.gpx":    testTrack,
		"bad-one.gpx": "not xml",
		"bad-two.gpx": "also not xml",
	})

	workouts, err := ParseZip(content)
	require.Error(t, err)
	assert.Len(t, workouts, 1)
	assert.Contains(t, err.Error(), "bad-one.gpx")
	assert.Contains(t, err.Error(), "bad-two.gpx")
}
