package database

import (
	"crypto/sha256"
	"errors"
	"html/template"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/google/uuid"
	"github.com/jovandeginste/workout-tracker/internal/pkg/converters"
	"github.com/microcosm-cc/bluemonday"
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

var ErrInvalidData = errors.New("could not convert data to a GPX structure")

type Workout struct {
	gorm.Model
	Name                string               `gorm:"not null"`                                  // The name of the workout
	Date                *time.Time           `gorm:"not null;uniqueIndex:idx_start_user"`       // The timestamp the workout was recorded
	UserID              uint                 `gorm:"not null;index;uniqueIndex:idx_start_user"` // The ID of the user who owns the workout
	Dirty               bool                 // Whether the workout has been modified and the details should be re-rendered
	PublicUUID          *uuid.UUID           `gorm:"type:uuid;uniqueIndex"` // UUID to publicly share a workout - this UUID can be rotated
	User                *User                // The user who owns the workout
	Notes               string               // The notes associated with the workout, in markdown
	Type                WorkoutType          // The type of the workout
	Data                *MapData             `json:",omitempty"`                                    // The map data associated with the workout
	GPX                 *GPXData             `json:",omitempty"`                                    // The file data associated with the workout
	Equipment           []Equipment          `json:",omitempty" gorm:"many2many:workout_equipment"` // Which equipment is used for this workout
	RouteSegmentMatches []*RouteSegmentMatch `json:",omitempty"`                                    // Which route segments match
}

type GPXData struct {
	gorm.Model
	WorkoutID uint   `gorm:"not null;uniqueIndex"` // The ID of the workout
	Content   []byte `gorm:"type:text"`            // The file content
	Checksum  []byte `gorm:"not null;uniqueIndex"` // The checksum of the content
	Filename  string // The filename of the file
}

func (w *Workout) GetDate() time.Time {
	if w.Date == nil {
		return time.Now()
	}

	return *w.Date
}

func (w *Workout) Filename() string {
	if !w.HasFile() {
		return w.Name + ".txt"
	}

	return w.GPX.Filename
}

func (w *Workout) HasFile() bool {
	if w.GPX == nil {
		return false
	}

	return w.GPX.Filename != "" && w.GPX.Content != nil
}

func (w *Workout) HasTracks() bool {
	if w.Data == nil {
		return false
	}

	if w.Data.Center.IsZero() {
		return false
	}

	if w.Data.Details == nil {
		return false
	}

	if len(w.Data.Details.Points) == 0 {
		return false
	}

	return w.Type.IsLocation()
}

func (w *Workout) Weight() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalWeight
}

func (w *Workout) Repetitions() int {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalRepetitions
}

func (w *Workout) Duration() time.Duration {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalDuration
}

func (w *Workout) FullAddress() string {
	if w.Data == nil {
		return ""
	}

	if w.Data.Address != nil {
		return w.Data.Address.FormattedAddress
	}

	return w.Data.AddressString
}

func (w *Workout) Address() string {
	if w.Data == nil {
		return ""
	}

	if w.Data.AddressString != "" {
		return w.Data.AddressString
	}

	return w.Data.addressString()
}

func (w *Workout) Distance() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalDistance
}

func (w *Workout) MarkdownNotes() template.HTML {
	doc := parser.NewWithExtensions(parser.CommonExtensions).Parse([]byte(w.Notes))
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(markdown.Render(doc, renderer))

	return template.HTML(safeHTML) //nolint:gosec // We escaped all unsafe HTML with bluemonday
}

func (d *GPXData) Save(db *gorm.DB) error {
	if d.Content == nil {
		return ErrInvalidData
	}

	return db.Save(d).Error
}

