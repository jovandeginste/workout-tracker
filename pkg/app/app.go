package app

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/jovandeginste/workouts/pkg/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	Version string
	Config  Config
	Assets  fs.FS
	Views   fs.FS

	echo           *echo.Echo
	log            *slog.Logger
	db             *gorm.DB
	sessionManager *scs.SessionManager
}

func (a *App) jwtSecret() []byte {
	if a.Config.JWTEncryptionKey == "" {
		a.log.Error("JWTEncryptionKey is not set; generating a random string at startup")

		s, err := util.GenerateRandomString(32)
		if err != nil {
			panic(err)
		}

		a.Config.JWTEncryptionKey = s
	}

	return []byte(a.Config.JWTEncryptionKey)
}

func (a *App) ConfigureDatabase() error {
	db, err := database.Connect(a.Config.DatabaseFile, a.Config.Debug, a.log.With("module", "database"))
	if err != nil {
		return err
	}

	if a.Config.Debug {
		db = db.Debug()
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

func NewApp(version string) *App {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With("app", "workout-tracker", "version", version)
	a := &App{
		log:     logger,
		Version: version,
	}

	return a
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
