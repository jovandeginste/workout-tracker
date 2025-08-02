package main

import (
	"fmt"
	"os"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
	"github.com/spf13/cobra"
)

func (c *cli) filesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "files",
		Short: "Operate on files",
	}

	cmd.AddCommand(c.workoutsParseCmd())
	cmd.AddCommand(c.workoutsCalculateCmd())

	return cmd
}

func (c *cli) workoutsParseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "parse",
		Short: "Parse a file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]
			content, err := os.ReadFile(filename)
			if err != nil {
				return err
			}

			wo, err := converters.ParseCollection(filename, content)
			if err != nil {
				return err
			}

			fmt.Println("Parsing was successful!")

			for _, f := range wo {
				fmt.Printf("- name: %s\n", f.Data.Name)

				if f.IsGPXBAsed() {
					fmt.Printf("  number of tracks: %d\n", len(f.GPX.Tracks))
					continue
				}

				fmt.Printf("  repetitions: %d\n", f.Data.TotalRepetitions)
				fmt.Printf("  duration: %s\n", f.Data.TotalDuration)
				fmt.Printf("  start: %s\n", f.Data.Start)
			}

			return nil
		},
	}
}

func (c *cli) workoutsCalculateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "calculate",
		Short: "Calculatet the information of a file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dst := "km"
			dstF := templatehelpers.HumanDistanceFor(dst)

			filename := args[0]
			content, err := os.ReadFile(filename)
			if err != nil {
				return err
			}

			workouts, err := database.NewWorkout(database.AnonymousUser(), database.WorkoutTypeAutoDetect, "", filename, content)
			if err != nil {
				return err
			}

			for _, workout := range workouts {
				fmt.Println("Parsing was successful!")
				fmt.Printf("- name: %s\n", workout.Name)
				fmt.Printf("- center: (%.5f, %.5f): %s\n", workout.Data.Center.Lat, workout.Data.Center.Lng, workout.Data.AddressString)
				fmt.Printf("- total distance: %.0fm (%s %s)\n", workout.TotalDistance(), dstF(workout.TotalDistance()), dst)
				fmt.Printf("- total duration: %.0fs (%s)\n", workout.TotalDuration().Seconds(), workout.TotalDuration().String())

				p := workout.Data.Details.Points
				lp := p[len(p)-1]
				fmt.Printf("- last point total distance: %sm (%s %s)\n", templatehelpers.RoundFloat64(lp.TotalDistance), dstF(lp.TotalDistance), dst)
			}

			return nil
		},
	}
}
