package database

import (
	"fmt"
	"os"

	"github.com/jovandeginste/workout-tracker/internal/pkg/templatehelpers"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID              uint        // The ID of the user who owns this profile
	APIActive           bool        `form:"api_active"`            // Whether the user's API key is active
	Language            string      `form:"language"`              // The user's preferred language
	TotalsShow          WorkoutType `form:"totals_show"`           // What workout type of totals to show
	Timezone            string      `form:"timezone"`              // The user's preferred timezone
	AutoImportDirectory string      `form:"auto_import_directory"` // The user's preferred directory for auto-import
	SocialsDisabled     bool        `form:"socials_disabled"`      // Whether social sharing buttons are disabled when viewing a workout
	PreferFullDate      bool        `form:"prefer_full_date"`      // Whether to show full dates in the workout details

	PreferredUnits UserPreferredUnits `gorm:"serializer:json"` // The user's preferred units

	User *User `gorm:"foreignKey:UserID" json:"-"` // The user who owns this profile
}

type UserPreferredUnits struct {
	SpeedRaw     string `form:"speed" json:"speed"`         // The user's preferred speed unit
	DistanceRaw  string `form:"distance" json:"distance"`   // The user's preferred distance unit
	ElevationRaw string `form:"elevation" json:"elevation"` // The user's preferred elevation unit
	WeightRaw    string `form:"weight" json:"weight"`       // The user's preferred weight unit
}

func (u UserPreferredUnits) Tempo() string {
	return "min/" + u.Distance()
}

func (u UserPreferredUnits) HeartRate() string {
	return "bpm"
}

func (u UserPreferredUnits) Cadence() string {
	return "spm"
}

func (u UserPreferredUnits) Elevation() string {
	switch u.ElevationRaw {
	case "ft":
		return "ft"
	default:
		return "m"
	}
}

func (u UserPreferredUnits) Weight() string {
	switch u.WeightRaw {
	case "lbs":
		return "lbs"
	default:
		return "kg"
	}
}

func (u UserPreferredUnits) Distance() string {
	switch u.DistanceRaw {
	case "mi":
		return "mi"
	default:
		return "km"
	}
}

func (u UserPreferredUnits) DistanceToDatabase(d float64) float64 {
	switch u.Distance() {
	case "mi":
		return d * templatehelpers.MeterPerMile
	default:
		return d * templatehelpers.MeterPerKM
	}
}

func (u UserPreferredUnits) Speed() string {
	switch u.SpeedRaw {
	case "mph":
		return "mph"
	default:
		return "km/h"
	}
}

func (p *Profile) ResetBools() {
	p.PreferFullDate = false
	p.APIActive = false
	p.SocialsDisabled = false
}

func (p *Profile) Save(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Profile) CanImportFromDirectory() (bool, error) {
	if p == nil {
		return false, nil
	}

	if p.AutoImportDirectory == "" {
		return false, nil
	}

	info, err := os.Stat(p.AutoImportDirectory)
	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("%v is not a directory", p.AutoImportDirectory)
	}

	return true, nil
}
