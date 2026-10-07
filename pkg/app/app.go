package app

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alitto/pond/v2"
	"github.com/cat-dealer/go-rand/v2"
	"github.com/fsouza/slognil"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"github.com/jovandeginste/workout-tracker/v2/pkg/version"
	"github.com/labstack/echo/v5"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"gorm.io/gorm"
)

type App struct {
	Assets       fs.FS
	AssetDir     string
	Translations fs.FS

	echo           *echo.Echo
	logger         *slog.Logger
	rawLogger      *slog.Logger
	db             *gorm.DB
	sessionManager *scs.SessionManager
	translator     *i18n.Locale
	Version        version.Version
	Config         database.Config
	workerPool     pond.Pool
	workerPoolGeo  pond.Pool
}

func (a *App) jwtSecret() []byte {
	if a.Config.JWTEncryptionKey != "" {
		return []byte(a.Config.JWTEncryptionKey)
	}

	if a.Config.JWTEncryptionKeyFile != "" {
		a.logger.Info("reading JWTEncryptionKeyFile", "file", a.Config.JWTEncryptionKeyFile)

		key, err := os.ReadFile(a.Config.JWTEncryptionKeyFile)
		if err == nil {
			a.Config.JWTEncryptionKey = strings.TrimSpace(string(key))
			return []byte(a.Config.JWTEncryptionKey)
		}

		a.logger.Error("could not read JWTEncryptionKeyFile", "error", err)
	}

	a.logger.Error("JWTEncryptionKey is not set; generating a random string at startup")
	a.Config.JWTEncryptionKey = rand.String(32, rand.GetAlphaNumericPool())

	return []byte(a.Config.JWTEncryptionKey)
}

func (a *App) Serve() error {
	go a.BackgroundWorker()

	// Close the database once the HTTP server has shut down, so SQLite runs
	// its WAL checkpoint and removes the -wal and -shm files.
	defer func() {
		if err := a.closeDB(); err != nil {
			a.logger.Error("failed to close database connection", "error", err)
		}
	}()

	a.logger.Info("Starting web server on " + a.Config.Bind)

	return a.echo.Start(a.Config.Bind)
}

func (a *App) closeDB() error {
	sqlDB, err := a.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (a *App) Configure() error {
	if err := a.ReadConfiguration(); err != nil {
		return err
	}

	a.ConfigureLogger()

	if err := a.ConfigureLocalizer(); err != nil {
		return err
	}

	if err := a.ConfigureDatabase(); err != nil {
		return err
	}

	a.ConfigureGeocoder()

	if err := database.InitTZFinder(); err != nil {
		return err
	}

	if err := a.Config.UpdateFromDatabase(a.db); err != nil {
		return err
	}

	if err := a.ConfigureWebserver(); err != nil {
		return err
	}

	return nil
}

func (a *App) ConfigureGeocoder() {
	if a.Config.Offline {
		geocoder.ForceOffline()
		return
	}

	geocoder.SetClient(a.logger, a.Version.UserAgent())
}

func (a *App) ConfigureDatabase() error {
	a.SetDSN()

	a.logger.Info("Connecting to the database '" + a.Config.DatabaseDriver + "': " + a.Config.DSN)

	db, err := database.Connect(a.Config.DatabaseDriver, a.Config.DSN, a.Config.Debug, a.rawLogger)
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

func newLogger(enabled bool) *slog.Logger {
	if !enabled {
		return slognil.NewLogger()
	}

	return slog.New(newLogHandler())
}

func newLogHandler() slog.Handler {
	w := os.Stderr
	if isatty.IsTerminal(w.Fd()) {
		return tint.NewTextHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		})
	}

	return slog.NewJSONHandler(w, nil)
}

func (a *App) ConfigureLogger() {
	logger := newLogger(a.Config.Logging).
		With("app", "workout-tracker").
		With("version", a.Version.RefName).
		With("sha", a.Version.Sha)

	a.rawLogger = logger
	a.logger = logger.With("module", "app")
}

func NewApp(v version.Version) *App {
	return &App{
		Version:   v,
		logger:    newLogger(false),
		rawLogger: newLogger(false),
	}
}

func (a *App) createAdminUser() error {
	u := &database.User{
		UserData: database.UserData{
			Username: "admin",
			Name:     "Administrator",
			Active:   true,
			Admin:    true,
		},
	}

	if err := u.SetPassword("admin"); err != nil {
		return err
	}

	a.logger.Warn("Creating admin user '" + u.Username + "', with password 'admin'")

	u.Profile.User = u

	return u.Create(a.db)
}

func (a *App) DB() *gorm.DB {
	return a.db
}

func (a *App) Logger() *slog.Logger {
	return a.logger
}

func (a *App) Reverse(name string, params ...any) string {
	if rev, err := a.echo.Router().Routes().Reverse(name, params...); err == nil {
		return rev
	}
	return ""
}
