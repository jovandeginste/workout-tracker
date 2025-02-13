package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/anyappinc/fitbit"
	"github.com/aquasecurity/table"
)

func (fs *fitbitSync) showProfile() {
	units := fs.fitbitClient.GetUnit()
	log.Print("Fetching Fitbit information...")

	fmt.Println("Information for:", fs.profile.FullName)
	fmt.Printf("- weight: %.2f %s\n", fs.profile.Weight, units.Weight)
	fmt.Printf("- height: %.2f %s\n", fs.profile.Height, units.Height)
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
				fmt.Sprintf("%.2f %s", a.Distance, units.Distance),
			)
		}

		if act.Summary == nil {
			continue
		}

		summaries.AddRow(
			d.Format("2006-01-02"),
			strconv.FormatInt(act.Summary.Steps, 10),
			strconv.FormatFloat(sum(act.Summary.Distances), 'g', 2, 64)+" "+units.Distance,
			strconv.FormatInt(act.Summary.SedentaryMinutes, 10)+" min",
			strconv.FormatInt(act.Summary.LightlyActiveMinutes, 10)+" min",
			strconv.FormatInt(act.Summary.FairlyActiveMinutes, 10)+" min",
			strconv.FormatInt(act.Summary.VeryActiveMinutes, 10)+" min",
		)
	}

	fmt.Println("Daily summaries:")
	summaries.Render()
	fmt.Println()

	fmt.Println("Activities:")
	activities.Render()

	return nil
}

func sum(s []fitbit.Distance) float64 {
	var sum float64

	for _, v := range s {
		sum += v.Distance
	}

	return sum
}
