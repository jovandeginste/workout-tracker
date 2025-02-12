package fitbit

import (
	"net/http"
	"strconv"
	"time"
)

type (
	// RateLimit represents the rate limit of API calls.
	//
	// Note: The rate limit headers are approximate and asynchronously updated.
	RateLimit struct {
		Quota     int64
		Remaining int64
		ResetTime *time.Time
	}
)

func extractRateLimit(h *http.Header) *RateLimit {
	quataString := h.Get("Fitbit-Rate-Limit-Limit")
	remainingString := h.Get("Fitbit-Rate-Limit-Remaining")
	resetString := h.Get("Fitbit-Rate-Limit-Reset")
	if quataString == "" || remainingString == "" || resetString == "" {
		return nil
	}
	quata, _ := strconv.ParseInt(quataString, 10, 64)
	remaining, _ := strconv.ParseInt(remainingString, 10, 64)
	reset, _ := strconv.ParseInt(resetString, 10, 64)
	return &RateLimit{
		Quota:     quata,
		Remaining: remaining,
		ResetTime: timeRef(time.Now().Add(time.Duration(reset * 1e9))),
	}
}
