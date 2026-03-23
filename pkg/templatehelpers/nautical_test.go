package templatehelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanDistanceNM(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "0"},
		{1852, "1"},
		{3704, "2"},
		{100, "0.05"},
		{2000, "1.08"},
	}

	for _, tt := range tests {
		result := HumanDistanceNM(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestHumanSpeedKnots(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "N/A"},
		{1, "1.94"},
		{5.14444, "10"},
		{10, "19.44"},
	}

	for _, tt := range tests {
		result := HumanSpeedKnots(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestHumanTempoNM(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "N/A"},
		{2.5, "12:20"},      // 12 min 20 sec per NM
		{5.0, "6:10"},       // 6 min 10 sec per NM
		{0.514444, "60:00"}, // 1 knot = 60 min/NM
	}

	for _, tt := range tests {
		result := HumanTempoNM(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestHumanTempoFor(t *testing.T) {
	tests := []struct {
		unit     string
		input    float64
		expected string
	}{
		{"min/mi", 2.68224, "10:00"},
		{"mi", 2.68224, "10:00"},
		{"min/nm", 5.0, "6:10"},
		{"nm", 5.0, "6:10"},
		{"min/km", 2.77778, "5:59"},
		{"km", 2.77778, "5:59"},
	}

	for _, tt := range tests {
		f := HumanTempoFor(tt.unit)
		result := f(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestHumanDistanceFor(t *testing.T) {
	tests := []struct {
		unit     string
		input    float64
		expected string
	}{
		{"mi", 1609.34, "1"},
		{"nm", 1852, "1"},
		{"km", 1000, "1"},
		{"m", 1000, "1"},
	}

	for _, tt := range tests {
		f := HumanDistanceFor(tt.unit)
		result := f(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestHumanSpeedFor(t *testing.T) {
	tests := []struct {
		unit     string
		input    float64
		expected string
	}{
		{"mph", 10, "22.37"},
		{"kn", 10, "19.44"},
		{"km/h", 10, "36"},
	}

	for _, tt := range tests {
		f := HumanSpeedFor(tt.unit)
		result := f(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}
