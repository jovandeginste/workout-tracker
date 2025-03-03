package converters

import (
	"archive/zip"
	"bytes"
	"io"
)

func ParseZip(content []byte) ([]*Workout, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		return nil, err
	}

	result := []*Workout{}

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		c, err := readFileFromZip(zipFile)
		if err != nil {
			return nil, err
		}

		gpx, err := ParseCollection(zipFile.Name, c)
		if err != nil {
			return nil, err
		}

		result = append(result, gpx...)
	}

	return result, nil
}

func readFileFromZip(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