func NewWorkout(u *User, workoutType WorkoutType, notes string, filename string, content []byte) (*Workout, error) {
	if u == nil {
		return nil, ErrNoUser
	}

	filename = filepath.Base(filename)

	gpxContent, err := converters.Parse(filename, content)
	if err != nil {
		return nil, err
	}

	data := gpxAsMapData(gpxContent)
	if filename == "" {
		filename = data.Name + ".gpx"
	}

	h := sha256.New()
	h.Write(content)

	if workoutType == WorkoutTypeAutoDetect {
		workoutType = autoDetectWorkoutType(data, gpxContent)
	}

	w := Workout{
		User:   u,
		UserID: u.ID,
		Dirty:  true,
		Name:   gpxName(gpxContent),
		Data:   data,
		Notes:  notes,
		Type:   workoutType,
		Date:   gpxDate(gpxContent),
		GPX: &GPXData{
			Content:  content,
			Checksum: h.Sum(nil),
			Filename: filename,
		},
	}

	return &w, nil
}

func workoutTypeFromGpxTrackType(gpxType string) (WorkoutType, bool) {
	switch strings.ToLower(gpxType) {
	case "running", "run":
		return WorkoutTypeRunning, true
	case "walking", "walk":
		return WorkoutTypeWalking, true
	case "cycling", "cycle":
		return WorkoutTypeCycling, true
	case "snowboarding":
		return WorkoutTypeSnowboarding, true
	case "skiing":
		return WorkoutTypeSkiing, true
	case "swimming":
		return WorkoutTypeSwimming, true
	case "kayaking":
		return WorkoutTypeKayaking, true
	case "golfing":
		return WorkoutTypeGolfing, true
	case "hiking":
		return WorkoutTypeHiking, true
	default:
		return WorkoutTypeAutoDetect, false
	}
}

func autoDetectWorkoutType(data *MapData, gpxContent *gpx.GPX) WorkoutType {
	// If the GPX file mentions a workout type (for the first track), use it
	if len(gpxContent.Tracks) > 0 {
		firstTrack := &gpxContent.Tracks[0]

		if workoutType, ok := workoutTypeFromGpxTrackType(firstTrack.Type); ok {
			return workoutType
		}
	}

	if 3.6*data.AverageSpeedNoPause() > 15.0 {
		return WorkoutTypeCycling
	}

	if 3.6*data.AverageSpeedNoPause() > 7.0 {
		return WorkoutTypeRunning
	}

	return WorkoutTypeWalking
}

