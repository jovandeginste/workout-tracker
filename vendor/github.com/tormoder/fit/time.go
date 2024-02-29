package fit

import "time"

const (
	systemTimeMarker = 0x10000000
	localZoneName    = "FITLOCAL"
)

var timeBase = time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)

// IsBaseTime reports if t represents the FIT base time.
func IsBaseTime(t time.Time) bool {
	return t.Equal(timeBase)
}

func decodeDateTime(dt uint32) time.Time {
	return timeBase.Add(time.Duration(dt) * time.Second)
}

func encodeTime(t time.Time) uint32 {
	return uint32(t.Sub(timeBase) / time.Second)
}
