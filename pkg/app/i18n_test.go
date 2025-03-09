package app

import (
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestTranslateWorkoutTypes(t *testing.T) {
	a := defaultApp(t)
	a.ConfigureLocalizer()

	wt := database.WorkoutTypes()
	tr := a.translator

	for _, w := range wt {
		t.Run(
			"translation of "+w.String(),
			func(t *testing.T) {
				assert.True(t, tr.Has(w.String()))
			},
		)
	}
}
