// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datetime

import (
	"time"

	"github.com/muktihari/fit/profile/basetype"
)

var epoch = time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)

// Epoch returns FIT epoch (31 Dec 1989 00:00:000 UTC) as time.Time
func Epoch() time.Time { return epoch }

// ToUint32 converts t into uint32 FIT representative time value.
func ToUint32(t time.Time) uint32 {
	if t.Before(epoch) {
		return basetype.Uint32Invalid
	}
	return uint32(t.Sub(epoch).Seconds())
}

// ToTime converts uint32 value into time.Time.
func ToTime(value uint32) time.Time {
	if value == basetype.Uint32Invalid {
		return time.Time{}
	}
	return epoch.Add(time.Duration(value) * time.Second)
}

// ToLocalTime converts time in local time zone by specifying the time zone offset hours (+7 for GMT+7).
func ToLocalTime(t time.Time, tzOffsetHours int) time.Time {
	if t.IsZero() {
		return t
	}
	return t.In(time.FixedZone("", tzOffsetHours*60*60)) // Use unnamed fixed zones
}

// TzOffsetHours calculates time zone offset.
//
// formula ilustration: (activity.LocalTimestamp - activity.Timestamp) / 3600
func TzOffsetHours(localDateTime, dateTime time.Time) int {
	return int(localDateTime.Sub(dateTime).Seconds()) / 3600
}

// TzOffsetHoursFromUint32 is similar to TzOffsetHours but it took uint32 as parameters.
func TzOffsetHoursFromUint32(localDateTime, dateTime uint32) int {
	return int(localDateTime-dateTime) / 3600
}
