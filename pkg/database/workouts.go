package database

import (
	"crypto/sha256"
	"errors"
	"slices"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

type WorkoutType string

const (
	WorkoutTypeAutoDetect WorkoutType = "auto"
	WorkoutTypeRunning    WorkoutType = "running"
	WorkoutTypeCycling    WorkoutType = "cycling"
	WorkoutTypeWalking    WorkoutType = "walking"
)

var ErrInvalidGPXData = errors.New("invalid gpx data")

func WorkoutTypes() []WorkoutType {
	return []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking}
}

func DistanceWorkoutTypes() []WorkoutType {
	return []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking}
}

func (wt WorkoutType) String() string {
	return string(wt)
}

func (wt WorkoutType) IsDistance() bool {
	return slices.Contains(DistanceWorkoutTypes(), wt)
}

type Workout struct {
	gorm.Model
	Name     string     `gorm:"nut null"`
	Date     *time.Time `gorm:"not null"`
	UserID   uint       `gorm:"not null;index"`
	Dirty    bool
	User     *User
	Notes    string
	Type     WorkoutType
	Data     *MapData `gorm:"serializer:json"`
	Checksum []byte   `gorm:"not null;uniqueIndex"`
	GPXData  []byte   `gorm:"type:mediumtext"`
}

func NewWorkout(u *User, workoutType WorkoutType, notes string, content []byte) *Workout {
	if u == nil {
		return nil
	}

	gpxContent, err := parseGPX(content)
	if err != nil {
		return nil
	}

	data := gpxAsMapData(gpxContent)

	h := sha256.New()
	h.Write(content)

	if workoutType == WorkoutTypeAutoDetect {
		workoutType = autoDetectWorkoutType(data)
	}

	w := Workout{
		User:     u,
		UserID:   u.ID,
		GPXData:  content,
		Name:     gpxName(gpxContent),
		Data:     data,
		Notes:    notes,
		Type:     workoutType,
		Date:     gpxContent.Time,
		Checksum: h.Sum(nil),
	}

	return &w
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
		return ErrInvalidGPXData
	}

	return db.Create(w).Error
}

func (w *Workout) Save(db *gorm.DB) error {
	if w.Data == nil {
		return ErrInvalidGPXData
	}

	return db.Save(w).Error
}

func (w *Workout) AsGPX() (*gpx.GPX, error) {
	return parseGPX(w.GPXData)
}

func (w *Workout) UpdateData(db *gorm.DB) error {
	gpxContent, err := parseGPX(w.GPXData)
	if err != nil {
		return err
	}

	w.Data = gpxAsMapData(gpxContent)
	w.Dirty = false

	return db.Save(w).Error
}
