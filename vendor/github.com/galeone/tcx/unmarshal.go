package tcx

import (
	"encoding/xml"
	"os"
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
