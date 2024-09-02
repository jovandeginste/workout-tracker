package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsItem_CalcultateSpeed(t *testing.T) {
	si := &BreakdownItem{
		Distance: 1000,
		Duration: 60 * time.Second,
	}

	assert.Zero(t, si.Speed)
	si.CalcultateSpeed()

	assert.InDelta(t, 16.67, si.Speed, 0.01)
}

func TestStatisticsItem_CanHave(t *testing.T) {
	si := &BreakdownItem{
		TotalDistance: 1800,
		Counter:       2,
	}

	assert.True(t, si.canHaveDistance(100, 2000))
	assert.True(t, si.canHaveDistance(199, 2000))
	assert.False(t, si.canHaveDistance(200, 2000))
	assert.False(t, si.canHaveDistance(203, 2000))
}

func TestStatisticsItem_StatisticsPerKilometer(t *testing.T) {
	w := defaultWorkout(t)
	assert.NotNil(t, w)

	stats, err := w.StatisticsPer(1, "km")
	assert.NoError(t, err)
	assert.Len(t, stats.Items, 4)

	assert.Equal(t, 1, stats.Items[0].Counter)
	assert.Equal(t, 2, stats.Items[1].Counter)
	assert.Equal(t, 4, stats.Items[3].Counter)

	assert.InDelta(t, 1000, stats.Items[0].TotalDistance, 20)
	assert.InDelta(t, 2000, stats.Items[1].TotalDistance, 20)
	assert.InDelta(t, 3100, stats.Items[3].TotalDistance, 20)

	assert.Equal(t, 242*time.Second, stats.Items[0].Duration)
	assert.Equal(t, 336*time.Second, stats.Items[1].Duration)
	assert.Equal(t, 31*time.Second, stats.Items[3].Duration)
}

func TestStatisticsItem_StatisticsPerMinute(t *testing.T) {
	w := defaultWorkout(t)
	assert.NotNil(t, w)

	stats, err := w.StatisticsPer(1, "min")
	assert.NoError(t, err)
	assert.Len(t, stats.Items, 16)

	assert.Equal(t, 1, stats.Items[0].Counter)
	assert.Equal(t, 3, stats.Items[2].Counter)
	assert.Equal(t, 8, stats.Items[7].Counter)
	assert.Equal(t, 16, stats.Items[15].Counter)

	assert.InDelta(t, 180, stats.Items[0].TotalDistance, 20)
	assert.InDelta(t, 700, stats.Items[2].TotalDistance, 20)
	assert.InDelta(t, 1765, stats.Items[7].TotalDistance, 20)
	assert.InDelta(t, 3100, stats.Items[15].TotalDistance, 20)

	assert.Equal(t, 56*time.Second, stats.Items[0].Duration)
	assert.Equal(t, 63*time.Second, stats.Items[1].Duration)
	assert.Equal(t, 106*time.Second, stats.Items[7].Duration)
	assert.Equal(t, 51*time.Second, stats.Items[15].Duration)

	assert.Equal(t, 56*time.Second, stats.Items[0].TotalDuration)
	assert.Equal(t, 119*time.Second, stats.Items[1].TotalDuration)
	assert.Equal(t, 509*time.Second, stats.Items[7].TotalDuration)
	assert.Equal(t, 948*time.Second, stats.Items[15].TotalDuration)
}
