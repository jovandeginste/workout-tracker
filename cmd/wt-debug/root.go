package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func (c *cli) rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wt-debug",
		Short: "Workout Tracker debugger is a CLI tool to debug the workout tracker database",
	}

	cmd.AddCommand(c.versionCmd())
	cmd.AddCommand(c.configCmd())
	cmd.AddCommand(c.workoutsCmd())

	return cmd
}

func (c *cli) versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Shows the version of the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version: " + c.app.Version.PrettyVersion())
			fmt.Println("Build time: " + c.app.Version.BuildTime)
		},
	}
}

func (c *cli) configCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Show the full configuration of the application (including sensitive information)",
		Run: func(cmd *cobra.Command, args []string) {
			spew.Dump(c.app.Config)
		},
	}
}
