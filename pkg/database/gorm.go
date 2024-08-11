package database

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/glebarez/sqlite"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Threshold at which point queries are logged as slow
const thresholdSlowQueries = 100 * time.Millisecond

var ErrUnsuportedDriver = errors.New("unsupported driver")

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

	if err := preMigrationActions(db); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&User{}, &Profile{}, &Config{}, &Equipment{}, &WorkoutEquipment{},
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

	q := db.Unscoped().
		Where("id < (select max(id) from map_data as m where m.workout_id = map_data.workout_id)").
		Delete(&MapData{})
	if q.Error != nil {
		return q.Error
	}

	q = db.Unscoped().
		Where("id < (select max(id) from workouts as w where w.date = workouts.date and w.user_id = workouts.user_id)").
		Delete(&Workout{})

	return q.Error
}

func postMigrationActions(db *gorm.DB) error {
	// Nothing to do for now
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
