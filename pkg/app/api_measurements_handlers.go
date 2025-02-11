package app

import (
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
)

type Measurement struct {
	Date       string  `form:"date" json:"date"`               // The date of the measurement
	Weight     float64 `form:"weight" json:"weight"`           // The weight of the user, in kilograms
	Height     float64 `form:"height" json:"height"`           // The height of the user, in centimeter
	Steps      float64 `form:"steps" json:"steps"`             // The number of steps taken
	WeightUnit string  `form:"weight_unit" json:"weight_unit"` // The unit of the weight (or the user's preferred unit)
	HeightUnit string  `form:"height_unit" json:"height_unit"` // The unit of the height (or the user's preferred unit)

	units *database.UserPreferredUnits
}

func (m *Measurement) Time() time.Time {
	if m.Date == "" {
		return time.Now()
	}

	d, err := time.Parse("2006-01-02", m.Date)
	if err != nil {
		return time.Now()
	}

	return d
}

func (m *Measurement) ToHeight() *float64 {
	if m.Height == 0 {
		return nil
	}

	if m.HeightUnit == "" {
		m.HeightUnit = m.units.Height()
	}

	d := templatehelpers.HeightToDatabase(m.Height, m.HeightUnit)

	return &d
}

func (m *Measurement) ToWeight() *float64 {
	if m.Weight == 0 {
		return nil
	}

	if m.WeightUnit == "" {
		m.WeightUnit = m.units.Weight()
	}

	d := templatehelpers.WeightToDatabase(m.Weight, m.WeightUnit)

	return &d
}

func (m *Measurement) Update(measurement *database.Measurement) {
	setIfNotNil(&measurement.Weight, m.ToWeight())
	setIfNotNil(&measurement.Height, m.ToHeight())
	setIfNotNil(&measurement.Steps, &m.Steps)
}
