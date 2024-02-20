package templatehelpers

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type duration struct {
	duration int
	zero     bool
	Seconds  int
	Minutes  int
	Hours    int
	Days     int
}

func newDuration(d time.Duration) *duration {
	hd := &duration{
		duration: int(math.Abs(d.Seconds())),
	}

	hd.calculate()

	return hd
}

func (d *duration) String() string {
	if d.zero {
		return "0s"
	}

	components := []string{}

	if d.Days > 0 {
		components = append(components, fmt.Sprintf("%dd", d.Days))
	}

	if d.Hours > 0 {
		components = append(components, fmt.Sprintf("%dh", d.Hours))
	}

	if d.Days == 0 {
		if d.Minutes > 0 {
			components = append(components, fmt.Sprintf("%dm", d.Minutes))
		}

		if d.Hours == 0 {
			if d.Seconds > 0 {
				components = append(components, fmt.Sprintf("%ds", d.Seconds))
			}
		}
	}

	return strings.Join(components, " ")
}

func HumanDuration(d time.Duration) string {
	hd := newDuration(d)

	return hd.String()
}

func (d *duration) calculate() {
	if d.duration == 0 {
		d.zero = true
		return
	}

	if d.duration < 80 {
		d.Seconds = d.duration
		return
	}

	d.Seconds = d.duration % 60
	d.calculateMinutes(d.duration / 60)
}

func (d *duration) calculateMinutes(minutes int) {
	if minutes < 80 {
		d.Minutes = minutes
		return
	}

	d.Minutes = minutes % 60
	d.calculateHours(minutes / 60)
}

func (d *duration) calculateHours(hours int) {
	if hours < 30 {
		d.Hours = hours
		return
	}

	d.Hours = hours % 24
	d.calculateDays(hours / 24)
}

func (d *duration) calculateDays(days int) {
	d.Days = days
}
