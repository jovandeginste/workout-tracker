package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/anyappinc/fitbit"
	"resty.dev/v3"
)

type WorkoutConfig struct {
	URL, APIKey string
	persist     bool
}

func (c *WorkoutConfig) CopyFrom(o *WorkoutConfig) {
	if c.URL == "" {
		c.URL = o.URL
	}

	if c.APIKey == "" {
		c.APIKey = o.APIKey
	}
}

type fitbitSync struct {
	cfg           config
	WorkoutConfig WorkoutConfig
	restClient    *resty.Client

	configDir   string
	configFile  string
	bind        string
	redirectURI *url.URL

	server              *http.Server
	authCodeURL         *url.URL
	state, codeVerifier string

	fitbitClient *fitbit.Client
	waitForAuth  chan bool
	profile      *fitbit.Profile
}

func newCLI() *fitbitSync {
	fs := &fitbitSync{}

	return fs
}

func configDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(dir, "workout-tracker")
}
