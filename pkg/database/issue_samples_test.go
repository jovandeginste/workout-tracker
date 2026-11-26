package database

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The samples in testdata/issues are workout files that were donated in GitHub
// issues. Every sample has a golden file with the expected result of the
// parsing pipeline. Regenerate the golden files with:
//
//	TZ=UTC go test ./pkg/database -run TestIssueSamples -update-issue-samples
//
// The test suite runs in UTC; use make test-go or set TZ=UTC when invoking
// go test directly.
var updateIssueSamples = flag.Bool("update-issue-samples", false, "rewrite the golden JSON files for the donated issue samples")

const issueSamplesDir = "testdata/issues"

type issueSampleSource struct {
	Issues       []int  `json:"issues"`
	File         string `json:"file"`
	OriginalName string `json:"original"`
	URL          string `json:"url"`
	SHA256       string `json:"sha256"`
	Bytes        int    `json:"bytes"`
}

type issueSampleResult struct {
	Issues       []int               `json:"issues"`
	SourceURL    string              `json:"sourceUrl"`
	OriginalName string              `json:"originalName"`
	File         string              `json:"file"`
	Panicked     bool                `json:"panicked,omitempty"`
	Error        string              `json:"error,omitempty"`
	Workout      *issueSampleWorkout `json:"workout,omitempty"`
}

type issueSampleWorkout struct {
	Name                string             `json:"name"`
	Type                string             `json:"type"`
	Start               string             `json:"start"`
	End                 string             `json:"end"`
	Creator             string             `json:"creator"`
	Center              issueSampleCenter  `json:"center"`
	TotalDistance       float64            `json:"totalDistance"`
	TotalDistance2D     float64            `json:"totalDistance2D"`
	TotalDuration       string             `json:"totalDuration"`
	PauseDuration       string             `json:"pauseDuration"`
	MaxSpeed            float64            `json:"maxSpeed"`
	MaxCadence          float64            `json:"maxCadence"`
	MaxPower            float64            `json:"maxPower"`
	AverageSpeed        float64            `json:"averageSpeed"`
	AverageSpeedNoPause float64            `json:"averageSpeedNoPause"`
	AverageCadence      float64            `json:"averageCadence"`
	AveragePower        float64            `json:"averagePower"`
	MinElevation        float64            `json:"minElevation"`
	MaxElevation        float64            `json:"maxElevation"`
	TotalUp             float64            `json:"totalUp"`
	TotalDown           float64            `json:"totalDown"`
	TotalRepetitions    int                `json:"totalRepetitions"`
	TotalWeight         float64            `json:"totalWeight"`
	TotalCalories       float64            `json:"totalCalories"`
	ExtraMetrics        []string           `json:"extraMetrics"`
	PointsCount         int                `json:"pointsCount"`
	PointsDigest        string             `json:"pointsDigest"`
	FirstPoint          *issueSamplePoint  `json:"firstPoint,omitempty"`
	LastPoint           *issueSamplePoint  `json:"lastPoint,omitempty"`
	Climbs              []issueSampleClimb `json:"climbs"`
	HeatmapPointsCount  int                `json:"heatmapPointsCount"`
	HeatmapPointsDigest string             `json:"heatmapPointsDigest"`
}

type issueSampleClimb struct {
	Index    int              `json:"index"`
	Type     string           `json:"type"`
	StartIdx int              `json:"startIdx"`
	EndIdx   int              `json:"endIdx"`
	Gain     float64          `json:"gain"`
	Length   float64          `json:"length"`
	AvgSlope float64          `json:"avgSlope"`
	Duration string           `json:"duration"`
	Category string           `json:"category"`
	Start    issueSamplePoint `json:"start"`
	End      issueSamplePoint `json:"end"`
}

type issueSampleCenter struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	TZ  string  `json:"tz"`
}

type issueSamplePoint struct {
	Time            string             `json:"time"`
	Lat             float64            `json:"lat"`
	Lng             float64            `json:"lng"`
	Elevation       float64            `json:"elevation"`
	Distance        float64            `json:"distance"`
	Distance2D      float64            `json:"distance2D"`
	TotalDistance   float64            `json:"totalDistance"`
	TotalDistance2D float64            `json:"totalDistance2D"`
	Duration        string             `json:"duration"`
	TotalDuration   string             `json:"totalDuration"`
	SlopeGrade      float64            `json:"slopeGrade"`
	ExtraMetrics    map[string]float64 `json:"extraMetrics,omitempty"`
}

