package user

import (
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name    string `gorm:"nut null"`
	UserID  uint   `gorm:"not null;index"`
	User    *User
	Notes   string
	Type    string
	GPXData string `gorm:"type:mediumtext"`
}

func NewWorkout(u *User, notes string, content string, gpxContent *gpx.GPX) *Workout {
	if u == nil {
		return nil
	}

	w := Workout{
		User:    u,
		UserID:  u.ID,
		GPXData: content,
		Name:    gpxContent.Name,
		Notes:   notes,
	}

	return &w
}

func (w *Workout) Create(db *gorm.DB) error {
	return db.Create(w).Error
}
