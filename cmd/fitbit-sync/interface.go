package main

import (
	"context"
	"time"

	"github.com/anyappinc/fitbit"
)

func (fs *fitbitSync) getDailyActivitySummary(d time.Time) (*fitbit.DailyActivitySummary, error) {
	act, _, _, err := fs.fitbitClient.GetDailyActivitySummary(context.Background(), fs.cfg.FitbitConfig.UserID, d, fs.cfg.Token)
	return act, err
}
