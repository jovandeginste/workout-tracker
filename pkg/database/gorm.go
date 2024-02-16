package database

import (
	"github.com/glebarez/sqlite"
	"github.com/jovandeginste/workouts/pkg/user"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&user.User{}, &user.Profile{}, user.Workout{}); err != nil {
		return nil, err
	}

	return db, nil
}
