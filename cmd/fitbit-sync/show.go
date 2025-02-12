package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/anyappinc/fitbit"
	"github.com/aquasecurity/table"
)

func (fs *fitbitSync) showProfile() {
	fmt.Println("Information for:", fs.profile.FullName)
	fmt.Printf("Weight: %.2f %s\n", fs.profile.Weight, weightUnit(fs.profile.WeightUnit))
	fmt.Printf("Height: %.2f %s\n", fs.profile.Height, heightUnit(fs.profile.HeightUnit))
}

func (fs *fitbitSync) showActivities() error {
	end := time.Now()
	start := end.AddDate(0, 0, -7)

	t := table.New(os.Stdout)
	t.SetHeaders("Date", "Steps", "Distance", "Sedentary", "Light", "Fair", "Very")

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		act, err := fs.getDailyActivitySummary(d)
		if err != nil {
			return err
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

	return nil
}

func sum(s []fitbit.Distance) float64 {
	var sum float64

	for _, v := range s {
		sum += v.Distance
	}

	return sum
}

func weightUnit(u string) string {
	switch u {
	case "METRIC":
		return "kg"
	default:
		return "lb"
	}
}

func heightUnit(u string) string {
	switch u {
	case "METRIC":
		return "cm"
	default:
		return "in"
	}
}