func GetRecentWorkouts(db *gorm.DB, count int) ([]Workout, error) {
	var w []Workout

	if err := db.Preload("Data").Preload("User").Order("date DESC").Limit(count).Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorkouts(db *gorm.DB) ([]*Workout, error) {
	var w []*Workout

	if err := db.Preload("Data").Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorkoutWithGPXByUUID(db *gorm.DB, u uuid.UUID) (*Workout, error) {
	return GetWorkoutByUUID(db.Preload("GPX"), u)
}

func GetWorkoutWithGPX(db *gorm.DB, id int) (*Workout, error) {
	return GetWorkout(db.Preload("GPX"), id)
}

func GetWorkoutDetailsByUUID(db *gorm.DB, u uuid.UUID) (*Workout, error) {
	return GetWorkoutWithGPXByUUID(db.Preload("Data.Details"), u)
}

func GetWorkoutDetails(db *gorm.DB, id int) (*Workout, error) {
	return GetWorkoutWithGPX(db.Preload("Data.Details"), id)
}

func GetWorkoutByUUID(db *gorm.DB, u uuid.UUID) (*Workout, error) {
	w := Workout{
		PublicUUID: &u,
	}

	if err := db.
		Preload("RouteSegmentMatches.RouteSegment").
		Preload("Data").
		Preload("User").
		Preload("Equipment").
		Where(&w).
		First(&w).
		Error; err != nil {
		return nil, err
	}

	sort.Slice(w.RouteSegmentMatches, func(i, j int) bool {
		return w.RouteSegmentMatches[i].Distance > w.RouteSegmentMatches[j].Distance
	})

	return &w, nil
}

func GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w Workout

	if err := db.
		Preload("RouteSegmentMatches.RouteSegment").
		Preload("Data").
		Preload("User").
		Preload("Equipment").
		First(&w, id).
		Error; err != nil {
		return nil, err
	}

	sort.Slice(w.RouteSegmentMatches, func(i, j int) bool {
		return w.RouteSegmentMatches[i].Distance > w.RouteSegmentMatches[j].Distance
	})

	return &w, nil
}

func (w *Workout) Delete(db *gorm.DB) error {
	return db.Unscoped().Select("GPX", "Data").Delete(w).Error
}

func (w *Workout) Create(db *gorm.DB) error {
	if w.Data == nil {
		return ErrInvalidData
	}

	return db.Create(w).Error
}

func (w *Workout) Save(db *gorm.DB) error {
	if w.Data == nil {
		return ErrInvalidData
	}

	if err := w.Data.Save(db); err != nil {
		return err
	}

	if w.RouteSegmentMatches != nil {
		if err := db.Model(w).Association("RouteSegmentMatches").Replace(w.RouteSegmentMatches); err != nil {
			return err
		}
	}

	return db.Save(w).Error
}

func (w *Workout) AsGPX() (*gpx.GPX, error) {
	if !w.HasFile() {
		return nil, errors.New("workout has no GPX")
	}

	return converters.Parse(w.GPX.Filename, w.GPX.Content)
}

func (w *Workout) setData(data *MapData) {
	if w.Data == nil {
		w.Data = data
		return
	}

	dataID := w.Data.ID
	dataCreatedAt := w.Data.CreatedAt

	w.Data = data
	w.Data.ID = dataID
	w.Data.CreatedAt = dataCreatedAt

	if !w.Data.Center.IsZero() {
		w.Data.Address = w.Data.Center.Address()
	}

	w.Data.UpdateAddress()
}

func (w *Workout) UpdateData(db *gorm.DB) error {
	if !w.HasFile() {
		// We only update data from (stored) GPX data
		w.Dirty = false

		return w.Save(db)
	}

	gpxContent, err := w.AsGPX()
	if err != nil {
		return err
	}

	w.setData(gpxAsMapData(gpxContent))

	if err := w.Data.Save(db); err != nil {
		return err
	}

	if err := w.UpdateRouteSegmentMatches(db); err != nil {
		return err
	}

	w.Dirty = false

	return w.Save(db)
}

func (w *Workout) UpdateRouteSegmentMatches(db *gorm.DB) error {
	routeSegments, err := GetRouteSegments(db)
	if err != nil {
		return err
	}

	w.RouteSegmentMatches = w.FindMatches(routeSegments)

	return nil
}

func (w *Workout) RepetitionFrequencyPerMinute() float64 {
	if w.Data == nil {
		return 0
	}

	return float64(w.Data.TotalRepetitions) / w.Duration().Minutes()
}

func (w *Workout) HasCalories() bool {
	return w.Type.IsDuration()
}

func (w *Workout) CaloriesBurned() float64 {
	if !w.Type.IsDuration() {
		return 0
	}

	weight := 70.0 // assume weight is 70 kg for now
	// Calories burned = weight * time * intensity (MET)
	cb := weight * w.Duration().Hours() * w.MET()

	return cb
}

func (w *Workout) HasElevation() bool {
	return w.HasExtraMetric("elevation")
}

func (w *Workout) HasCadence() bool {
	return w.HasExtraMetric("cadence")
}

func (w *Workout) HasHeartRate() bool {
	return w.HasExtraMetric("heart-rate")
}

func (w *Workout) HasHeading() bool {
	return w.HasExtraMetric("heading")
}

func (w *Workout) HasExtraMetric(name string) bool {
	if w.Data == nil || w.Data.Details == nil {
		return false
	}

	for _, d := range w.Data.Details.Points {
		if _, ok := d.ExtraMetrics[name]; ok {
			return true
		}
	}

	return false
}

func (w *Workout) EquipmentIDs() []uint {
	ids := make([]uint, 0, len(w.Equipment))

	for _, e := range w.Equipment {
		ids = append(ids, e.ID)
	}

	return ids
}

func (w *Workout) Uses(e Equipment) bool {
	return slices.Contains(w.EquipmentIDs(), e.ID)
}

func (w *Workout) PreferredAverageSpeedMetric(preferredUnits *UserPreferredUnits) template.HTML {
	return w.Type.PreferredSpeedMetric(w.Data.AverageSpeedNoPause(), preferredUnits)
}
