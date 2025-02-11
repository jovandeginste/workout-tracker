package database

import "time"

type Measurement struct {
	Model
	User   *User      `gorm:"foreignKey:UserID"`                                    // The user who owns the workout
	Date   *time.Time `json:"date" gorm:"not null;index;uniqueIndex:idx_user_date"` // The date of the measurement
	Weight float64    `json:"weight"`                                               // The weight of the user, in kilograms
	Steps  uint64     `json:"steps"`                                                // The number of steps taken
	UserID uint64     `gorm:"not null;index;uniqueIndex:idx_user_date"`             // The ID of the user who owns the workout
}
