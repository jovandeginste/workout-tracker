package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anyappinc/fitbit"
	"github.com/aquasecurity/table"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
	"github.com/spf13/cast"
)

func (fs *fitbitSync) showProfile() {
	units := fs.fitbitClient.GetUnit()
	log.Print("Fetching Fitbit information...")

	fmt.Println("Information for:", fs.profile.FullName)
	fmt.Println("- weight:", templatehelpers.RoundFloat64(fs.profile.Weight), units.Weight)
	fmt.Println("- height:", templatehelpers.RoundFloat64(fs.profile.Height), units.Height)
	fmt.Println()
}

func (fs *fitbitSync) showActivities(days int) error {
	fs.fitbitClient.SetLocale("en-UK")
	units := fs.fitbitClient.GetUnit()
	end := time.Now()
	start := end.AddDate(0, 0, -days)

	summaries := table.New(os.Stdout)
	summaries.SetHeaders("Date", "Steps", "Distance", "Sedentary", "Light", "Fair", "Very")

	activities := table.New(os.Stdout)
	activities.SetHeaders("Date", "Type", "Duration", "Distance")

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		act, err := fs.getDailyActivitySummary(d)
		if err != nil {
			return err
		}

		for _, a := range act.Activities {
			if !a.HasStartTime {
				continue
			}

			activities.AddRow(
				a.StartDateTime.Format("2006-01-02T15:04"),
				a.Name,
				a.Duration.String(),
				templatehelpers.RoundFloat64(a.Distance)+" "+units.Distance,
			)
		}

		if act.Summary == nil {
			continue
		}

		summaries.AddRow(
			d.Format("2006-01-02"),
			cast.ToString(act.Summary.Steps),
			templatehelpers.RoundFloat64(findTotal(act.Summary.Distances))+" "+units.Distance,
			cast.ToString(act.Summary.SedentaryMinutes)+" min",
			cast.ToString(act.Summary.LightlyActiveMinutes)+" min",
			cast.ToString(act.Summary.FairlyActiveMinutes)+" min",
			cast.ToString(act.Summary.VeryActiveMinutes)+" min",
		)
	}

	fmt.Println("Daily summaries:")
	summaries.Render()
	fmt.Println()

	fmt.Println("Activities:")
	activities.Render()

	return nil
}

func findTotal(s []fitbit.Distance) float64 {
	for _, v := range s {
		if v.Activity == "total" {
			return v.Distance
		}
	}

	return 0
}
