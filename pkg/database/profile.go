package database

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID     uint
	APIActive  bool        `form:"api_active"`
	Language   string      `form:"language"`
	TotalsShow WorkoutType `form:"totals_show"`

	User *User `gorm:"foreignKey:UserID"`
}

func (p *Profile) Save(db *gorm.DB) error {
	return db.Save(p).Error
}
