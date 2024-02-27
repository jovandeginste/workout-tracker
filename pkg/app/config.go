package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	Bind             string `mapstructure:"bind"`
	JWTEncryptionKey string `mapstructure:"jwt_encryption_key"`
	Logging          bool   `mapstructure:"logging"`
	Debug            bool   `mapstructure:"debug"`
	DatabaseDriver   string `mapstructure:"database_driver"`
	DSN              string `mapstructure:"dsn"`
}

func (a *App) ReadConfiguration() error {
	viper.SetConfigName("workout-tracker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("WT")

	viper.SetDefault("bind", "[::]:8080")
	viper.SetDefault("logging", "true")
	viper.SetDefault("debug", "false")
	viper.SetDefault("database_driver", "sqlite")
	viper.SetDefault("dsn", "./database.db")

	for _, envVar := range []string{
		"bind",
		"jwt_encryption_key",
		"logging",
		"debug",
		"database_driver",
		"dsn",
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
