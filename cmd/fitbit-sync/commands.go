package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func (fs *fitbitSync) rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "fitbit-sync",
	}

	cmd.PersistentFlags().StringVarP(&fs.configDir, "config-dir", "c", configDir(), "Configuration directory")
	cmd.PersistentFlags().StringVarP(&fs.configFile, "config-file", "f", "fitbit.json", "Configuration file inside the directory")

	cmd.AddCommand(fs.initCmd())
	cmd.AddCommand(fs.showCmd())
	cmd.AddCommand(fs.syncCmd())

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
	var days int

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

			fs.showProfile()

			if err := fs.showActivities(days); err != nil {
				return fmt.Errorf("could not show activities: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().IntVarP(&days, "days", "d", 7, "Number of days to show")

	return cmd
}

func (fs *fitbitSync) syncCmd() *cobra.Command {
	var days int

	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync the Fitbit activities to your Workout Tracker",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fs.loadConfig(); err != nil {
				return fmt.Errorf("could not load config: %w", err)
			}

			fs.WorkoutConfig.CopyFrom(&fs.cfg.WorkoutConfig)

			if fs.WorkoutConfig.persist {
				fs.cfg.WorkoutConfig = fs.WorkoutConfig
				if err := fs.saveConfig(); err != nil {
					return fmt.Errorf("could not save config: %w", err)
				}
			}

			if err := fs.reinitClient(); err != nil {
				return fmt.Errorf("could not initialize Fitbit client: %w", err)
			}

			fs.showProfile()
			fs.syncActivities(days)

			return nil
		},
	}

	cmd.Flags().IntVarP(&days, "days", "d", 7, "Number of days to show")
	cmd.Flags().StringVarP(&fs.WorkoutConfig.URL, "url", "u", "", "Workout Tracker root URL")
	cmd.Flags().StringVarP(&fs.WorkoutConfig.APIKey, "key", "k", "", "Workout Tracker API key")
	cmd.Flags().BoolVarP(&fs.WorkoutConfig.persist, "persist", "p", false, "Persist Workout Tracker configuration")

	return cmd
}
