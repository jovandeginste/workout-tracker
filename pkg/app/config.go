package app

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func (a *App) ReadConfiguration() error {
	viper.SetConfigName("workout-tracker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("WT")

	viper.SetDefault("bind", "[::]:8080")
	viper.SetDefault("web_root", "")
	viper.SetDefault("logging", "true")
	viper.SetDefault("debug", "false")
	viper.SetDefault("database_driver", "sqlite")
	viper.SetDefault("dsn", "./database.db")
	viper.SetDefault("registration_disabled", "false")
	viper.SetDefault("socials_disabled", "false")

	for _, envVar := range []string{
		"bind",
		"web_root",
		"jwt_encryption_key",
		"jwt_encryption_key_file",
		"logging",
		"debug",
		"database_driver",
		"dsn",
		"dsn_file",
		"registration_disabled",
		"socials_disabled",
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

	return nil
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
