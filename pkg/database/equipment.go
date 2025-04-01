package database

import (
	"slices"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Equipment struct {
	db *gorm.DB
	Model

	Name        string        `gorm:"not null;uniqueIndex" json:"name" form:"name"`                            // The name of the gear
	Description string        `gorm:"" json:"description" form:"description"`                                  // More information about the equipment
	DefaultFor  []WorkoutType `gorm:"serializer:json;column:default_for" form:"default_for" json:"defaultFor"` // Which workout types to add this equipment by default

	Workouts []Workout `gorm:"many2many:workout_equipment" json:"workouts"`

	User User `json:"user"`

	UserID uint64 `gorm:"not null;index" json:"userID"`             // The ID of the user who owns the workout
	Active bool   `gorm:"default:true" json:"active" form:"active"` // Whether this equipment is active
}

type WorkoutEquipment struct {
	Model
	Workout     Workout   `json:"workout"`
	Equipment   Equipment `json:"equipment"`
	WorkoutID   uint64    `gorm:"not null;uniqueIndex:idx_workout_equipment" json:"workoutID"`   // The ID of the workout
	EquipmentID uint64    `gorm:"not null;uniqueIndex:idx_workout_equipment" json:"equipmentID"` // The ID of the equipment
}

func GetEquipment(db *gorm.DB, id int) (*Equipment, error) {
	var e Equipment

	if err := db.Preload("User").Preload("Workouts").Preload("Workouts.Data").First(&e, id).Error; err != nil {
		return nil, err
	}

	e.SetDB(db)

	return &e, nil
}

func (e *Equipment) ValidFor(wt *WorkoutType) bool {
	return slices.Contains(e.DefaultFor, *wt)
}

func (e *Equipment) Delete(db *gorm.DB) error {
	if err := db.Model(e).Association("Workouts").Clear(); err != nil {
		return err
	}

	return db.Select("workout_equipment").Delete(e).Error
}

func (e *Equipment) Save(db *gorm.DB) error {
	return db.Omit(clause.Associations).Save(e).Error
}

func GetEquipmentByIDs(db *gorm.DB, userID uint64, ids []uint64) ([]*Equipment, error) {
	var equipment []*Equipment

	if len(ids) == 0 {
		return equipment, nil
	}

	if err := db.Where("user_id = ?", userID).Find(&equipment, ids).Error; err != nil {
		return nil, err
	}

	return equipment, nil
}

func (e *Equipment) SetDB(db *gorm.DB) {
	e.db = db
}

func (e *Equipment) GetTotals() (WorkoutTotals, error) {
	rs := WorkoutTotals{}

	for _, w := range e.Workouts {
		if w.Type.IsDistance() {
			rs.Distance += w.Distance()
		}

		if w.Type.IsDuration() {
			rs.Duration += w.Duration()
		}

		if w.Type.IsRepetition() {
			rs.Repetitions += w.Repetitions()
		}
	}

	return rs, nil
}

type WorkoutTotals struct {
	Distance    float64
	Duration    time.Duration
	Repetitions int
}
