package database

import (
	"log/slog"
	"time"

	"github.com/glebarez/sqlite"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/gorm"
)

func Connect(file string, logger *slog.Logger) (*gorm.DB, error) {
	gormLogger := slogGorm.New(
		slogGorm.WithLogger(logger),
		slogGorm.WithSlowThreshold(time.Second),
		slogGorm.WithTraceAll(),
	)

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&User{}, &Profile{}, &Workout{}); err != nil {
		return nil, err
	}

	return db, nil
}
