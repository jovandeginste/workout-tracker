package user

import (
	"crypto/sha256"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name       string     `gorm:"nut null"`
	Date       *time.Time `gorm:"not null"`
	UserID     uint       `gorm:"not null;index"`
	User       *User
	Notes      string
	Type       string
	Data       MapData `gorm:"serializer:json"`
	Checksum   []byte  `gorm:"not null;uniqueIndex"`
	GPXData    []byte  `gorm:"type:mediumtext"`
	FAIconName string
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
		User:       u,
		UserID:     u.ID,
		GPXData:    content,
		Name:       gpxName(gpxContent),
		Data:       data,
		Notes:      notes,
		Type:       workoutType,
		Date:       gpxContent.Time,
		Checksum:   h.Sum(nil),
		FAIconName: faIconNameFor(workoutType),
	}

	return &w
}

func (w *Workout) Create(db *gorm.DB) error {
	return db.Create(w).Error
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

	return db.Save(w).Error
}

func faIconNameFor(wType string) string {
	if wType == "running" {
		return "person-running"
	}

	return "question"
}
