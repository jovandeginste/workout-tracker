package user

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
	User     *User
	Notes    string
	Type     string
	Checksum []byte `gorm:"not null;uniqueIndex"`
	GPXData  []byte `gorm:"type:mediumtext"`
}

func NewWorkout(u *User, workoutType, notes string, content []byte) *Workout {
	if u == nil {
		return nil
	}

	gpxContent, err := parseGPX(content)
	if err != nil {
		return nil
	}

	h := sha256.New()
	h.Write(content)

	w := Workout{
		User:     u,
		UserID:   u.ID,
		GPXData:  content,
		Name:     gpxName(gpxContent),
		Notes:    notes,
		Type:     workoutType,
		Date:     gpxContent.Time,
		Checksum: h.Sum(nil),
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
