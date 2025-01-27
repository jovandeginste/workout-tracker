package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/anyappinc/fitbit"
	"github.com/aquasecurity/table"
	"github.com/skratchdot/open-golang/open"
)

const (
	fitbitConfig = "fitbit.json"
	redirectURI  = "http://localhost:8080/link"
)

var (
	clientID     string
	clientSecret string

	fitbitClient *fitbit.Client
	state        string
	codeVerifier string
	cfg          config

	profile *fitbit.Profile
)

type config struct {
	RedirectURI  string
	ClientID     string
	ClientSecret string

	UserID string
	Token  *fitbit.Token

	waitForAuth chan bool
}

func updateTokenFunc(oldToken, newToken *fitbit.Token) error {
	if oldToken == newToken {
		return nil
	}

	log.Print("Updating token.")

	cfg.Token = newToken

	return cfg.saveConfig()
}

func (cfg *config) setDefaults() {
	cfg.RedirectURI = redirectURI
	cfg.ClientID = clientID
	cfg.ClientSecret = clientSecret
	cfg.waitForAuth = make(chan bool, 1)
}

func (cfg *config) loadConfig() error {
	cfg.setDefaults()

	if _, err := os.Stat(fitbitConfig); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	file, err := os.Open(fitbitConfig)
	if err != nil {
		return fmt.Errorf("could not open '%s': %w", fitbitConfig, err)
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return fmt.Errorf("could not parse '%s': %w", fitbitConfig, err)
	}

	return nil
}

func (cfg *config) saveConfig() error {
	file, err := os.Create(fitbitConfig)
	if err != nil {
		return fmt.Errorf("could not create '%s': %w", fitbitConfig, err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("could not encode '%s': %w", fitbitConfig, err)
	}

	return nil
}

func (cfg *config) Valid() bool {
	if cfg.UserID == "" || cfg.Token == nil || cfg.Token.RefreshToken == "" {
		return false
	}

	p, _, _, err := fitbitClient.GetProfile(context.Background(), cfg.UserID, cfg.Token)
	if err != nil {
		return false
	}

	profile = p

	return true
}

func main() {
	if err := cfg.loadConfig(); err != nil {
		log.Fatal("could not load config: ", err)
	}

	fitbitClient = fitbit.NewClient(cfg.ClientID, cfg.ClientSecret, fitbit.ServerApplication, &fitbit.Scope{
		Activity: true,
		Weight:   true,
		Profile:  true,
	})
	fitbitClient.SetUpdateTokenFunc(updateTokenFunc)
	fitbitClient.EnableDebugMode()

	if !cfg.Valid() {
		getToken()
	}

	log.Print("Fetching Fitbit information...")

	fmt.Println("Information for:", profile.FullName)

	end := time.Now()
	start := end.AddDate(0, 0, -7)

	t := table.New(os.Stdout)
	t.SetHeaders("Date", "Steps", "Distance", "Sedentary", "Light", "Fair", "Very")

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		act, _, _, err := fitbitClient.GetDailyActivitySummary(context.Background(), cfg.UserID, d, cfg.Token)
		if err != nil {
			log.Fatal(err)
		}

		t.AddRow(
			d.Format("2006-01-02"),
			strconv.FormatInt(act.Summary.Steps, 10),
			strconv.FormatFloat(sum(act.Summary.Distances), 'g', 2, 64)+" km",
			strconv.FormatInt(act.Summary.SedentaryMinutes, 10)+" min",
			strconv.FormatInt(act.Summary.LightlyActiveMinutes, 10)+" min",
			strconv.FormatInt(act.Summary.FairlyActiveMinutes, 10)+" min",
			strconv.FormatInt(act.Summary.VeryActiveMinutes, 10)+" min",
		)
	}

	fmt.Println("Fitbit information:")
	t.Render()
}

func sum(s []fitbit.Distance) float64 {
	var sum float64

	for _, v := range s {
		sum += v.Distance
	}

	return sum
}

func getToken() {
	go configureServer()

	var authCodeURL *url.URL

	authCodeURL, state, codeVerifier = fitbitClient.AuthCodeURL(redirectURI)

	fmt.Println("Please open the following link in your browser, if it does not open automatically:")
	fmt.Println(authCodeURL.String())

	if err := open.Run(authCodeURL.String()); err != nil {
		fmt.Println("Could not open browser:", err)
	}

	fmt.Println("Allow this application access, then continue here.")

	<-cfg.waitForAuth
}

func configureServer() {
	http.HandleFunc("/link", linkFunc)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func linkFunc(w http.ResponseWriter, req *http.Request) {
	defer func() { cfg.waitForAuth <- true }()

	requestQuery := req.URL.Query()
	if requestQuery.Has("error") {
		http.Error(w, requestQuery.Get("error"), http.StatusInternalServerError)
		return
	}

	if requestQuery.Get("state") != state {
		http.Error(w, "state mismatched.", http.StatusBadRequest)
		return
	}

	linkResp, err := fitbitClient.Link(req.Context(), requestQuery.Get("code"), codeVerifier, redirectURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cfg.UserID = linkResp.UserID
	cfg.Token = linkResp.Token

	if err := cfg.saveConfig(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.WriteString(w, "You should close this site."); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
