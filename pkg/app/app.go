package app

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/jovandeginste/workout-tracker/pkg/util"
	"github.com/labstack/echo/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"gorm.io/gorm"
)

type Version struct {
	BuildTime string
	Ref       string
	Commit    string
}

type App struct {
	Version Version
	Config  Config
	Assets  fs.FS
	Views   fs.FS

	echo           *echo.Echo
	logger         *slog.Logger
	rawLogger      *slog.Logger
	db             *gorm.DB
	sessionManager *scs.SessionManager
}

func (a *App) jwtSecret() []byte {
	if a.Config.JWTEncryptionKey == "" {
		a.logger.Error("JWTEncryptionKey is not set; generating a random string at startup")

		s, err := util.GenerateRandomString(32)
		if err != nil {
			panic(err)
		}

		a.Config.JWTEncryptionKey = s
	}

	return []byte(a.Config.JWTEncryptionKey)
}

func (a *App) Configure() error {
	if err := a.ReadConfiguration(); err != nil {
		return err
	}

	if err := a.ConfigureDatabase(); err != nil {
		return err
	}

	if err := a.ConfigureWebserver(); err != nil {
		return err
	}

	a.logger = a.logger.With("module", "app")

	return nil
}

func (a *App) ConfigureDatabase() error {
	a.logger.Info("Connecting to the database: " + a.Config.DatabaseFile)

	db, err := database.Connect(a.Config.DatabaseFile, a.Config.Debug, a.rawLogger)
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

func NewApp(version Version) *App {
	logger := slog.New(newLogger()).
		With("app", "workout-tracker", "version", version)

	a := &App{
		rawLogger: logger,
		logger:    logger.With("module", "app"),
		Version:   version,
	}

	return a
}

func (a *App) createAdminUser() error {
	u := &database.User{
		Username: "admin@localhost",
		Name:     "Administrator",
		Active:   true,
		Admin:    true,
	}

	if err := u.SetPassword("admin"); err != nil {
		return err
	}

	a.logger.Warn("Creating admin user '" + u.Username + "', with password 'admin'")

	return u.Create(a.db)
}
