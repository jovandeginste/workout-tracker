package database

import (
	"math"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

const (
	liftMinDuration     = 90 * time.Second
	liftMaxDuration     = 30 * time.Minute
	liftMinElevation    = 80.0
	liftMinLength       = 100.0
	liftMaxStraightness = 1.3
	liftMinSpeed        = 1.0
	liftMaxSpeed        = 15.0
)

// DetectLiftSegments finds likely lift/uplift sections: a sustained,
// mostly straight ascent or descent at a mechanical pace.
func DetectLiftSegments(points []MapPoint) []Segment {
	lifts := []Segment{}

	for start := 0; start < len(points)-1; {
		end, ok := liftCandidateEnd(points, start)
		if !ok {
			start++
			continue
		}

		segment, ok := newLiftSegment(points, start, end)
		if !ok {
			start++
			continue
		}

		segment.Index = len(lifts) + 1
		lifts = append(lifts, segment)
		start = end
	}

	return lifts
}

func liftCandidateEnd(points []MapPoint, start int) (int, bool) {
	direction := 0
	startElevation := points[start].Elevation
	startTime := points[start].TotalDuration

	end := start

	for i := start + 1; i < len(points); i++ {
		elapsed := points[i].TotalDuration - startTime
		if elapsed > liftMaxDuration {
			break
		}

		elevationChange := points[i].Elevation - startElevation
		if direction == 0 {
			if math.Abs(elevationChange) < 5 {
				continue
			}

			direction = 1
			if elevationChange < 0 {
				direction = -1
			}
		}

		if elevationChange*float64(direction) < -20 {
			break
		}

		pathLength := points[i].TotalDistance2D - points[start].TotalDistance2D
		if pathLength > liftMinLength {
			beeline := beelineDistance(points[start], points[i])
			if beeline > 0 && pathLength/beeline > liftMaxStraightness {
				break
			}
		}

		if elapsed >= liftMinDuration && math.Abs(elevationChange) >= liftMinElevation {
			end = i
		}
	}

	return end, end > start
}

func newLiftSegment(points []MapPoint, start, end int) (Segment, bool) {
	startPoint := points[start]
	endPoint := points[end]
	duration := endPoint.TotalDuration - startPoint.TotalDuration
	length := endPoint.TotalDistance2D - startPoint.TotalDistance2D
	gain := math.Abs(endPoint.Elevation - startPoint.Elevation)

	if duration < liftMinDuration || duration > liftMaxDuration {
		return Segment{}, false
	}

	if length < liftMinLength || gain < liftMinElevation {
		return Segment{}, false
	}

	speed := length / duration.Seconds()
	if speed < liftMinSpeed || speed > liftMaxSpeed {
		return Segment{}, false
	}

	beeline := beelineDistance(startPoint, endPoint)
	if beeline == 0 || length/beeline > liftMaxStraightness {
		return Segment{}, false
	}

	return Segment{
		Type:     SlopeKindLift,
		StartIdx: start,
		EndIdx:   end,
		Start:    startPoint,
		End:      endPoint,
		Length:   length,
		Gain:     gain,
		Duration: duration,
		AvgSlope: gain / length,
		Category: CategoryUncategorized,
	}, true
}

func liftPointMask(points []MapPoint, lifts []Segment) []bool {
	mask := make([]bool, len(points))

	for _, lift := range lifts {
		for i := max(0, lift.StartIdx); i <= min(len(points)-1, lift.EndIdx); i++ {
			mask[i] = true
		}
	}

	return mask
}

func beelineDistance(a, b MapPoint) float64 {
	return distance2DBetween(
		gpx.GPXPoint{Point: gpx.Point{Latitude: a.Lat, Longitude: a.Lng}},
		gpx.GPXPoint{Point: gpx.Point{Latitude: b.Lat, Longitude: b.Lng}},
	)
}
