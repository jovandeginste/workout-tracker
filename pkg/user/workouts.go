package user

import (
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name    string
	UserID  uint `gorm:"not null;index"`
	User    *User
	Notes   string
	Type    string
	GPXData string `gorm:"type:mediumtext"`
}

func NewWorkout(u *User, notes string, gpxContent *gpx.GPX) *Workout {
	if u == nil {
		return nil
	}

	w := Workout{
		User:    u,
		UserID:  u.ID,
		GPXData: gpxContent.XMLNs,
		Name:    gpxContent.Name,
		Notes:   notes,
	}

	return &w
}

func (w *Workout) Create(db *gorm.DB) error {
	return db.Create(w).Error
}
