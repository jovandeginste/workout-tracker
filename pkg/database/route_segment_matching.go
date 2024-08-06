package database

import (
	"time"
)

// MaxDeltaMeter is the maximum distance in meters that a point can be away from
// the route segment
const MaxDeltaMeter = 20

// RouteSegmentMatch is a match between a route segment and a workout
type RouteSegmentMatch struct {
	Username        string        // The name of the user
	UserID          uint          // The ID of the user
	WorkoutName     string        // The name of the workout
	WorkoutID       uint          // The ID of the workout
	FirstID, LastID int           // The index of the first and last point of the route
	Distance        float64       // The total distance of the route segment for this workout
	Duration        time.Duration // The total duration of the route segment for this workout
	Points          int           // The total number of points of the route segment for this workout

	first, last MapPoint   // The first and last point of the route
	points      []MapPoint // The points of the route
}

// NewRouteSegmentMatch will create a new route segment match from a workout and
// the first and last point of the route along the route segment
func NewRouteSegmentMatch(workout *Workout, p, last int) *RouteSegmentMatch {
	rs := &RouteSegmentMatch{
		Username:    workout.User.Name,
		UserID:      workout.User.ID,
		WorkoutName: workout.Name,
		WorkoutID:   workout.ID,
		FirstID:     p,
		LastID:      last,

		first:  workout.Data.Details.Points[p],
		last:   workout.Data.Details.Points[last],
		points: workout.Data.Details.Points[p:last],
	}

	rs.calculate()

	return rs
}

// IsBetterThan returns true if the new route segment match is better than the
// current one
func (rsm *RouteSegmentMatch) IsBetterThan(current *RouteSegmentMatch) bool {
	return current == nil || rsm.Distance < current.Distance
}

// calculate will calculate the total distance and duration of the route
// segment, and the total number of points of this workout along the route
// segment
func (rsm *RouteSegmentMatch) calculate() {
	rsm.Distance = rsm.last.TotalDistance - rsm.first.TotalDistance
	rsm.Duration = rsm.last.TotalDuration - rsm.first.TotalDuration
	rsm.Points = len(rsm.points)
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
		if last, ok := rs.MatchSegment(workout, p); ok {
			rs := NewRouteSegmentMatch(workout, p, last)

			if rs.IsBetterThan(bestMatch) {
				bestMatch = rs
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
func (rs *RouteSegment) MatchSegment(workout *Workout, start int) (int, bool) {
	cur := 0

	for i := start; i < len(workout.Data.Details.Points); i++ {
		d := rs.Points[cur].DistanceTo(&workout.Data.Details.Points[i])
		if d > MaxDeltaMeter {
			continue
		}

		cur++
		if cur == len(rs.Points) {
			return i, true
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
		if start.DistanceTo(&p) < MaxDeltaMeter {
			r = append(r, i)
		}
	}

	return r
}
