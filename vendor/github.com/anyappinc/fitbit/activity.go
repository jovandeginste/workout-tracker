package fitbit

import (
	"context"
	"encoding/json"
	"time"
)

type (
	rawActivity struct {
		ActivityID           int64      `json:"activityId"`
		Name                 string     `json:"name"`
		Description          string     `json:"description"`
		DetailsLink          string     `json:"detailsLink"`
		Calories             float64    `json:"calories"`
		StartDate            string     `json:"startDate"`
		StartTime            string     `json:"startTime"`
		Duration             int64      `json:"duration"` // in milliseconds
		Distance             float64    `json:"distance"`
		Steps                int64      `json:"steps"`
		HasActiveZoneMinutes bool       `json:"hasActiveZoneMinutes"`
		HasStartTime         bool       `json:"hasStartTime"`
		IsFavorite           bool       `json:"isFavorite"`
		LastModified         *time.Time `json:"lastModified"`
		LogID                int64      `json:"logId"`
		ActivityParentID     int64      `json:"activityParentId"`
		ActivityParentName   string     `json:"activityParentName"`
	}

	// Activity represents a user's activity.
	Activity struct {
		ID                   int64
		Name                 string
		Description          string
		DetailsLink          string
		Calories             float64
		StartDateTime        *time.Time
		Duration             time.Duration
		Distance             float64
		Steps                int64
		HasActiveZoneMinutes bool
		HasStartTime         bool
		IsFavorite           bool
		LastModified         *time.Time
		LogID                int64
		ParentID             int64
		ParentName           string
	}

	// Goals represents a user's activity goals.
	Goals struct {
		ActiveMinutes int64   `json:"activeMinutes"`
		CaloriesOut   int64   `json:"caloriesOut"`
		Distance      float64 `json:"distance"`
		Floors        int64   `json:"floors"`
		Steps         int64   `json:"steps"`
	}

	// Distance represents a user's activity distance.
	Distance struct {
		Activity string  `json:"activity"`
		Distance float64 `json:"distance"`
	}

	// HeartRateZone represents a user's heart rate zone.
	HeartRateZone struct {
		Name        string  `json:"name"`
		CaloriesOut float64 `json:"caloriesOut"`
		Max         int64   `json:"max"`
		Min         int64   `json:"min"`
		Minutes     int64   `json:"minutes"`
	}

	// Summary represents a user's daily activity summary.
	Summary struct {
		ActiveScore            int64           `json:"activeScore"`
		ActivityCalories       int64           `json:"activityCalories"`
		CaloriesEstimationMu   int64           `json:"caloriesEstimationMu"`
		CaloriesBMR            int64           `json:"caloriesBMR"`
		CaloriesOut            int64           `json:"caloriesOut"`
		CaloriesOutUnestimated int64           `json:"caloriesOutUnestimated"`
		MarginalCalories       int64           `json:"marginalCalories"`
		Distances              []Distance      `json:"distances"`
		Elevation              float64         `json:"elevation"`
		Floors                 int64           `json:"floors"`
		Steps                  int64           `json:"steps"`
		HeartRateZones         []HeartRateZone `json:"heartRateZones"`
		RestingHeartRate       int64           `json:"restingHeartRate"`
		FairlyActiveMinutes    int64           `json:"fairlyActiveMinutes"`
		LightlyActiveMinutes   int64           `json:"lightlyActiveMinutes"`
		SedentaryMinutes       int64           `json:"sedentaryMinutes"`
		VeryActiveMinutes      int64           `json:"veryActiveMinutes"`
		UseEstimation          bool            `json:"useEstimation"`
	}

	// DailyActivitySummary represents a summary and list of a user’s
	// activities and activity log entries.
	DailyActivitySummary struct {
		Activities []Activity `json:"activities"`
		Goals      *Goals     `json:"goals"`
		Summary    *Summary   `json:"summary"`
	}
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Activity) UnmarshalJSON(b []byte) error {
	var raw rawActivity
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	startDateTime, err := parseTime("2006-01-0215:04", raw.StartDate+raw.StartTime)
	if err != nil {
		return err
	}

	a.ID = raw.ActivityID
	a.Name = raw.Name
	a.Description = raw.Description
	a.DetailsLink = raw.DetailsLink
	a.Calories = raw.Calories
	a.StartDateTime = startDateTime
	a.Duration = time.Duration(raw.Distance * 1000 * 1000) // convert milliseconds to nanoseconds.
	a.Distance = raw.Distance
	a.Steps = raw.Steps
	a.HasActiveZoneMinutes = raw.HasActiveZoneMinutes
	a.HasStartTime = raw.HasStartTime
	a.IsFavorite = raw.IsFavorite
	a.LastModified = raw.LastModified
	a.LogID = raw.LogID
	a.ParentID = raw.ActivityParentID
	a.ParentName = raw.ActivityParentName
	return nil
}

// GetDailyActivitySummary retrieves a summary and list of a user’s activities and activity log entries for a given day.
//
// Scope.Activity is required.
//
// Scope.HeartRate is required to obtain `DailyActivitySummary.Summary.HeartRateZones`
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/activity/get-daily-activity-summary/
func (c *Client) GetDailyActivitySummary(ctx context.Context, userID string, date time.Time, token *Token) (*DailyActivitySummary, *RateLimit, []byte, error) {
	endpoint := c.getEndpoint("GetDailyActivitySummary", userID, date.Format(dateFormat))
	b, rateLimit, err := c.getRequest(ctx, token, endpoint)
	if err != nil {
		return nil, nil, b, err
	}
	var dailyActivitySummary DailyActivitySummary
	if err := json.Unmarshal(b, &dailyActivitySummary); err != nil {
		return nil, rateLimit, b, err
	}
	return &dailyActivitySummary, rateLimit, b, nil
}
