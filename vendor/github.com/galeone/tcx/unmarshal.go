package tcx

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
)

func ReadTpts(path string) (track *Track, err error) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	track = new(Track)
	err = xml.Unmarshal(filebytes, track)
	return
}

func ReadLap(path string) (lap *Lap, err error) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lap = new(Lap)
	err = xml.Unmarshal(filebytes, lap)
	return
}

func ReadActivity(path string) (act *Activity, err error) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	act = new(Activity)
	err = xml.Unmarshal(filebytes, act)
	return
}

func ReadActivities(path string) (acts *Activities, err error) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	acts = new(Activities)
	err = xml.Unmarshal(filebytes, acts)
	return
}

func ReadFile(path string) (db *TCXDB, err error) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	db = new(TCXDB)
	err = xml.Unmarshal(filebytes, db)
	return
}

// ReadDir recursively walks and loads directories of tcx files
func ReadDir(dir string) (dbs []*TCXDB, err error) {
	walk := func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			split := strings.Split(path, ".")
			if split[len(split)-1] == "tcx" {
				db, err := ReadFile(path)
				if err != nil {
					return err
				}
				dbs = append(dbs, db)
			}
		}
		if info.Mode().IsDir() && path != dir {
			_, err = ReadDir(path)
		}
		return err
	}
	if err = filepath.Walk(dir, walk); err != nil {
		return nil, err
	}
	return dbs, nil
}