func TestIssueSamples(t *testing.T) {
	sources := issueSampleSources(t)
	files := issueSampleFiles(t)
	require.NotEmpty(t, files, "no issue samples found in %s", issueSamplesDir)

	filesSeen := make(map[string]bool, len(files))

	for _, file := range files {
		base := filepath.Base(file)
		filesSeen[base] = true

		t.Run(base, func(t *testing.T) {
			source, ok := sources[base]
			require.True(t, ok, "no entry for %s in sources.json", base)

			content, err := os.ReadFile(file)
			require.NoError(t, err)
			require.Len(t, content, source.Bytes, "sample %s does not match sources.json", base)
			require.Equal(t, source.SHA256, checksum(content), "sample %s does not match sources.json", base)

			result, panicValue := parseIssueSample(file, content)
			if panicValue != nil {
				t.Logf("parser panicked: %v", panicValue)
			}

			result.File = base
			result.Issues = source.Issues
			result.SourceURL = source.URL
			result.OriginalName = source.OriginalName

			// The extension is part of the golden name: the same workout can be
			// donated as both GPX and TCX (issue #57).
			golden := file + ".json"

			if *updateIssueSamples {
				writeIssueSampleGolden(t, golden, result)
			}

			expected := readIssueSampleGolden(t, golden)

			assert.Equal(t, expected, result)
		})
	}

	for base := range sources {
		assert.True(t, filesSeen[base], "sources.json mentions %s, but the file is missing", base)
	}
}

func issueSampleFiles(t *testing.T) []string {
	t.Helper()

	extensions := []string{".gpx", ".fit", ".tcx", ".ftb"}
	files := []string{}

	for _, ext := range extensions {
		matches, err := filepath.Glob(filepath.Join(issueSamplesDir, "*"+ext))
		require.NoError(t, err)

		files = append(files, matches...)
	}

	sort.Strings(files)

	return files
}

func issueSampleSources(t *testing.T) map[string]issueSampleSource {
	t.Helper()

	content, err := os.ReadFile(filepath.Join(issueSamplesDir, "sources.json"))
	require.NoError(t, err)

	var list []issueSampleSource
	require.NoError(t, json.Unmarshal(content, &list))

	sources := make(map[string]issueSampleSource, len(list))

	for _, source := range list {
		sources[source.File] = source
	}

	return sources
}

func parseIssueSample(filename string, content []byte) (result issueSampleResult, panicValue any) {
	defer func() {
		if r := recover(); r != nil {
			result.Panicked = true
			panicValue = r
		}
	}()

	workouts, err := NewWorkout(defaultUser(), WorkoutTypeAutoDetect, "", filepath.Base(filename), content)
	if err != nil {
		result.Error = err.Error()
		return result, nil
	}

	if len(workouts) != 1 {
		result.Error = fmt.Sprintf("expected 1 workout, got %d", len(workouts))
		return result, nil
	}

	workout := workouts[0]

	// The background worker runs these derived calculations after importing.
	// Include them so the goldens cover climb detection, slope grades, and
	// cadence/power averages too.
	workout.UpdateAverages()

	if workout.Data != nil && workout.Data.Details != nil {
		workout.Data.CalculateSlopes(workout.Type)
	}

	result.Workout = issueSampleWorkoutFrom(workout)
	if workout.Data != nil && workout.Data.Details != nil {
		result.Workout.PointsCount = len(workout.Data.Details.Points)
		result.Workout.PointsDigest = issueSamplePointsDigest(workout.Data.Details.Points)
		result.Workout.FirstPoint = issueSamplePointFrom(workout.Data.Details.Points, 0)
		result.Workout.LastPoint = issueSamplePointFrom(workout.Data.Details.Points, -1)
		result.Workout.HeatmapPointsCount = len(workout.Data.Details.HeatmapPoints)
		result.Workout.HeatmapPointsDigest = issueSampleHeatmapDigest(workout.Data.Details.HeatmapPoints)
	}

	return result, nil
}

func issueSampleWorkoutFrom(w *Workout) *issueSampleWorkout {
	data := w.data()

	result := &issueSampleWorkout{
		Name:                w.Name,
		Type:                w.Type.String(),
		Start:               w.Date.Format(time.RFC3339Nano),
		End:                 w.GetEnd().Format(time.RFC3339Nano),
		Creator:             data.Creator,
		Center:              issueSampleCenter{Lat: roundFloat(data.Center.Lat, 9), Lng: roundFloat(data.Center.Lng, 9), TZ: data.Center.TZ},
		TotalDistance:       roundFloat(data.TotalDistance, 6),
		TotalDistance2D:     roundFloat(data.TotalDistance2D, 6),
		TotalDuration:       formatDuration(data.TotalDuration),
		PauseDuration:       formatDuration(data.PauseDuration),
		MaxSpeed:            roundFloat(data.MaxSpeed, 6),
		MaxCadence:          roundFloat(data.MaxCadence, 6),
		MaxPower:            roundFloat(data.MaxPower, 6),
		AverageSpeed:        roundFloat(data.AverageSpeed, 6),
		AverageSpeedNoPause: roundFloat(data.AverageSpeedNoPause, 6),
		AverageCadence:      roundFloat(data.AverageCadence, 6),
		AveragePower:        roundFloat(data.AveragePower, 6),
		MinElevation:        roundFloat(data.MinElevation, 6),
		MaxElevation:        roundFloat(data.MaxElevation, 6),
		TotalUp:             roundFloat(data.TotalUp, 6),
		TotalDown:           roundFloat(data.TotalDown, 6),
		TotalRepetitions:    data.TotalRepetitions,
		TotalWeight:         roundFloat(data.TotalWeight, 6),
		TotalCalories:       roundFloat(data.TotalCalories, 6),
		ExtraMetrics:        data.ExtraMetrics,
		Climbs:              issueSampleClimbsFrom(data),
	}

	if result.ExtraMetrics == nil {
		result.ExtraMetrics = []string{}
	}

	return result
}

