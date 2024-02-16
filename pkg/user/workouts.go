package user

import (
	"time"

	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name    string     `gorm:"nut null"`
	Date    *time.Time `gorm:"not null"`
	UserID  uint       `gorm:"not null;index"`
	User    *User
	Notes   string
	Type    string
	GPXData []byte `gorm:"type:mediumtext"`
}

func NewWorkout(u *User, notes string, content []byte) *Workout {
	if u == nil {
		return nil
	}

	gpxContent, err := parseGPX(content)
	if err != nil {
		return nil
	}

	w := Workout{
		User:    u,
		UserID:  u.ID,
		GPXData: content,
		Name:    gpxName(gpxContent),
		Notes:   notes,
		Date:    gpxContent.Time,
	}

	return &w
}

func (w *Workout) Create(db *gorm.DB) error {
	return db.Create(w).Error
}

func (w *Workout) AsGPX() (*gpx.GPX, error) {
	return parseGPX(w.GPXData)
}

func (w *Workout) MapData() MapData {
	gpxContent, _ := w.AsGPX()
	return gpxAsMapData(gpxContent)
}
