package database

import (
	"cmp"
	"math"
	"slices"
)

// SlopeState represents the type of slope detected.
type SlopeState string

const (
	// State machine states for the detector logic.
	StateSearching       = "SEARCHING"
	StateInSegment       = "IN_SEGMENT"
	StateEvaluatingPause = "EVALUATING_PAUSE"
	StateStartClimb      = "START_CLIMB"
	StateEndClimb        = "END_CLIMB"
	StateStartDescent    = "START_DESCENT"
	StateEndDescent      = "END_DESCENT"

	// Thresholds from Python code
	StartClimbThreshold   float64 = 0.02
	EndClimbThreshold     float64 = 0.01
	MaxPauseLengthMeters  float64 = 200.0
	MaxPauseDescentMeters float64 = 10.0
	MinGain               float64 = 20.0
	MinLength             float64 = 300.0
)

// Segment represents a detected climb or descent.
type Segment struct {
	Index         int     `json:"index"`
	Type          string  `json:"type"`
	StartDistance float64 `json:"start_km"`
	EndDistance   float64 `json:"end_km"`
	Elevation     float64 `json:"elev_gain,omitempty"`
	ElevLoss      float64 `json:"elev_loss,omitempty"`
	Length        float64 `json:"length_m"`
	AvgSlope      float64 `json:"avg_slope"`
	StartIdx      int     `json:"start_idx"`
	EndIdx        int     `json:"end_idx"`
	Category      string  `json:"category"`
}

// Detector holds the state for the segment detection process.
type Detector struct {
	segments  []Segment
	kind      string
	slopeSign float64

	state                string
	currentSegmentPoints []*MapPoint

	startIdx      int
	pauseStartIdx int
	pauseLength   float64
	pauseDescent  float64
}

// CalculateSlopes processes a slice of MapPoints and returns a slice of ClimbDetection.
func (m *MapData) CalculateSlopes() {
	climbs := DetectSignificantSegments(m.Details.Points, "climb")
	descents := DetectSignificantSegments(m.Details.Points, "descent")

	climbs = append(climbs, descents...)
	slices.SortFunc(climbs, func(a, b Segment) int {
		return cmp.Compare(a.StartDistance, b.StartDistance)
	})

	m.Climbs = climbs
}

// NewDetector initializes a new Detector for a given kind ("climb" or "descent").
func NewDetector(kind string) *Detector {
	slopeSign := 1.0
	if kind != "climb" {
		slopeSign = -1.0
	}

	return &Detector{
		kind:      kind,
		slopeSign: slopeSign,
		state:     StateSearching,
	}
}

// SmoothSlopeGrades computes a weighted average slope at each point.
func SmoothSlopeGrades(points []MapPoint, windowMeters float64) {
	for i := range points {
		centerDist := points[i].TotalDistance2D
		var weightedSlopeSum, totalWeight float64

		for j := range points {
			distDiff := points[j].TotalDistance2D - centerDist
			distFromCenter := math.Abs(distDiff)
			if distFromCenter > windowMeters/2 || distFromCenter < MaxDeltaMeter/2 {
				continue
			}

			elevDiff := points[j].Elevation - points[i].Elevation
			slope := elevDiff / distDiff

			weight := 1.0 / distFromCenter
			weightedSlopeSum += slope * weight
			totalWeight += weight
		}

		if totalWeight > 0 {
			points[i].SlopeGrade = weightedSlopeSum / totalWeight
		} else {
			points[i].SlopeGrade = 0
		}
	}
}

