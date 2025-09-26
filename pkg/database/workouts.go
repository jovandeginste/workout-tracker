package database

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrInvalidData          = errors.New("could not convert data to a GPX structure")
	ErrWorkoutAlreadyExists = errors.New("user already has workout with exact start time")
)

const minEventDuration = 1 * time.Second

type Workout struct {
	Model
	Date                time.Time            `gorm:"not null;uniqueIndex:idx_start_user" json:"date"`                                    // The timestamp the workout was recorded
	PublicUUID          *uuid.UUID           `gorm:"type:uuid;uniqueIndex" json:"publicUUID"`                                            // UUID to publicly share a workout - this UUID can be rotated
	User                *User                `gorm:"foreignKey:UserID" json:"user"`                                                      // The user who owns the workout
	Data                *MapData             `gorm:"foreignKey:WorkoutID;constraint:OnDelete:CASCADE" json:"data,omitempty"`             // The map data associated with the workout
	GPX                 *GPXData             `gorm:"foreignKey:WorkoutID;constraint:OnDelete:CASCADE" json:"gpx,omitempty"`              // The file data associated with the workout
	Name                string               `gorm:"not null" json:"name"`                                                               // The name of the workout
	Notes               string               `json:"notes"`                                                                              // The notes associated with the workout, in markdown
	Type                WorkoutType          `json:"type"`                                                                               // The type of the workout
	CustomType          string               `json:"custom_type"`                                                                        // The type of the workout, custom
	Equipment           []Equipment          `json:"equipment,omitempty" gorm:"constraint:OnDelete:CASCADE;many2many:workout_equipment"` // Which equipment is used for this workout
	RouteSegmentMatches []*RouteSegmentMatch `gorm:"constraint:OnDelete:CASCADE" json:"routeSegmentMatches,omitempty"`                   // Which route segments match
	UserID              uint64               `gorm:"not null;index;uniqueIndex:idx_start_user" json:"userID"`                            // The ID of the user who owns the workout
	Locked              bool                 `json:"locked"`                                                                             // Whether the workout's main attributes should be auto-updated
	Dirty               bool                 `json:"dirty"`                                                                              // Whether the workout has been modified and the details should be re-rendered
}

type GPXData struct {
	Model
	Filename  string `json:"filename"`                              // The filename of the file
	Content   []byte `gorm:"type:bytes" json:"content"`             // The file content
	Checksum  []byte `gorm:"not null;uniqueIndex" json:"checksum"`  // The checksum of the content
	WorkoutID uint64 `gorm:"not null;uniqueIndex" json:"workoutID"` // The ID of the workout
}

func (w *Workout) HasCustomType() bool {
	return w.Type == WorkoutTypeOther
}

func (w *Workout) AfterFind(tx *gorm.DB) error {
	if w.User != nil {
		w.User.db = tx
	}

	return nil
}

func (w *Workout) GetDate() time.Time {
	return w.Date
}

func (w *Workout) Filename() string {
	if !w.HasFile() {
		return w.Name + ".txt"
	}

	return w.GPX.Filename
}

func (w *Workout) HasElevationData() bool {
	// If both min & max elevation are 0, we don't have elevation information
	return w.MinElevation() != 0 || w.MaxElevation() != 0
}

func (w *Workout) HasPause() bool {
	return w.PauseDuration() == 0
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

func (w *Workout) TotalRepetitions() int {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalRepetitions
}

func (w *Workout) Weight() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalWeight
}

func (w *Workout) AverageSpeed() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.AverageSpeed
}

func (w *Workout) GetEnd() time.Time {
	if w.TotalDuration() <= 0 {
		return w.GetDate().Add(minEventDuration)
	}

	return w.GetDate().Add(w.Duration())
}

func (w *Workout) TotalDuration() time.Duration {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalDuration
}

func (w *Workout) TotalDistance() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalDistance
}

func (w *Workout) Repetitions() int {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalRepetitions
}

func (w *Workout) Duration() time.Duration {
	return w.TotalDuration()
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

func (w *Workout) Center() *MapCenter {
	if w.Data == nil {
		return nil
	}

	return &w.Data.Center
}

func (w *Workout) Details() *MapDataDetails {
	if w.Data == nil {
		return nil
	}

	return w.Data.Details
}

func (w *Workout) TotalDown() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalDown
}

func (w *Workout) TotalUp() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.TotalUp
}

func (w *Workout) MaxElevation() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.MaxElevation
}

func (w *Workout) MinElevation() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.MinElevation
}

func (w *Workout) MaxSpeed() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.MaxSpeed
}

func (w *Workout) MaxCadence() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.MaxCadence
}

func (w *Workout) AverageSpeedNoPause() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.AverageSpeedNoPause
}

func (w *Workout) AverageCadence() float64 {
	if w.Data == nil {
		return 0
	}

	return w.Data.AverageCadence
}

func (w *Workout) PauseDuration() time.Duration {
	if w.Data == nil {
		return 0
	}

	return w.Data.PauseDuration
}

