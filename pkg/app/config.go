package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	Bind             string `mapstructure:"bind"`
	JWTEncryptionKey string `mapstructure:"jwt_encryption_key"`
	Debug            bool   `mapstructure:"debug"`
}

func (a *App) ReadConfiguration() error {
	viper.SetConfigName("workout-tracker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("WORKOUT_TRACKER")

	viper.SetDefault("bind", "[::]:8080")
	viper.SetDefault("debug", "false")

	if err := viper.BindEnv("jwt_encryption_key"); err != nil {
		return err
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
