package app

import (
	"os"
	"strings"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/spf13/viper"
)

func (a *App) ReadConfiguration() error {
	viper.Reset()

	viper.SetConfigName("workout-tracker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("WT")

	viper.SetDefault("bind", "[::]:8080")
	viper.SetDefault("web_root", "")
	viper.SetDefault("logging", true)
	viper.SetDefault("debug", false)
	viper.SetDefault("offline", false)
	viper.SetDefault("compress", true)
	viper.SetDefault("database_driver", "sqlite")
	viper.SetDefault("dsn", "./database.db")
	viper.SetDefault("registration_disabled", false)
	viper.SetDefault("socials_disabled", false)
	viper.SetDefault("login_custom_text", "")
	viper.SetDefault("worker_delay_seconds", 60)

	for _, envVar := range []string{
		"bind",
		"web_root",
		"jwt_encryption_key",
		"jwt_encryption_key_file",
		"logging",
		"offline",
		"compress",
		"debug",
		"database_driver",
		"dsn",
		"dsn_file",
		"registration_disabled",
		"socials_disabled",
		"login_custom_text",
		"worker_delay_seconds",
	} {
		if err := viper.BindEnv(envVar); err != nil {
			return err
		}
	}

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	if err := viper.Unmarshal(&a.Config); err != nil {
		return err
	}

	// Only non-default values from the configuration file or environment take
	// precedence over the database, so only those are marked as overridden.
	a.Config.Overrides = database.ConfigOverrides{
		RegistrationDisabled: a.Config.RegistrationDisabled && isConfigOverridden("registration_disabled"),
		SocialsDisabled:      a.Config.SocialsDisabled && isConfigOverridden("socials_disabled"),
		LoginCustomText:      a.Config.LoginCustomText != "" && isConfigOverridden("login_custom_text"),
	}

	return nil
}

// isConfigOverridden reports whether a setting is explicitly set in the
// configuration file or through an environment variable.
func isConfigOverridden(key string) bool {
	if viper.InConfig(key) {
		return true
	}

	_, ok := os.LookupEnv("WT_" + strings.ToUpper(key))

	return ok
}

func (a *App) ResetConfiguration() error {
	if err := a.ReadConfiguration(); err != nil {
		return err
	}

	return a.Config.UpdateFromDatabase(a.db)
}

func (a *App) SetDSN() {
	if a.Config.DSN != "" {
		return
	}

	if a.Config.DSNFile == "" {
		return
	}

	a.logger.Info("reading DSNFile", "file", a.Config.DSNFile)

	dsn, err := os.ReadFile(a.Config.DSNFile)
	if err != nil {
		a.logger.Error("could not read DSN file", "error", err)
		return
	}

	a.Config.DSN = strings.TrimSpace(string(dsn))
}
