package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/anyappinc/fitbit"
	"github.com/spf13/cobra"
)

type fitbitSync struct {
	cfg         config
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

func (fs *fitbitSync) rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "fitbit-sync",
	}

	cmd.AddCommand(fs.initCmd())
	cmd.AddCommand(fs.showCmd())

	return cmd
}

func (fs *fitbitSync) initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the Fitbit oauth configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fs.loadConfig(); err != nil {
				return fmt.Errorf("could not load config: %w", err)
			}

			fs.initClient()

			if err := fs.getToken(); err != nil {
				return fmt.Errorf("could not get token: %w", err)
			}

			fmt.Println("Configuration saved. You should now be able to use the other commands.")

			return nil
		},
	}

	cmd.Flags().StringVarP(&fs.bind, "redirect-bind", "r", "http://localhost:8080", "IP and port to listen for redirect request")
	cmd.Flags().StringVarP(&fs.cfg.FitbitConfig.ClientID, "client-id", "i", "", "Fitbit client ID")
	cmd.Flags().StringVarP(&fs.cfg.FitbitConfig.ClientSecret, "client-secret", "s", "", "Fitbit client secret")
	cmd.Flags().StringVarP(&fs.cfg.FitbitConfig.UserID, "user-id", "u", "", "Fitbit user ID")

	for _, flag := range []string{"client-id", "client-secret", "user-id"} {
		if err := cmd.MarkFlagRequired(flag); err != nil {
			log.Fatal(err)
		}
	}

	return cmd
}

func (fs *fitbitSync) showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show the Fitbit activities",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fs.loadConfig(); err != nil {
				return fmt.Errorf("could not load config: %w", err)
			}

			if err := fs.reinitClient(); err != nil {
				return fmt.Errorf("could not initialize Fitbit client: %w", err)
			}

			log.Print("Fetching Fitbit information...")

			fs.showProfile()

			if err := fs.showActivities(); err != nil {
				return fmt.Errorf("could not show activities: %w", err)
			}

			return nil
		},
	}

	return cmd
}
