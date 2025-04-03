package database

import (
	"math"
	"time"
)

// MaxDeltaMeter is the maximum distance in meters that a point can be away from
// the route segment
const MaxDeltaMeter = 20.0

// MaxTotalDistanceFraction is the maximum percentage of the total distance of
// the route segment that can be exceeded by the total distance matching part of
// the route (1.0 = 100%)
const MaxTotalDistanceFraction = 0.9

// RouteSegmentMatch is a match between a route segment and a workout
type RouteSegmentMatch struct {
	Workout      *Workout      `json:"workout"`
	RouteSegment *RouteSegment `json:"routeSegment"`

	first, last MapPoint // The first and last point of the route
	end         MapPoint // The last point of the workout

	RouteSegmentID uint64        `gorm:"primaryKey" json:"routeSegmentID"` // The ID of the route segment
	WorkoutID      uint64        `gorm:"primaryKey" json:"workoutID"`      // The ID of the workout
	FirstID        int           `json:"firstID"`                          // The index of the first point of the route
	LastID         int           `json:"lastID"`                           // The index of the last point of the route
	Distance       float64       `json:"distance"`                         // The total distance of the route segment for this workout
	Duration       time.Duration `json:"duration"`                         // The total duration of the route segment for this workout
}

func (rsm *RouteSegmentMatch) AverageSpeed() float64 {
	return rsm.Distance / rsm.Duration.Seconds()
}

// NewRouteSegmentMatch will create a new route segment match from a workout and
// the first and last point of the route along the route segment
func (rs *RouteSegment) NewRouteSegmentMatch(workout *Workout, p, last int) *RouteSegmentMatch {
	rsm := &RouteSegmentMatch{
		Workout:      workout,
		RouteSegment: rs,
		FirstID:      p,
		LastID:       last,
	}

	rsm.calculate()

	return rsm
}

// IsBetterThan returns true if the new route segment match is better than the
// current one
func (rsm *RouteSegmentMatch) IsBetterThan(current *RouteSegmentMatch) bool {
	return current == nil || rsm.Distance < current.Distance
}

// MatchesDistance returns true if the distance of the route segment match is
// within MaxTotalDistancePercentage of the distance of the current route
// segment
func (rsm *RouteSegmentMatch) MatchesDistance(distance float64) bool {
	return math.Abs(rsm.Distance/distance) > MaxTotalDistanceFraction
}

// calculate will calculate the total distance and duration of the route
// segment, and the total number of points of this workout along the route
// segment
func (rsm *RouteSegmentMatch) calculate() {
	rsm.RouteSegmentID = rsm.RouteSegment.ID
	rsm.WorkoutID = rsm.Workout.ID
	rsm.first = rsm.Workout.Data.Details.Points[rsm.FirstID]
	rsm.last = rsm.Workout.Data.Details.Points[rsm.LastID]
	rsm.end = rsm.Workout.Data.Details.Points[len(rsm.Workout.Data.Details.Points)-1]

	if rsm.FirstID <= rsm.LastID {
		rsm.Distance = rsm.last.TotalDistance - rsm.first.TotalDistance
		rsm.Duration = rsm.last.TotalDuration - rsm.first.TotalDuration
	} else {
		rsm.Distance = rsm.last.TotalDistance + rsm.end.TotalDistance - rsm.first.TotalDistance
		rsm.Duration = rsm.last.TotalDuration + rsm.end.TotalDuration - rsm.first.TotalDuration
	}
}

// FindMatches will find all workouts that match the current route segment
// The result will contain a list of RouteSegmentMatches, which will contain
// the workout, the point of the workout along the segment, and the total
// distance and duration of the segment for this workout.
func (rs *RouteSegment) FindMatches(workouts []*Workout) []*RouteSegmentMatch {
	if len(rs.Points) == 0 {
		return nil
	}

	var result []*RouteSegmentMatch

	for _, w := range workouts {
		if m := rs.Match(w); m != nil {
			result = append(result, m)
		}
	}

	return result
}

// Match will find the best match (if any) of the route segment in the workout
// First calculate all possible starting points, then find the best one that
// actually matches the segment.
func (rs *RouteSegment) Match(workout *Workout) *RouteSegmentMatch {
	if !workout.Type.IsLocation() {
		return nil
	}

	if !workout.HasTracks() {
		return nil
	}

	sp := rs.StartingPoints(workout.Data.Details.Points)
	if len(sp) == 0 {
		return nil
	}

	var bestMatch *RouteSegmentMatch

	for _, p := range sp {
		if last, ok := rs.MatchSegment(workout, p, true); ok {
			rsm := rs.NewRouteSegmentMatch(workout, p, last)
			if rsm.MatchesDistance(rs.TotalDistance) && rsm.IsBetterThan(bestMatch) {
				bestMatch = rsm
			}
		}

		if !rs.Bidirectional {
			continue
		}

		if last, ok := rs.MatchSegment(workout, p, false); ok {
			rsm := rs.NewRouteSegmentMatch(workout, p, last)
			if rsm.MatchesDistance(rs.TotalDistance) && rsm.IsBetterThan(bestMatch) {
				bestMatch = rsm
			}
		}
	}

	return bestMatch
}

// MatchSegment starts at a point and continues the workout track while it finds
// each next point of the route segment, assuming there are many more points in
// the workout track than the route segment.
// If it can't find all points of the segment in the correct order, it returns false.
// Otherwise it returns the last point index of the route that matches the final
// point of the route segment.
// If forward is true, we increment the index, otherwise we decrement it
func (rs *RouteSegment) MatchSegment(workout *Workout, start int, forward bool) (int, bool) {
	workoutLength := len(workout.Data.Details.Points)
	segmentLength := len(rs.Points)

	cur := 0
	if !forward {
		cur = segmentLength - 1
	}

	for i := range workoutLength {
		index := (start + i) % workoutLength

		d := rs.Points[cur].DistanceTo(&workout.Data.Details.Points[index])
		if d > MaxDeltaMeter {
			continue
		}

		if forward {
			cur++

			if cur == segmentLength {
				return index, true
			}
		} else {
			cur--

			if cur == 0 {
				return index, true
			}
		}

		if !rs.Circular && index < start {
			break
		}
	}

	return 0, false
}

// StartingPoints finds all points that are closer than MaxDeltaMeter to the
// segment's starting point
func (rs *RouteSegment) StartingPoints(points []MapPoint) []int {
	var r []int

	start := rs.Points[0]

	for i, p := range points {
		d := start.DistanceTo(&p)
		if d < MaxDeltaMeter {
			r = append(r, i)
		}
	}

	return r
}

// FindMatches will find all workouts that match the current route segment
// The result will contain a list of RouteSegmentMatches, which will contain
// the workout, the point of the workout along the segment, and the total
// distance and duration of the segment for this workout.
func (w *Workout) FindMatches(routeSegments []*RouteSegment) []*RouteSegmentMatch {
	if !w.HasTracks() {
		return nil
	}

	var result []*RouteSegmentMatch

	for _, rs := range routeSegments {
		if m := rs.Match(w); m != nil {
			result = append(result, m)
		}
	}

	return result
}
