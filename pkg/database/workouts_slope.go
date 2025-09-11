package database

// SlopeState represents the type of slope detected.
type SlopeState string

const (
	// Ascent states
	StartClimb SlopeState = "start_climb"
	InClimb    SlopeState = "in_climb"
	EndClimb   SlopeState = "end_climb"

	// Descent states
	StartDescent SlopeState = "start_descent"
	InDescent    SlopeState = "in_descent"
	EndDescent   SlopeState = "end_descent"

	// Flat state
	Flat SlopeState = "flat"
)

// Slope detection thresholds. These are suggestions and can be adjusted.
const (
	// StartClimbThreshold defines the minimum slope to be considered a climb.
	StartClimbThreshold float64 = 0.02 // 2% grade
	// StartDescentThreshold defines the minimum negative slope to be considered a descent.
	StartDescentThreshold float64 = -0.02 // -2% grade
	// SmoothingDistance defines the look-back window for the average slope calculation.
	SmoothingDistance float64 = 200.0 // 200 meters
)

// SlopeDetection holds the calculated slope and climb state for a given point.
type SlopeDetection struct {
	ClimbNumber int        `json:"climb_number"`
	SlopeGrade  float64    `json:"slope_grade"` // Slope as a decimal (-1.0 to 1.0)
	SlopeState  SlopeState `json:"slope_state"` // The detected climb/descent state
}

// calculateSlopes processes a slice of MapPoints and returns a slice of ClimbDetection.
func (d *MapData) CalculateSlopes() {
	if d.Details == nil {
		return
	}

	points := d.Details.Points
	if len(points) < 2 {
		return
	}

	climbNumber := 0

	// The first point can't have a slope since it has no previous point.
	points[0].SlopeGrade = 0.0
	points[0].SlopeState = Flat
	points[0].ClimbNumber = climbNumber

	for i := 1; i < len(points); i++ {
		currentPoint := points[i]

		// Find the index of the point SmoothingDistance meters back.
		var lookbackIndex int
		for j := i - 1; j >= 0; j-- {
			if currentPoint.TotalDistance-points[j].TotalDistance >= SmoothingDistance {
				lookbackIndex = j
				break
			}
		}

		// Use the point at the look-back distance.
		lookbackPoint := points[lookbackIndex]

		// Calculate the total distance and elevation change over the smoothing window.
		deltaDistance := currentPoint.TotalDistance - lookbackPoint.TotalDistance
		deltaElevation := currentPoint.Elevation - lookbackPoint.Elevation

		var slopeGrade float64
		if deltaDistance > 0 {
			slopeGrade = deltaElevation / deltaDistance
		}

		state := detectSlopeState(slopeGrade, points[i-1].SlopeState)

		if state == StartClimb {
			climbNumber++
		}

		points[i].SlopeGrade = slopeGrade
		points[i].SlopeState = state
		points[i].ClimbNumber = climbNumber
	}
}

// detectSlopeState determines the current climb state based on the slope grade and previous state.
func detectSlopeState(slopeGrade float64, prevState SlopeState) SlopeState {
	switch {
	// A new climb starts if the slope is positive and steep enough.
	case slopeGrade > StartClimbThreshold:
		if prevState == StartClimb || prevState == InClimb {
			return InClimb
		}
		return StartClimb

	// A new descent starts if the slope is negative and steep enough.
	case slopeGrade < StartDescentThreshold:
		if prevState == StartDescent || prevState == InDescent {
			return InDescent
		}
		return StartDescent

	// If the slope is close to zero, we check the previous state to see if a climb or descent is ending.
	case slopeGrade >= StartDescentThreshold && slopeGrade <= StartClimbThreshold:
		if prevState == StartClimb || prevState == InClimb {
			return EndClimb
		} else if prevState == StartDescent || prevState == InDescent {
			return EndDescent
		}
		return Flat

	default:
		return Flat
	}
}
