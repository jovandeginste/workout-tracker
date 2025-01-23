package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

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
	cmd.AddCommand(c.workoutsDiagCmd())
	cmd.AddCommand(c.workoutsShowCmd())
	cmd.AddCommand(c.workoutsExportCmd())
	cmd.AddCommand(c.workoutsImportCmd())

	return cmd
}

func (c *cli) workoutsDiagCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "diag",
		Short: "Perform diagnose on all workouts",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := table.New(os.Stdout)
			t.SetHeaders("ID", "Name", "Issues")

			var ids []uint64

			if err := c.getDatabase().Model(&database.Workout{}).Pluck("ID", &ids).Error; err != nil {
				return err
			}

			for _, id := range ids {
				issues := []string{}

				wo, err := database.GetWorkout(c.getDatabase(), id)
				if err != nil {
					issues = append(issues, err.Error())
				}

				if len(issues) == 0 {
					issues = []string{"OK"}
				}

				t.AddRow(strconv.FormatUint(id, 10), wo.Name, strings.Join(issues, "; "))
			}

			t.Render()

			return nil
		},
	}
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
				t.AddRow(strconv.FormatUint(wo.ID, 10), wo.Date.String(), wo.Name)
			}

			t.Render()

			return nil
		},
	}
}

func (c *cli) workoutsShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show workout-id",
		Short: "Show information about a workout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			t := table.New(os.Stdout)
			t.SetRowLines(false)

			wo, err := database.GetWorkout(c.getDatabase(), id)
			if err != nil {
				return err
			}

			t.AddRow("ID", strconv.FormatUint(wo.ID, 10))
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

func (c *cli) workoutsExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "export [workout-id...]",
		Short: "Export all, or some workouts to json (stdout)",
		RunE: func(cmd *cobra.Command, args []string) error {
			var ids []uint64

			if err := c.getDatabase().Model(&database.Workout{}).Pluck("ID", &ids).Error; err != nil {
				return err
			}

			filter := len(args) > 0

			for _, id := range ids {
				if filter && !slices.Contains(args, strconv.FormatUint(id, 10)) {
					continue
				}

				c.app.Logger().Info("Exporting workout", "id", id)

				wo, err := database.GetWorkoutDetails(c.getDatabase(), id)
				if err != nil {
					return err
				}

				e, err := wo.Export()
				if err != nil {
					return err
				}

				os.Stdout.Write(e)
			}

			return nil
		},
	}
}

func (c *cli) workoutsImportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "import <file>",
		Short: "Import workouts from a json file (eg. from an earlier export)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]
			c.app.Logger().Info(fmt.Sprintf("Importing from '%s'", filename))

			dat, err := os.ReadFile(filename)
			if err != nil {
				return err
			}

			lines := strings.Split(string(dat), "\n")

			for _, line := range lines {
				if line == "" {
					continue
				}

				var wo database.Workout

				if err := json.Unmarshal([]byte(line), &wo); err != nil {
					return err
				}

				c.app.Logger().Info("Importing workout", "id", wo.ID)

				if err := wo.Save(c.app.DB()); err != nil {
					return err
				}
			}

			return nil
		},
	}
}
