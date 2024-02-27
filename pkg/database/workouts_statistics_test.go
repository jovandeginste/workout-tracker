package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsItem_CalcultateSpeed(t *testing.T) {
	si := &StatisticsItem{
		Distance: 1000,
		Duration: 60 * time.Second,
	}

	assert.Zero(t, si.Speed)
	si.CalcultateSpeed()

	assert.InDelta(t, 16.67, si.Speed, 0.01)
}

func TestStatisticsItem_CanHave(t *testing.T) {
	si := &StatisticsItem{
		TotalDistance: 1800,
		Kilometer:     2,
	}

	assert.True(t, si.CanHave(100))
	assert.True(t, si.CanHave(199))
	assert.False(t, si.CanHave(200))
	assert.False(t, si.CanHave(203))
}

func TestStatisticsItem_StatisticsPerKilometer(t *testing.T) {
	w := defaultWorkout(t)
	assert.NotNil(t, w)

	stats := w.StatisticsPerKilometer()
	assert.Len(t, stats, 4)

	assert.Equal(t, 1, stats[0].Kilometer)
	assert.Equal(t, 2, stats[1].Kilometer)
	assert.Equal(t, 3, stats[2].Kilometer)
	assert.Equal(t, 4, stats[3].Kilometer)

	assert.InDelta(t, 1000, stats[0].TotalDistance, 20)
	assert.InDelta(t, 2000, stats[1].TotalDistance, 20)
	assert.InDelta(t, 3000, stats[2].TotalDistance, 20)
	assert.InDelta(t, 3100, stats[3].TotalDistance, 20)

	assert.Equal(t, 4*time.Minute+2*time.Second, stats[0].Duration)
	assert.Equal(t, 5*time.Minute+36*time.Second, stats[1].Duration)
	assert.Equal(t, 5*time.Minute+39*time.Second, stats[2].Duration)
	assert.Equal(t, 31*time.Second, stats[3].Duration)
}
