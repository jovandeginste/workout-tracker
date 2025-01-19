package main

import (
	"os"
	"strconv"

	"github.com/aquasecurity/table"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/spf13/cobra"
)

func (c *cli) workoutsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workouts",
		Short: "Operate on workouts",
	}

	cmd.AddCommand(c.workoutsListCmd())
	cmd.AddCommand(c.workoutsShowCmd())

	return cmd
}

func (c *cli) workoutsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all workouts",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := table.New(os.Stdout)
			t.SetHeaders("ID", "Date", "Name")

			workouts, err := database.GetWorkouts(c.getDatabase())
			if err != nil {
				return err
			}

			for _, wo := range workouts {
				t.AddRow(strconv.FormatUint(uint64(wo.ID), 10), wo.Date.String(), wo.Name)
			}

			t.Render()

			return nil
		},
	}
}

func (c *cli) workoutsShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show information about a workout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			t := table.New(os.Stdout)
			t.SetRowLines(false)

			wo, err := database.GetWorkout(c.getDatabase(), id)
			if err != nil {
				return err
			}

			t.AddRow("ID", strconv.FormatUint(uint64(wo.ID), 10))
			t.AddRow("Date", wo.Date.String())
			t.AddRow("Name", wo.Name)

			t.AddRow("Location", wo.Address())
			t.AddRow("Distance (m)", strconv.FormatFloat(wo.TotalDistance(), 'f', 2, 64))
			t.AddRow("Duration (s)", strconv.FormatFloat(wo.TotalDuration().Seconds(), 'f', 2, 64))

			t.Render()

			return nil
		},
	}
}