// DetectSignificantSegments processes a slice of points to find climbs or descents.
func DetectSignificantSegments(points []MapPoint, kind string) []Segment {
	detector := NewDetector(kind)

	if len(points) < 2 {
		return nil
	}

	SmoothSlopeGrades(points, 300.0)

	// Start with the first point.
	detector.currentSegmentPoints = append(detector.currentSegmentPoints, &points[0])

	for i := 1; i < len(points); i++ {
		prevPoint := &points[i-1]
		currentPoint := &points[i]

		distDiff := currentPoint.TotalDistance2D - prevPoint.TotalDistance2D
		elevDiff := (currentPoint.Elevation - prevPoint.Elevation)

		slope := currentPoint.SlopeGrade

		// Adjust slope and elevation diff based on "kind" (climb or descent).
		effectiveSlope := slope * detector.slopeSign
		effectiveElevDiff := elevDiff * detector.slopeSign

		// State machine logic
		switch detector.state {
		case StateSearching:
			if effectiveSlope >= StartClimbThreshold {
				detector.state = StateInSegment
				detector.startIdx = i - 1
				// Add previous and current point to start the segment.
				detector.currentSegmentPoints = []*MapPoint{prevPoint, currentPoint}
			}

		case StateInSegment:
			if effectiveSlope >= EndClimbThreshold {
				detector.currentSegmentPoints = append(detector.currentSegmentPoints, currentPoint)
			} else {
				detector.state = StateEvaluatingPause
				detector.pauseStartIdx = i - 1
				detector.pauseLength = 0
				detector.pauseDescent = 0
				detector.currentSegmentPoints = append(detector.currentSegmentPoints, currentPoint)
			}

		case StateEvaluatingPause:
			detector.currentSegmentPoints = append(detector.currentSegmentPoints, currentPoint)
			detector.pauseLength += distDiff
			if effectiveElevDiff < 0 {
				detector.pauseDescent += math.Abs(effectiveElevDiff)
			}

			if effectiveSlope >= EndClimbThreshold {
				detector.state = StateInSegment
			} else if detector.pauseLength > MaxPauseLengthMeters || detector.pauseDescent > MaxPauseDescentMeters {
				// The pause is too long or a significant descent occurred.
				// Finalize the segment before the pause.
				finalSegmentPoints := detector.currentSegmentPoints[:len(detector.currentSegmentPoints)-(i-detector.pauseStartIdx)]

				detector.validateAndAppendSegment(finalSegmentPoints)

				// Reset state to search for a new segment.
				detector.state = StateSearching
				detector.currentSegmentPoints = []*MapPoint{}
			}
		}
	}

	// Final check for any segment in progress at the end of the data.
	if detector.state == StateInSegment || detector.state == StateEvaluatingPause {
		detector.validateAndAppendSegment(detector.currentSegmentPoints)
	}

	return detector.segments
}

// validateAndAppendSegment is a private method that validates and appends a segment to the detector's slice.
func (d *Detector) validateAndAppendSegment(segmentPoints []*MapPoint) {
	if len(segmentPoints) < 2 {
		return
	}

	length := segmentPoints[len(segmentPoints)-1].TotalDistance2D - segmentPoints[0].TotalDistance2D

	var gain float64
	for i := 1; i < len(segmentPoints); i++ {
		elevDiff := segmentPoints[i].Elevation - segmentPoints[i-1].Elevation
		if (d.kind == "climb" && elevDiff > 0) || (d.kind == "descent" && elevDiff < 0) {
			gain += math.Abs(elevDiff)
		}
	}

	if length > MinLength && gain > MinGain {
		avgSlope := 0.0
		if length > 0 {
			avgSlope = gain / length
		}

		endIdx := d.startIdx + len(segmentPoints) - 1
		category := ClassifyClimbCategory(length, avgSlope)

		segment := Segment{
			Type:          d.kind,
			StartDistance: segmentPoints[0].TotalDistance2D,
			EndDistance:   segmentPoints[len(segmentPoints)-1].TotalDistance2D,
			Length:        length,
			StartIdx:      d.startIdx,
			EndIdx:        endIdx,
			Category:      category,
			Index:         len(d.segments) + 1,
		}

		if d.kind == "climb" {
			segment.AvgSlope = avgSlope
			segment.Elevation = gain
		} else {
			segment.AvgSlope = -avgSlope
			segment.Elevation = -gain
		}

		d.segments = append(d.segments, segment)
	}
}

func ClassifyClimbCategory(length, slope float64) string {
	switch {
	case length >= 10000 && slope >= 0.06:
		return "Hors Catégorie"
	case length >= 8000 && slope >= 0.05:
		return "Category 1"
	case length >= 5000 && slope >= 0.04:
		return "Category 2"
	case length >= 3000 && slope >= 0.03:
		return "Category 3"
	case length >= 2000 && slope >= 0.03:
		return "Category 4"
	case length >= 1000 && slope >= 0.02:
		return "Category 5"
	case length >= 500 && slope >= 0.01:
		return "Category 6"
	default:
		return "Uncategorized"
	}
}
