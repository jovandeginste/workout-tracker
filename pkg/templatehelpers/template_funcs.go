package templatehelpers

import (
	"fmt"
	"html/template"
	"time"

	"github.com/dustin/go-humanize"
	emojiflag "github.com/jayco/go-emoji-flag"
)

func NumericDuration(d time.Duration) float64 {
	return d.Seconds()
}

func CountryCodeToFlag(cc string) string {
	return emojiflag.GetFlag(cc)
}

func LocalDate(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04") //nolint:gosmopolitan
}

func HumanDistance(d float64) string {
	value, prefix := humanize.ComputeSI(d)

	return fmt.Sprintf("%.2f %sm", value, prefix)
}

func HumanSpeed(mps float64) string {
	mph := mps * 3600
	value, prefix := humanize.ComputeSI(mph)

	return fmt.Sprintf("%.2f %sm/h", value, prefix)
}

func HumanTempo(mps float64) string {
	mpk := 1000000 / (mps * 60)
	value, prefix := humanize.ComputeSI(mpk)

	return fmt.Sprintf("%.2f min/%sm", value, prefix)
}

func BoolToHTML(b bool) template.HTML {
	if b {
		return `<i class="text-green-500 fas fa-check"></i>`
	}

	return `<i class="text-rose-500 fas fa-times"></i>`
}

func BoolToCheckbox(b bool) template.HTML {
	if b {
		return "checked"
	}

	return ""
}

func BuildDecoratedAttribute(icon, name string, value interface{}) interface{} {
	return struct {
		Icon  string
		Name  string
		Value interface{}
	}{
		Icon:  icon,
		Name:  name,
		Value: value,
	}
}
