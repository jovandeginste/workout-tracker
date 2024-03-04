package database

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/glebarez/sqlite"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Threshold at which point queries are logged as slow
const thresholdSlowQueries = 100 * time.Millisecond

var ErrUnsuportedDriver = errors.New("unsupported driver")

func Connect(driver, dsn string, debug bool, logger *slog.Logger) (*gorm.DB, error) {
	loggerOptions := []slogGorm.Option{
		slogGorm.WithLogger(logger.With("module", "database")),
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

	if err := db.AutoMigrate(&User{}, &Profile{}, &Workout{}, &GPXData{}, &MapData{}, &MapDataDetails{}); err != nil {
		return nil, err
	}

	if err := convertWorkouts(db); err != nil {
		return nil, err
	}

	if err := setUserAPIKeys(db); err != nil {
		return nil, err
	}

	return db, nil
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
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsuportedDriver, driver)
	}
}
