package app

import (
	"errors"
	"io/fs"
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	Version        string
	echo           *echo.Echo
	log            *slog.Logger
	db             *gorm.DB
	Assets         fs.FS
	sessionManager *scs.SessionManager
	jwtSecret      []byte
}

func (a *App) Connect() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	a.db = db

	err = db.First(&database.User{}).Error
	if err == nil {
		return nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return a.createAdminUser()
}

func NewApp(l *slog.Logger) *App {
	return &App{
		log:       l,
		jwtSecret: []byte("secret"), // TODO: change to configuration
	}
}

func (a *App) createAdminUser() error {
	u := &database.User{
		Username: "admin",
		Name:     "Administrator",
		Active:   true,
		Admin:    true,
	}

	if err := u.SetPassword("admin"); err != nil {
		return err
	}

	a.log.Info("Creating admin user 'admin/admin'")

	return u.Create(a.db)
}
