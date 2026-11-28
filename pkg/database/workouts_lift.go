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
	liftMinSpeed        = 1.8
	liftMaxSpeed        = 15.0
	liftMinGrade        = 0.08
	// liftMaxReverse is how far the elevation may drop (for an ascent) or rise
	// (for a descent) before the movement is no longer considered one lift.
	liftMaxReverse = 10.0
	// liftMinMovingRatio is the fraction of the segment that must be spent
	// moving. Lifts run non-stop; ski runs are full of short stops.
	liftMinMovingRatio = 0.75
	// liftMovingSpeed is the speed (in m/s) above which a point counts as moving.
	liftMovingSpeed = 1.0
	// liftMaxSpeedVariation is the maximum coefficient of variation of the
	// speed while moving. A lift has a near-constant speed.
	liftMaxSpeedVariation = 0.35
	// liftGainWindow is the window in which the elevation change must be
	// spread out. A lift climbs continuously; a barometer that sticks and then
	// jumps does not.
	liftGainWindow = 15 * time.Second
	// liftMaxGainConcentration is the maximum share of the total gain that may
	// happen within liftGainWindow. Real lifts stay well below this; elevation
	// jumps exceed it.
	liftMaxGainConcentration = 0.35
)

// DetectLiftSegments finds likely lift/uplift sections: a sustained,
// mostly straight ascent or descent at a mechanical pace.
func DetectLiftSegments(points []MapPoint) []Segment {
	lifts := []Segment{}

	for start := 0; start < len(points)-1; {
		run, ok := liftRun(points, start)
		if !ok {
			start++
			continue
		}

		segment, ok := newLiftSegment(points, run.start, run.end)
		if !ok {
			start++
			continue
		}

		segment.Index = len(lifts) + 1
		lifts = append(lifts, segment)
		start = run.end
	}

	return lifts
}

type pointRange struct {
	start, end int
}

// liftRun returns the next stretch of points that moves in one direction. The
// run ends at the furthest point in that direction once the elevation reverses
// by more than liftMaxReverse.
func liftRun(points []MapPoint, start int) (pointRange, bool) {
	direction := 0
	extreme := start
	startTime := points[start].TotalDuration

	for i := start + 1; i < len(points); i++ {
		elapsed := points[i].TotalDuration - startTime
		if elapsed > liftMaxDuration {
			break
		}

		if direction == 0 {
			elevationChange := points[i].Elevation - points[start].Elevation
			if math.Abs(elevationChange) < 5 {
				continue
			}

			direction = 1
			if elevationChange < 0 {
				direction = -1
			}

			extreme = i

			continue
		}

		progress := float64(direction) * (points[i].Elevation - points[extreme].Elevation)
		if progress > 0 {
			extreme = i
			continue
		}

		if -progress >= liftMaxReverse {
			return trimLiftRun(points, pointRange{start: start, end: extreme}), extreme > start
		}
	}

	return trimLiftRun(points, pointRange{start: start, end: extreme}), extreme > start
}

// trimLiftRun drops the standing still parts at both ends of a run. Waiting at
// the bottom station or stopping at the top is not part of the ride.
func trimLiftRun(points []MapPoint, run pointRange) pointRange {
	for run.start < run.end && !liftPointMoves(points[run.start+1]) {
		run.start++
	}

	for run.end > run.start && !liftPointMoves(points[run.end]) {
		run.end--
	}

	return run
}

func liftPointMoves(point MapPoint) bool {
	if point.Duration <= 0 {
		return false
	}

	return point.Distance2D/point.Duration.Seconds() > liftMovingSpeed
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

	grade := gain / length
	if grade < liftMinGrade {
		return Segment{}, false
	}

	beeline := beelineDistance(startPoint, endPoint)
	if beeline == 0 || length/beeline > liftMaxStraightness {
		return Segment{}, false
	}

	movingRatio, speedVariation := liftMovementStats(points[start : end+1])
	if movingRatio < liftMinMovingRatio || speedVariation > liftMaxSpeedVariation {
		return Segment{}, false
	}

	if liftGainConcentration(points, start, end) > liftMaxGainConcentration {
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

// liftMovementStats returns how much of the segment was spent moving and how
// much the moving speed varied. Lifts move almost continuously at a steady
// pace; skiing up or down does not.
func liftMovementStats(points []MapPoint) (float64, float64) {
	var (
		moving     int
		count      int
		sum, sumSq float64
	)

	for i := 1; i < len(points); i++ {
		if points[i].Duration <= 0 {
			continue
		}

		count++

		speed := points[i].Distance2D / points[i].Duration.Seconds()
		if speed <= liftMovingSpeed {
			continue
		}

		moving++
		sum += speed
		sumSq += speed * speed
	}

	if count == 0 || moving == 0 {
		return 0, math.Inf(1)
	}

	mean := sum / float64(moving)
	variance := max(0, sumSq/float64(moving)-mean*mean)

	return float64(moving) / float64(count), math.Sqrt(variance) / mean
}

// liftGainConcentration returns the largest share of the segment's total gain
// that happens within liftGainWindow. A lift climbs continuously, so this
// stays small; a stuck altitude that jumps does not.
func liftGainConcentration(points []MapPoint, start, end int) float64 {
	totalGain := math.Abs(points[end].Elevation - points[start].Elevation)
	if totalGain == 0 {
		return 0
	}

	maxWindowGain := 0.0
	windowStart := start

	for i := start; i <= end; i++ {
		for windowStart < i && points[i].TotalDuration-points[windowStart].TotalDuration > liftGainWindow {
			windowStart++
		}

		minElevation, maxElevation := points[windowStart].Elevation, points[windowStart].Elevation

		for j := windowStart; j <= i; j++ {
			minElevation = min(minElevation, points[j].Elevation)
			maxElevation = max(maxElevation, points[j].Elevation)
		}

		maxWindowGain = max(maxWindowGain, maxElevation-minElevation)
	}

	return maxWindowGain / totalGain
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
