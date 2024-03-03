package database

import (
	"crypto/sha256"
	"errors"
	"time"

	"github.com/jovandeginste/workout-tracker/pkg/converters"
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

var ErrInvalidData = errors.New("could not convert data to a GPX structure")

type Workout struct {
	gorm.Model
	Name   string     `gorm:"not null"`
	Date   *time.Time `gorm:"not null"`
	UserID uint       `gorm:"not null;index"`
	Dirty  bool
	User   *User
	Notes  string
	Type   WorkoutType
	Data   *MapData `gorm:"serializer:json"`
	GPX    *GPXData

	GPXData  []byte `gorm:"type:mediumtext"`
	Filename string
	Checksum []byte
}

type GPXData struct {
	gorm.Model
	WorkoutID uint   `gorm:"not null;uniqueIndex"`
	Content   []byte `gorm:"type:mediumtext"`
	Checksum  []byte `gorm:"not null;uniqueIndex"`
	Filename  string
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
		workoutType = autoDetectWorkoutType(data)
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

func autoDetectWorkoutType(data *MapData) WorkoutType {
	if 3.6*data.AverageSpeedNoPause() > 15.0 {
		return WorkoutTypeCycling
	}

	if 3.6*data.AverageSpeedNoPause() > 5.0 {
		return WorkoutTypeRunning
	}

	return WorkoutTypeWalking
}

func GetRecentWorkouts(db *gorm.DB, count int) ([]Workout, error) {
	var w []Workout

	if err := db.Preload("User").Order("date DESC").Limit(count).Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorkouts(db *gorm.DB) ([]*Workout, error) {
	var w []*Workout

	if err := db.Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w Workout

	if err := db.Preload("User").First(&w, id).Error; err != nil {
		return nil, err
	}

	return &w, nil
}

func (w *Workout) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(w).Error
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

func (w *Workout) UpdateData(db *gorm.DB) error {
	gpxContent, err := w.AsGPX()
	if err != nil {
		return err
	}

	w.Data = gpxAsMapData(gpxContent)
	w.Dirty = false

	return db.Save(w).Error
}
