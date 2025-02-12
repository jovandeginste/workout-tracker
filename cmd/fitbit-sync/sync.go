package main

import (
	"errors"
	"log"
	"time"

	"github.com/anyappinc/fitbit"
	"github.com/jovandeginste/workout-tracker/v2/pkg/app"
	"resty.dev/v3"
)

func (fs *fitbitSync) initRESTClient() {
	client := resty.New()
	client.SetBaseURL(fs.WorkoutConfig.URL + "/api/v1/")
	client.SetAuthToken(fs.WorkoutConfig.APIKey)

	fs.restClient = client
}

func (fs *fitbitSync) syncActivities(days int) {
	fs.initRESTClient()

	end := time.Now()
	endDate := end.Format("2006-01-02")
	start := end.AddDate(0, 0, -days)

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		date := d.Format("2006-01-02")
		log.Printf("Syncing: %s\n", date)

		act, err := fs.getDailyActivitySummary(d)
		if err != nil {
			log.Printf("could not get daily activity summary: %v", err)
		}

		if act == nil {
			continue
		}

		final := date == endDate

		mw := fs.buildMeasurement(date, final, act)
		if err := fs.postMeasurement(mw); err != nil {
			log.Printf("could not post measurement: %v", err)
		}
	}
}

func (fs *fitbitSync) buildMeasurement(date string, final bool, act *fitbit.DailyActivitySummary) *app.Measurement {
	mw := &app.Measurement{
		Date: date,
	}

	if act.Summary != nil {
		mw.Steps = float64(act.Summary.Steps)
	}

	if !final {
		return mw
	}

	mw.Height = fs.profile.Height
	mw.HeightUnit = heightUnit(fs.profile.HeightUnit)
	mw.Weight = fs.profile.Weight
	mw.WeightUnit = weightUnit(fs.profile.WeightUnit)

	return mw
}

func (fs *fitbitSync) postMeasurement(m *app.Measurement) error {
	res, err := fs.restClient.R().
		SetBody(m).
		Post("/daily")
	if err != nil {
		return err
	}

	if !res.IsSuccess() {
		return errors.New(res.Status())
	}

	return nil
}
