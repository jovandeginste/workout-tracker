package database

import (
	"crypto/sha256"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name     string     `gorm:"nut null"`
	Date     *time.Time `gorm:"not null"`
	UserID   uint       `gorm:"not null;index"`
	Dirty    bool
	User     *User
	Notes    string
	Type     string
	Data     MapData `gorm:"serializer:json"`
	Checksum []byte  `gorm:"not null;uniqueIndex"`
	GPXData  []byte  `gorm:"type:mediumtext"`
}

func NewWorkout(u *User, workoutType, notes string, content []byte) *Workout {
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
	return db.Create(w).Error
}

func (w *Workout) Save(db *gorm.DB) error {
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
