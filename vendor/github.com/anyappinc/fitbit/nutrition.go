package fitbit

import (
	"context"
	"encoding/json"
	"time"
)

type (
	// WaterLog represents a user's water log.
	WaterLog struct {
		Amount float64 `json:"amount"`
		LogID  int64   `json:"logId"`
	}

	rawWater struct {
		Summary struct {
			Water float64 `json:"water"`
		} `json:"summary"`
		Water []WaterLog `json:"water"`
	}

	// Water represents a summary and list of a user's water log entries.
	Water struct {
		Total float64
		Logs  []WaterLog
	}
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (w *Water) UnmarshalJSON(b []byte) error {
	var raw rawWater
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	w.Total = raw.Summary.Water
	w.Logs = raw.Water
	return nil
}

// GetWater retrieves a summary and list of a user's water log entries for a given day.
//
// Scope.Nutrition is required.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/nutrition/get-water-log/
func (c *Client) GetWater(ctx context.Context, userID string, date time.Time, token *Token) (*Water, *RateLimit, []byte, error) {
	endpoint := c.getEndpoint("GetWater", userID, date.Format(dateFormat))
	b, rateLimit, err := c.getRequest(ctx, token, endpoint)
	if err != nil {
		return nil, nil, b, err
	}
	var water Water
	if err := json.Unmarshal(b, &water); err != nil {
		return nil, rateLimit, b, err
	}
	return &water, rateLimit, b, nil
}
