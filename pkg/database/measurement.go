package database

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Measurement struct {
	Model
	Date   datatypes.Date `form:"date" json:"date" gorm:"not null;index;uniqueIndex:idx_user_date"` // The date of the measurement
	Weight float64        `form:"weight" json:"weight"`                                             // The weight of the user, in kilograms
	Height uint64         `form:"height" json:"height"`                                             // The height of the user, in centimeter
	Steps  uint64         `form:"steps" json:"steps"`                                               // The number of steps taken
	UserID uint64         `gorm:"not null;index;uniqueIndex:idx_user_date"`                         // The ID of the user who owns the workout
}

func (u *User) NewMeasurement(d time.Time) *Measurement {
	return &Measurement{
		UserID: u.ID,
		Date:   datatypes.Date(d.UTC()),
	}
}

func (m *Measurement) Save(db *gorm.DB) error {
	return db.Save(m).Error
}

func (m *Measurement) Time() *time.Time {
	t := time.Time(m.Date)
	return &t
}

func (m *Measurement) Delete(db *gorm.DB) error {
	return db.Delete(m).Error
}

func (m *Measurement) DateString() string {
	return m.Time().Format("2006-01-02")
}
