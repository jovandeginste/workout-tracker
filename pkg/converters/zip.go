package converters

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"path"
	"slices"
	"strings"
)

func ParseZip(content []byte) ([]*Workout, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		return nil, err
	}

	result := []*Workout{}
	var errs []error

	// Read all the supported files from the zip archive. Unsupported files
	// (READMEs, images, thumbnails, ...) are ignored; a broken workout file is
	// collected and reported together with the other failures.
	for _, zipFile := range zipReader.File {
		name := zipEntryName(zipFile)
		if skipZipEntry(zipFile, name) {
			continue
		}

		c, err := readFileFromZip(zipFile)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
			continue
		}

		gpx, err := ParseCollection(name, c)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
			continue
		}

		result = append(result, gpx...)
	}

	return result, errors.Join(errs...)
}

func zipEntryName(zf *zip.File) string {
	name := strings.ReplaceAll(zf.Name, "\\", "/")

	return path.Clean(strings.TrimPrefix(name, "/"))
}

// skipZipEntry reports whether an archive entry should not be parsed. It skips
// directories, hidden files and directories (including __MACOSX and AppleDouble
// files), path traversal and anything that is not a supported workout file
// type.
func skipZipEntry(zf *zip.File, name string) bool {
	if zf.FileInfo().IsDir() {
		return true
	}

	if name == "." || name == ".." || strings.HasPrefix(name, "../") {
		return true
	}

	if !slices.Contains(SupportedFileTypes, strings.ToLower(path.Ext(name))) {
		return true
	}

	for _, part := range strings.Split(name, "/") {
		if strings.HasPrefix(part, ".") || strings.EqualFold(part, "__MACOSX") {
			return true
		}
	}

	return false
}

func readFileFromZip(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
