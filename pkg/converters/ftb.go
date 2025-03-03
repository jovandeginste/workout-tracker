package converters

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"time"
)

func ParseFTB(content []byte) ([]*Workout, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		return nil, err
	}

	result := []*Workout{}

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		if zipFile.Name != "data.xml" {
			continue
		}

		gpx, err := readZipFile(zipFile)
		if err != nil {
			return nil, err
		}

		result = append(result, gpx...)
	}

	return result, nil
}

func readZipFile(zf *zip.File) ([]*Workout, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	data := &FitoTrackBackup{}
	if err := xml.Unmarshal(d, &data); err != nil {
		return nil, err
	}

	result := []*Workout{}

	for _, is := range data.IndoorWorkouts.IndoorWorkouts {
		result = append(result, convertToWorkout(is))
	}

	return result, nil
}

func convertToWorkout(iw indoorWorkout) *Workout {
	wd := WorkoutData{
		Name:             iw.ExportFileName,
		Type:             iw.WorkoutType,
		Start:            iw.StartTime(),
		Stop:             iw.EndTime(),
		TotalDuration:    time.Duration(iw.Duration * int64(time.Millisecond)),
		TotalRepetitions: iw.Repetitions,
	}

	return &Workout{
		Data:     wd,
		FileType: "xml",
		Content:  nil,
	}
}
