package database

import (
	"math"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

type MapPoint struct {
	Lat           float64       // The latitude of the point
	Lng           float64       // The longitude of the point
	Distance      float64       // The distance from the previous point
	TotalDistance float64       // The total distance of the workout up to this point
	Duration      time.Duration // The duration from the previous point
	TotalDuration time.Duration // The total duration of the workout up to this point
	Time          time.Time     // The time the point was recorded

	ExtraMetrics ExtraMetrics // Extra metrics at this point
}

func (m *MapPoint) AverageSpeed() float64 {
	return m.Distance / m.Duration.Seconds()
}

func calcDistance(p1, p2 []MapPoint) float64 {
	switch {
	case len(p1) == 0 && len(p2) == 0:
		return 0.0
	case len(p1) == 0:
		return p2[len(p2)-1].TotalDistance
	case len(p2) == 0:
		return p1[len(p1)-1].TotalDistance
	}

	score := 0.0
	d1 := p1[len(p1)-1].TotalDistance
	d2 := p2[len(p2)-1].TotalDistance

	for _, p := range p1 {
		score += nearestPointScore(p, p2)
	}

	for _, p := range p2 {
		score += nearestPointScore(p, p1)
	}

	return score * math.Abs(d1-d2)
}

func nearestPointScore(p1 MapPoint, points []MapPoint) float64 {
	if len(points) == 0 {
		return 0
	}

	bestScore := gpx.HaversineDistance(p1.Lat, p1.Lng, points[0].Lat, points[0].Lng)

	for _, p := range points {
		score := gpx.HaversineDistance(p1.Lat, p1.Lng, p.Lat, p.Lng)

		if score < bestScore {
			bestScore = score
		}
	}

	return bestScore
}