func (w *Workout) Creator() string {
	if w.Data == nil {
		return ""
	}

	return w.Data.Creator
}

func (w *Workout) City() string {
	if w.Data == nil || w.Data.Address == nil {
		return ""
	}

	return w.Data.Address.City
}

func (w *Workout) Timezone() string {
	if w.Data == nil {
		return ""
	}

	return w.Data.Center.TZ
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

func (d *GPXData) Save(db *gorm.DB) error {
	if d.Content == nil {
		return ErrInvalidData
	}

	return db.Save(d).Error
}

func NewWorkout(u *User, workoutType WorkoutType, notes string, filename string, content []byte) ([]*Workout, error) {
	if u == nil {
		return nil, ErrNoUser
	}

	filename = filepath.Base(filename)

	gpxContent, err := converters.ParseCollection(filename, content)
	if err != nil {
		return nil, err
	}

	workouts := make([]*Workout, len(gpxContent))

	for i, g := range gpxContent {
		data := &MapData{
			WorkoutData: g.Data,
		}

		d := &g.Data.Start

		if g.IsGPXBAsed() {
			d = gpxDate(g.GPX)
			data = gpxAsMapData(g.GPX)
		}

		if workoutType == WorkoutTypeAutoDetect {
			workoutType = autoDetectWorkoutType(data, g.GPX, g.Data.Name)
		}

		w := &Workout{
			User:   u,
			UserID: u.ID,
			Dirty:  true,
			Name:   g.Data.Name,
			Data:   data,
			Notes:  notes,
			Type:   workoutType,
			Date:   *d,
		}

		// If multiple GPX files are extracted (e.g., from a zip), use the individual GPX filename.
		if filename == "" || len(gpxContent) > 1 {
			filename = g.Filename()
		}

		w.setContent(filename, g.Content)
		w.UpdateAverages()
		w.UpdateExtraMetrics()

		workouts[i] = w
	}

	return workouts, nil
}

func (w *Workout) setContent(filename string, content []byte) {
	if content == nil {
		return
	}

	h := sha256.New()
	h.Write(content)

	w.GPX = &GPXData{
		Content:  content,
		Checksum: h.Sum(nil),
		Filename: filename,
	}
}

func workoutTypeFromData(gpxType string) (WorkoutType, bool) {
	switch strings.ToLower(gpxType) {
	case "running", "run":
		return WorkoutTypeRunning, true
	case "walking", "walk":
		return WorkoutTypeWalking, true
	case "cycling", "cycle":
		return WorkoutTypeCycling, true
	case "snowboarding":
		return WorkoutTypeSnowboarding, true
	case "horse-riding", "horseback-riding":
		return WorkoutTypeHorseRiding, true
	case "inline-skating", "skating", "skate":
		return WorkoutTypeInlineSkating, true
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
	case "push-ups":
		return WorkoutTypePushups, true
	case "rowing":
		return WorkoutTypeRowing, true
	default:
		return WorkoutTypeAutoDetect, false
	}
}

func autoDetectWorkoutType(data *MapData, gpxContent *gpx.GPX, dataName string) WorkoutType {
	if gpxContent == nil {
		if workoutType, ok := workoutTypeFromData(data.Type); ok {
			return workoutType
		}

		return WorkoutTypeAutoDetect
	}

	// If the GPX file mentions a workout type (for the first track), use it
	if len(gpxContent.Tracks) > 0 {
		firstTrack := &gpxContent.Tracks[0]

		if workoutType, ok := workoutTypeFromData(firstTrack.Type); ok {
			return workoutType
		}
	}

	// If the GPX file mentions a workout type in the name (Runkeeper), use it
	if len(dataName) > 0 {
		nameField := strings.Fields(dataName)
		if len(nameField) > 0 {
			if workoutType, ok := workoutTypeFromData(nameField[0]); ok {
				return workoutType
			}
		}
	}

	if 3.6*data.AverageSpeedNoPause > 15.0 {
		return WorkoutTypeCycling
	}

	if 3.6*data.AverageSpeedNoPause > 7.0 {
		return WorkoutTypeRunning
	}

	return WorkoutTypeWalking
}

func GetRecentWorkouts(db *gorm.DB, count int) ([]*Workout, error) {
	var w []*Workout

	if err := db.Preload("Data").Preload("User").Order("date DESC").Limit(count).Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorkouts(db *gorm.DB) ([]*Workout, error) {
	var w []*Workout

	if err := db.Preload("Data").Preload("Data.Details").Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorkoutDetailsByUUID(db *gorm.DB, u uuid.UUID) (*Workout, error) {
	return GetWorkoutByUUID(db.Preload("GPX").Preload("Data.Details"), u)
}

func GetWorkoutDetails(db *gorm.DB, id uint64) (*Workout, error) {
	return GetWorkout(db.Preload("GPX").Preload("Data.Details"), id)
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

func GetWorkout(db *gorm.DB, id uint64) (*Workout, error) {
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
	return db.Select(clause.Associations).Delete(w).Error
}

func (w *Workout) Create(db *gorm.DB) error {
	err := w.create(db)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return fmt.Errorf("%w: user_id=%d, date=%s", ErrWorkoutAlreadyExists, w.UserID, w.Date)
	}

	return err
}

func (w *Workout) create(db *gorm.DB) error {
	if w.Data == nil {
		return ErrInvalidData
	}

	return db.Create(w).Error
}

func (w *Workout) Save(db *gorm.DB) error {
	err := w.save(db)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrWorkoutAlreadyExists
	}

	return err
}

func (w *Workout) save(db *gorm.DB) error {
	if w.Data == nil {
		return ErrInvalidData
	}

	if w.HasFile() {
		w.UpdateAverages()
	}

	if w.ID == 0 {
		if err := db.Save(w).Error; err != nil {
			return err
		}
	}

	w.Data.WorkoutID = w.ID
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

	wo, err := converters.Parse(w.GPX.Filename, w.GPX.Content)
	if err != nil {
		return nil, err
	}

	return wo.GPX, nil
}

func (w *Workout) setData(data *MapData) {
	if w.Data == nil {
		w.Data = data
		w.Data.WorkoutID = w.ID

		return
	}

	data.ID = w.Data.ID
	data.CreatedAt = w.Data.CreatedAt
	data.WorkoutID = w.ID

	if data.Details == nil {
		data.Details = &MapDataDetails{}
	}

	if w.Data.Details != nil {
		data.Details.ID = w.Data.Details.ID
		data.Details.MapDataID = w.Data.Details.MapDataID
	}

	if w.Locked {
		data.TotalDistance = w.Data.TotalDistance
		data.TotalDuration = w.Data.TotalDuration
		data.Address = w.Data.Address
	}

	data.UpdateAddress()
	data.UpdateExtraMetrics()
	data.CalculateSlopes()
	data.correctNaN()

	w.Data = data
}

func (w *Workout) UpdateAverages() {
	if w.Data == nil {
		return
	}

	w.calculateAverageSpeeds()
	w.calculateCadence()
}

func (w *Workout) calculateAverageSpeeds() {
	w.Data.AverageSpeed = 0
	w.Data.AverageSpeedNoPause = 0

	if w.Data.TotalDuration == 0 {
		return
	}

	w.Data.AverageSpeed = w.Data.TotalDistance / w.Data.TotalDuration.Seconds()

	if w.Data.TotalDuration == w.Data.PauseDuration {
		w.Data.AverageSpeedNoPause = w.Data.AverageSpeed
		return
	}

	w.Data.AverageSpeedNoPause = w.Data.TotalDistance / (w.Data.TotalDuration - w.Data.PauseDuration).Seconds()
}

func (w *Workout) calculateCadence() {
	w.Data.MaxCadence = 0
	w.Data.AverageCadence = 0

	if !w.HasCadence() {
		return
	}

	trackedFor := time.Duration(0)
	avgCadence := 0.0

	for _, p := range w.Data.Details.Points {
		c, ok := p.ExtraMetrics["cadence"]
		if !ok {
			continue
		}

		w.Data.MaxCadence = max(w.Data.MaxCadence, c)
		avgCadence += c * p.Duration.Seconds()
		trackedFor += p.Duration
	}

	if trackedFor.Seconds() == 0 {
		return
	}

	w.Data.AverageCadence = avgCadence / trackedFor.Seconds()
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

	w.UpdateAverages()
	w.UpdateExtraMetrics()
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

	weight := w.User.WeightAt(w.Date)
	// Calories burned = weight * time * intensity (MET)
	cb := weight * w.Duration().Hours() * w.MET()

	return cb
}

func (w *Workout) HasElevation() bool {
	return w.HasExtraMetric("elevation")
}

func (w *Workout) HasEnhancedSpeed() bool {
	return w.HasExtraMetric("speed")
}

func (w *Workout) HasTemperature() bool {
	return w.HasExtraMetric("temperature")
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

func (w *Workout) HasAccuracy() bool {
	return w.HasExtraMetric("accuracy")
}

func (w *Workout) UpdateExtraMetrics() {
	if w.Data == nil {
		return
	}

	w.Data.UpdateExtraMetrics()
}

func (w *Workout) HasExtraMetrics() bool {
	if w.Data == nil {
		return false
	}

	return len(w.Data.ExtraMetrics) > 0
}

func (w *Workout) HasExtraMetric(name string) bool {
	if w.Data == nil {
		return false
	}

	return slices.Contains(w.Data.ExtraMetrics, name)
}

func (w *Workout) EquipmentIDs() []uint64 {
	ids := make([]uint64, 0, len(w.Equipment))

	for _, e := range w.Equipment {
		ids = append(ids, e.ID)
	}

	return ids
}

func (w *Workout) Uses(e Equipment) bool {
	return slices.Contains(w.EquipmentIDs(), e.ID)
}

func (w *Workout) Export() ([]byte, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(w); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
