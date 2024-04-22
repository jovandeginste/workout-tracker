package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func convertWorkouts(db *gorm.DB) error {
	workouts, err := GetWorkouts(db.Preload("Data").Preload("GPX").Where("LENGTH(gpx_data) > 0 or data IS NOT NULL"))
	if err != nil {
		return err
	}

	for _, w := range workouts {
		db.Logger.Info(context.Background(), fmt.Sprintf("Converting workout gpx data: %d", w.ID))

		if w.GPXData != nil {
			if err := convertGPXData(db, w); err != nil {
				return err
			}
		}

		if w.MapData != nil {
			if err := convertMapData(db, w); err != nil {
				return err
			}
		}

		if err := w.Save(db); err != nil {
			return err
		}
	}

	return nil
}

func convertGPXData(db *gorm.DB, w *Workout) error {
	w.GPX = &GPXData{
		WorkoutID: w.ID,
		Content:   w.GPXData,
		Filename:  w.Filename,
		Checksum:  w.Checksum,
	}
	w.GPXData = nil

	return w.GPX.Save(db)
}

func convertMapData(db *gorm.DB, w *Workout) error {
	w.Data = w.MapData
	w.Data.WorkoutID = w.ID
	w.MapData = nil

	if w.Data.Points == nil {
		return w.Data.Save(db)
	}

	w.Data.Details.Points = w.Data.Points
	w.Data.Points = nil

	if err := w.Data.Details.Save(db); err != nil {
		return err
	}

	return w.Data.Save(db)
}
