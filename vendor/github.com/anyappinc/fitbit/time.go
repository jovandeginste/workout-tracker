package fitbit

import "time"

func parseTime(layout, value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}
	t, err := time.Parse(layout, value)
	return &t, err
}

func timeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func timeRef(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}
