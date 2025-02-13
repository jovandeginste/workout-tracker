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

func (fs *fitbitSync) getActivityTCX(a fitbit.Activity) ([]byte, error) {
	tcx, _, _, err := fs.fitbitClient.GetActivityTCX(context.Background(), fs.cfg.FitbitConfig.UserID, a.LogID, fs.cfg.Token)
	return tcx, err
}
