package database

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/labstack/gommon/log"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Threshold at which point queries are logged as slow
const thresholdSlowQueries = 100 * time.Millisecond

var ErrUnsuportedDriver = errors.New("unsupported driver")

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uint64 `gorm:"primaryKey"`
}

func Connect(driver, dsn string, debug bool, logger *slog.Logger) (*gorm.DB, error) {
	loggerOptions := []slogGorm.Option{
		slogGorm.WithHandler(logger.With("module", "database").Handler()),
		slogGorm.WithSlowThreshold(thresholdSlowQueries),
	}

	if debug {
		loggerOptions = append(loggerOptions, slogGorm.WithTraceAll())
	}

	gormLogger := slogGorm.New(
		loggerOptions...,
	)

	d, err := dialectorFor(driver, dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(d, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	if err := db.Use(NewMemoryCache()); err != nil {
		return nil, err
	}

	if err := preMigrationActions(db); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&User{}, &Profile{}, &Config{}, &Equipment{}, &WorkoutEquipment{}, &Measurement{},
		&Workout{}, &GPXData{}, &MapData{}, &MapDataDetails{}, &RouteSegment{}, &RouteSegmentMatch{},
	); err != nil {
		return nil, err
	}

	if err := postMigrationActions(db); err != nil {
		return nil, err
	}

	if err := setUserAPIKeys(db); err != nil {
		return nil, err
	}

	return db, nil
}

func preMigrationActions(db *gorm.DB) error {
	if !db.Migrator().HasTable(&MapData{}) {
		return nil
	}

	q := db.
		Where("id < (select max(id) from map_data as m where m.workout_id = map_data.workout_id)").
		Delete(&MapData{})
	if q.Error != nil {
		return q.Error
	}

	q = db.
		Where("id < (select max(id) from workouts as w where w.date = workouts.date and w.user_id = workouts.user_id)").
		Delete(&Workout{})
	if q.Error != nil {
		return q.Error
	}

	q = db.
		Where("map_data_id IN (SELECT map_data_id FROM map_data_details as mdd where map_data_details.created_at < mdd.created_at)").
		Delete(&MapDataDetails{})

	return q.Error
}

func postMigrationActions(db *gorm.DB) error {
	workouts, err := GetWorkouts(db)
	if err != nil {
		return err
	}

	for _, w := range workouts {
		if !w.HasTracks() || w.Data.ExtraMetrics != nil {
			continue
		}

		w.Data.UpdateExtraMetrics()

		if err := w.Save(db); err != nil {
			log.Error("Failed to update extra metrics", "err", err)
		}
	}

	return nil
}

func setUserAPIKeys(db *gorm.DB) error {
	users, err := GetUsers(db)
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.APIKey != "" {
			continue
		}

		if err := u.Save(db); err != nil {
			return err
		}
	}

	return nil
}

func dialectorFor(driver, dsn string) (gorm.Dialector, error) {
	switch driver {
	case "sqlite":
		return sqlite.Open(dsn), nil
	case "memory":
		return sqlite.Open(":memory:"), nil
	case "mysql":
		return mysql.Open(dsn), nil
	case "postgres":
		return postgres.Open(dsn), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsuportedDriver, driver)
	}
}
