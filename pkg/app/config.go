package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	Bind             string `mapstructure:"bind"`
	JWTEncryptionKey string `mapstructure:"jwt_encryption_key"`
	Debug            bool   `mapstructure:"debug"`
	DatabaseFile     string `mapstructure:"database_file"`
}

func (a *App) ReadConfiguration() error {
	viper.SetConfigName("workout-tracker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("WT")

	viper.SetDefault("bind", "[::]:8080")
	viper.SetDefault("debug", "false")
	viper.SetDefault("database_file", "./database.db")

	for _, envVar := range []string{
		"bind",
		"jwt_encryption_key",
		"debug",
		"database_file",
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