func issueSampleClimbsFrom(data *MapData) []issueSampleClimb {
	result := make([]issueSampleClimb, 0, len(data.Climbs))

	for _, climb := range data.Climbs {
		result = append(result, issueSampleClimb{
			Index:    climb.Index,
			Type:     string(climb.Type),
			StartIdx: climb.StartIdx,
			EndIdx:   climb.EndIdx,
			Gain:     roundFloat(climb.Gain, 6),
			Length:   roundFloat(climb.Length, 6),
			AvgSlope: roundFloat(climb.AvgSlope, 6),
			Duration: formatDuration(climb.Duration),
			Category: string(climb.Category),
			Start:    issueSamplePointValue(climb.Start),
			End:      issueSamplePointValue(climb.End),
		})
	}

	return result
}

func issueSamplePointFrom(points []MapPoint, index int) *issueSamplePoint {
	if len(points) == 0 {
		return nil
	}

	if index < 0 {
		index += len(points)
	}

	if index < 0 || index >= len(points) {
		return nil
	}

	result := issueSamplePointValue(points[index])

	return &result
}

func issueSamplePointValue(p MapPoint) issueSamplePoint {
	result := issueSamplePoint{
		Time:            p.Time.Format(time.RFC3339Nano),
		Lat:             roundFloat(p.Lat, 9),
		Lng:             roundFloat(p.Lng, 9),
		Elevation:       roundFloat(p.Elevation, 6),
		Distance:        roundFloat(p.Distance, 6),
		Distance2D:      roundFloat(p.Distance2D, 6),
		TotalDistance:   roundFloat(p.TotalDistance, 6),
		TotalDistance2D: roundFloat(p.TotalDistance2D, 6),
		Duration:        formatDuration(p.Duration),
		TotalDuration:   formatDuration(p.TotalDuration),
		SlopeGrade:      roundFloat(p.SlopeGrade, 6),
	}

	if len(p.ExtraMetrics) > 0 {
		result.ExtraMetrics = make(map[string]float64, len(p.ExtraMetrics))

		for key, value := range p.ExtraMetrics {
			result.ExtraMetrics[key] = roundFloat(value, 6)
		}
	}

	return result
}

// issueSamplePointsDigest hashes every parsed track point. Storing all points
// would make the golden files larger than the samples themselves, so the test
// only keeps this digest plus the first and last point.
func issueSamplePointsDigest(points []MapPoint) string {
	hash := sha256.New()

	for _, p := range points {
		keys := make([]string, 0, len(p.ExtraMetrics))
		for key := range p.ExtraMetrics {
			keys = append(keys, key)
		}

		slices.Sort(keys)

		fmt.Fprintf(
			hash,
			"%d|%.9f|%.9f|%.6f|%.6f|%.6f|%.6f|%.6f|%d|%d|%.6f",
			p.Time.UnixNano(),
			p.Lat,
			p.Lng,
			p.Elevation,
			p.Distance,
			p.Distance2D,
			p.TotalDistance,
			p.TotalDistance2D,
			p.Duration,
			p.TotalDuration,
			p.SlopeGrade,
		)

		for _, key := range keys {
			fmt.Fprintf(hash, "|%s=%.6f", key, p.ExtraMetrics[key])
		}

		fmt.Fprintln(hash)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

// issueSampleHeatmapDigest hashes the thinned-out heatmap points that are
// derived from the track points.
func issueSampleHeatmapDigest(points []HeatmapPoint) string {
	hash := sha256.New()

	for _, p := range points {
		fmt.Fprintf(hash, "%.5f|%.5f\n", p[0], p[1])
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func roundFloat(value float64, places int) float64 {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return value
	}

	factor := math.Pow10(places)

	return math.Round(value*factor) / factor
}

func formatDuration(d time.Duration) string {
	return d.Round(time.Microsecond).String()
}

func checksum(content []byte) string {
	hash := sha256.Sum256(content)

	return hex.EncodeToString(hash[:])
}

func writeIssueSampleGolden(t *testing.T, path string, result issueSampleResult) {
	t.Helper()

	content, err := json.MarshalIndent(result, "", "  ")
	require.NoError(t, err)

	//nolint:gosec // the path comes from a fixed glob over testdata
	require.NoError(t, os.WriteFile(path, append(content, '\n'), 0o600))
}

func readIssueSampleGolden(t *testing.T, path string) issueSampleResult {
	t.Helper()

	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		t.Fatalf("missing golden file %s; run: go test ./pkg/database -run TestIssueSamples -update-issue-samples", path)
	}

	require.NoError(t, err)

	var result issueSampleResult
	require.NoError(t, json.Unmarshal(content, &result))

	return result
}
