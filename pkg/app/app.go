package app

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cat-dealer/go-rand/v2"
	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"golang.org/x/text/language"
	"gorm.io/gorm"

	"github.com/vorlif/spreak"
	"github.com/vorlif/spreak/humanize"
)

type Version struct {
	BuildTime string
	Ref       string
	RefName   string
	RefType   string
	Sha       string
}

type App struct {
	Version      Version
	Config       Config
	Assets       fs.FS
	Views        fs.FS
	Translations fs.FS

	echo           *echo.Echo
	logger         *slog.Logger
	rawLogger      *slog.Logger
	db             *gorm.DB
	sessionManager *scs.SessionManager
	translator     *spreak.Bundle
	humanizer      *humanize.Collection
}

func (a *App) jwtSecret() []byte {
	if a.Config.JWTEncryptionKey == "" {
		a.logger.Error("JWTEncryptionKey is not set; generating a random string at startup")

		a.Config.JWTEncryptionKey = rand.String(32, rand.GetAlphaNumericPool())
	}

	return []byte(a.Config.JWTEncryptionKey)
}

func (a *App) ConfigureLocalizer() error {
	bundle, err := spreak.NewBundle(
		// Set the language used in the program code/templates
		spreak.WithSourceLanguage(language.English),
		// Set the path from which the translations should be loaded
		spreak.WithDomainFs(spreak.NoDomain, a.Translations),
		// Specify the languages you want to load
		spreak.WithLanguage(translations()...),
	)
	if err != nil {
		return err
	}

	a.translator = bundle

	a.humanizer = humanize.MustNew(
		humanize.WithLocale(humanLocales()...),
	)

	return nil
}

func (a *App) Serve() error {
	go a.BackgroundWorker()

	a.logger.Info("Starting web server on " + a.Config.Bind)
	return a.echo.Start(a.Config.Bind)
}

func (a *App) Configure() error {
	if err := a.ReadConfiguration(); err != nil {
		return err
	}

	if err := a.ConfigureLocalizer(); err != nil {
		return err
	}

	if err := a.ConfigureDatabase(); err != nil {
		return err
	}

	if err := a.ConfigureWebserver(); err != nil {
		return err
	}

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
		With("app", "workout-tracker", "version", version.RefName, "sha", version.Sha)

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

func (a *App) BackgroundWorker() {
	l := a.logger.With("module", "worker")

	for {
		l.Info("Worker started...")

		var w []database.Workout

		if err := a.db.Where(&database.Workout{Dirty: true}).Limit(10).Find(&w).Error; err != nil {
			l.Error("Worker error: " + err.Error())
		}

		for _, v := range w {
			if err := v.UpdateData(a.db); err != nil {
				l.Error("Worker error: " + err.Error())
			}
		}

		l.Info("Worker finished...")
		time.Sleep(time.Minute)
	}
}
