package database

import (
	"math"
	"time"
)

const (
	earthRadius = 6371 // Radius of Earth in kilometers
	gapPenalty  = 1e3  // A large penalty for gaps to prevent unnecessary insertions/deletions
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

// haversineDistance calculates the distance between two GPS points using the Haversine formula.
func haversineDistance(p1, p2 MapPoint) float64 {
	lat1, lon1, lat2, lon2 := p1.Lat*math.Pi/180, p1.Lng*math.Pi/180, p2.Lat*math.Pi/180, p2.Lng*math.Pi/180
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// initializeMatrices initializes the scoring and traceback matrices.
func initializeMatrices(n, m int) ([][]float64, [][]int) {
	score := make([][]float64, n+1)
	traceback := make([][]int, n+1)

	for i := range score {
		score[i] = make([]float64, m+1)
		traceback[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		score[i][0] = float64(i) * gapPenalty
	}

	for j := 0; j <= m; j++ {
		score[0][j] = float64(j) * gapPenalty
	}

	return score, traceback
}

// fillScoringMatrix fills the scoring matrix using the Needleman-Wunsch algorithm.
func fillScoringMatrix(track1, track2 []MapPoint, score [][]float64, traceback [][]int) {
	n, m := len(track1), len(track2)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			match := score[i-1][j-1] + haversineDistance(track1[i-1], track2[j-1])
			deletion := score[i-1][j] + gapPenalty
			insertion := score[i][j-1] + gapPenalty

			score[i][j] = math.Min(match, math.Min(deletion, insertion))

			switch {
			case score[i][j] == match:
				traceback[i][j] = 0 // Diagonal (match)
			case score[i][j] == deletion:
				traceback[i][j] = 1 // Up (delete)
			case score[i][j] == insertion:
				traceback[i][j] = 2 // Left (insert)
			}
		}
	}
}

// performTraceback performs the traceback to get the aligned tracks.
func performTraceback(track1, track2 []MapPoint, traceback [][]int) ([]MapPoint, []MapPoint) {
	var alignedTrack1, alignedTrack2 []MapPoint

	i, j := len(track1), len(track2)

	for i > 0 || j > 0 {
		switch {
		case i > 0 && j > 0 && traceback[i][j] == 0:
			alignedTrack1 = append([]MapPoint{track1[i-1]}, alignedTrack1...)
			alignedTrack2 = append([]MapPoint{track2[j-1]}, alignedTrack2...)
			i--
			j--
		case i > 0 && (j == 0 || traceback[i][j] == 1):
			alignedTrack1 = append([]MapPoint{track1[i-1]}, alignedTrack1...)
			alignedTrack2 = append([]MapPoint{{}}, alignedTrack2...)
			i--
		default:
			alignedTrack1 = append([]MapPoint{{}}, alignedTrack1...)
			alignedTrack2 = append([]MapPoint{track2[j-1]}, alignedTrack2...)
			j--
		}
	}

	return alignedTrack1, alignedTrack2
}

// calculateSimilarity calculates the normalized similarity score.
func calculateSimilarity(score [][]float64, n, m int) float64 {
	finalScore := score[n][m]
	maxPossibleScore := math.Max(float64(n), float64(m)) * gapPenalty

	return 1 - finalScore/maxPossibleScore
}

// needlemanWunsch aligns two GPS tracks using the Needleman-Wunsch algorithm and calculates a similarity score.
func needlemanWunsch(track1, track2 []MapPoint) (float64, []MapPoint, []MapPoint) {
	n, m := len(track1), len(track2)
	score, traceback := initializeMatrices(n, m)
	fillScoringMatrix(track1, track2, score, traceback)
	alignedTrack1, alignedTrack2 := performTraceback(track1, track2, traceback)
	similarity := calculateSimilarity(score, n, m)

	return similarity, alignedTrack1, alignedTrack2
}
