package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapPoint_IsMoving(t *testing.T) {
	tests := []struct {
		name     string
		distance float64
		duration time.Duration
		want     bool
	}{
		{name: "standing still", distance: 0, duration: time.Minute, want: false},
		{name: "very slow", distance: 10, duration: time.Minute, want: false},
		{name: "walking", distance: 50, duration: time.Minute, want: true},
		{name: "cycling", distance: 500, duration: time.Minute, want: true},
		{name: "no duration", distance: 0, duration: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MapPoint{Distance: tt.distance, Duration: tt.duration}
			assert.Equal(t, tt.want, p.IsMoving())
		})
	}
}
