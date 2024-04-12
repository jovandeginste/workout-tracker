package database

import (
	"crypto/sha256"
	"errors"
	"html/template"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jovandeginste/workout-tracker/pkg/converters"
	"github.com/microcosm-cc/bluemonday"
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

var ErrInvalidData = errors.New("could not convert data to a GPX structure")

type Workout struct {
	gorm.Model
	Name    string     `gorm:"not null"`
	Date    *time.Time `gorm:"not null"`
	UserID  uint       `gorm:"not null;index"`
	Dirty   bool
	User    *User
	Notes   string
	Type    WorkoutType
	Data    *MapData `json:",omitempty"`
	GPX     *GPXData `json:",omitempty"`
	MapData *MapData `gorm:"serializer:json;column:data" json:"-"`

	GPXData  []byte `gorm:"type:text" json:"-"` // To be removed
	Filename string `json:"-"`                        // To be removed
	Checksum []byte `gorm:"default:legacy" json:"-"`  // To be removed
}

type GPXData struct {
	gorm.Model
	WorkoutID uint   `gorm:"not null;uniqueIndex"`
	Content   []byte `gorm:"type:text"`
	Checksum  []byte `gorm:"not null;uniqueIndex"`
	Filename  string
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

	gpxContent, err := converters.Parse(filename, content)
	if err != nil {
		return nil, err
	}

	data := gpxAsMapData(gpxContent)

	h := sha256.New()
	h.Write(content)

	if workoutType == WorkoutTypeAutoDetect {
		workoutType = autoDetectWorkoutType(data, gpxContent)
	}

	w := Workout{
		User:   u,
		UserID: u.ID,
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

func GetWorkoutWithGPX(db *gorm.DB, id int) (*Workout, error) {
	return GetWorkout(db.Preload("GPX"), id)
}

func GetWorkoutDetails(db *gorm.DB, id int) (*Workout, error) {
	return GetWorkout(db.Preload("Data.Details"), id)
}

func GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w Workout

	if err := db.Preload("Data").Preload("User").First(&w, id).Error; err != nil {
		return nil, err
	}

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

	return db.Save(w).Error
}

func (w *Workout) AsGPX() (*gpx.GPX, error) {
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
}

func (w *Workout) UpdateData(db *gorm.DB) error {
	gpxContent, err := w.AsGPX()
	if err != nil {
		return err
	}

	w.setData(gpxAsMapData(gpxContent))

	if err := w.Data.Save(db); err != nil {
		return err
	}

	w.Dirty = false

	return w.Save(db)
}
