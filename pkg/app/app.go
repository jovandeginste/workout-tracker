package app

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/jovandeginste/workouts/pkg/util"
	"github.com/labstack/echo/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
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

func newLogger() slog.Handler {
	w := os.Stdout
	if isatty.IsTerminal(w.Fd()) {
		return tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		})
	}

	return slog.NewJSONHandler(w, nil)
}

func NewApp(version string) *App {
	logger := slog.New(newLogger()).
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

	a.log.Warn("Creating admin user 'admin/admin'")

	return u.Create(a.db)
}
