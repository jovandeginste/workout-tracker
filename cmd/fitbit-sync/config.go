package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/anyappinc/fitbit"
)

type config struct {
	WorkoutConfig WorkoutConfig

	FitbitConfig struct {
		ClientID     string
		ClientSecret string
		UserID       string
	}
	Token *fitbit.Token
}

func (fs *fitbitSync) setDefaults() {
	fs.waitForAuth = make(chan bool, 1)
}

func (fs *fitbitSync) ConfigFile() string {
	p := fs.configDir
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(p, 0o700); err != nil {
			log.Fatal(err)
		}
	}

	c := path.Join(p, fs.configFile)

	return c
}

func (fs *fitbitSync) loadConfig() error {
	fs.setDefaults()

	f := fs.ConfigFile()
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	file, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("could not open '%s': %w", f, err)
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fs.cfg); err != nil {
		return fmt.Errorf("could not parse '%s': %w", f, err)
	}

	return nil
}

func (fs *fitbitSync) saveConfig() error {
	f := fs.ConfigFile()

	file, err := os.Create(f)
	if err != nil {
		return fmt.Errorf("could not create '%s': %w", f, err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(fs.cfg); err != nil {
		return fmt.Errorf("could not encode '%s': %w", f, err)
	}

	return nil
}

func (fs *fitbitSync) initClient() {
	fs.fitbitClient = fitbit.NewClient(
		fs.cfg.FitbitConfig.ClientID,
		fs.cfg.FitbitConfig.ClientSecret,
		fitbit.ServerApplication,
		&fitbit.Scope{Activity: true, Weight: true, Profile: true, Heartrate: true, Location: true},
	)
	fs.fitbitClient.SetUpdateTokenFunc(fs.updateTokenFunc)
	fs.fitbitClient.EnableDebugMode()
}

func (fs *fitbitSync) reinitClient() error {
	if err := fs.cfg.valid(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	fs.initClient()

	p, _, _, err := fs.fitbitClient.GetProfile(context.Background(), fs.cfg.FitbitConfig.UserID, fs.cfg.Token)
	if err != nil {
		return fmt.Errorf("could not get profile: %w", err)
	}

	fs.profile = p

	return nil
}

func (cfg *config) valid() error {
	if cfg.FitbitConfig.UserID == "" || cfg.Token == nil || cfg.Token.RefreshToken == "" {
		return errors.New("configuration missing - did you run init?")
	}

	return nil
}

func (fs *fitbitSync) updateTokenFunc(oldToken, newToken *fitbit.Token) error {
	if oldToken == newToken {
		return nil
	}

	log.Print("Updating token.")

	fs.cfg.Token = newToken

	return fs.saveConfig()
}
