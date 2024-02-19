package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port             int    `mapstructure:"port"`
	JWTEncryptionKey string `mapstructure:"jwt_encryption_key"`
}

func (a *App) ReadConfiguration() error {
	viper.SetConfigName("workout-tracker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("WORKOUT_TRACKER")

	viper.SetDefault("port", 8080)

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
