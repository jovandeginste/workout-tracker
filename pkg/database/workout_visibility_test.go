package database

import (
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func privateTestUser(t *testing.T, db *gorm.DB, username string, admin bool) *User {
	t.Helper()

	u := &User{
		UserData: UserData{
			Username: username,
			Name:     username,
			Active:   true,
			Admin:    admin,
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}

	require.NoError(t, u.Create(db))

	u.SetDB(db)

	return u
}

func privateTestWorkout(t *testing.T, db *gorm.DB, u *User, name string, private bool, date time.Time) *Workout {
	t.Helper()

	w := &Workout{
		UserID:  u.ID,
		Date:    date,
		Name:    name,
		Type:    WorkoutTypeRunning,
		Private: private,
		Data: &MapData{
			WorkoutData: converters.WorkoutData{TotalDistance: 1000, TotalDuration: time.Hour},
		},
	}

	require.NoError(t, w.Create(db))

	return w
}

func TestWorkout_IsVisibleTo(t *testing.T) {
	owner := &User{UserData: UserData{Username: "owner"}}
	owner.ID = 1

	other := &User{UserData: UserData{Username: "other"}}
	other.ID = 2

	admin := &User{UserData: UserData{Username: "admin", Admin: true}}
	admin.ID = 3

	publicWorkout := &Workout{UserID: owner.ID}
	privateWorkout := &Workout{UserID: owner.ID, Private: true}

	tests := []struct {
		name    string
		workout *Workout
		viewer  *User
		want    bool
	}{
		{name: "nil workout", workout: nil, viewer: owner, want: false},
		{name: "public, owner", workout: publicWorkout, viewer: owner, want: true},
		{name: "public, other user", workout: publicWorkout, viewer: other, want: true},
		{name: "public, anonymous", workout: publicWorkout, viewer: nil, want: true},
		{name: "private, owner", workout: privateWorkout, viewer: owner, want: true},
		{name: "private, other user", workout: privateWorkout, viewer: other, want: false},
		{name: "private, admin", workout: privateWorkout, viewer: admin, want: true},
		{name: "private, anonymous", workout: privateWorkout, viewer: nil, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.workout.IsVisibleTo(tt.viewer))
		})
	}
}

func TestVisibleWorkouts(t *testing.T) {
	db := createMemoryDB(t)

	owner := privateTestUser(t, db, "owner", false)
	other := privateTestUser(t, db, "other", false)
	admin := privateTestUser(t, db, "admin", true)

	createPrivateWorkout := privateTestWorkout(t, db, owner, "private workout", true, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	privateTestWorkout(t, db, owner, "public workout", false, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	require.NoError(t, db.Model(&MapData{}).Where("workout_id = ?", createPrivateWorkout.ID).Update("total_distance", 42000).Error)

	names := func(viewer *User) []string {
		t.Helper()

		var workouts []*Workout

		require.NoError(t, VisibleWorkouts(db, viewer).Order("name").Find(&workouts).Error)

		result := make([]string, 0, len(workouts))
		for _, w := range workouts {
			result = append(result, w.Name)
		}

		return result
	}

	assert.Equal(t, []string{"private workout", "public workout"}, names(owner))
	assert.Equal(t, []string{"public workout"}, names(other))
	assert.Equal(t, []string{"private workout", "public workout"}, names(admin))
	assert.Equal(t, []string{"public workout"}, names(nil))

	// Combined with another filter, the visibility condition must stay grouped.
	var ownerWorkouts []*Workout

	require.NoError(t, VisibleWorkouts(db.Where(&Workout{UserID: owner.ID}), other).Order("name").Find(&ownerWorkouts).Error)
	require.Len(t, ownerWorkouts, 1)
	assert.Equal(t, "public workout", ownerWorkouts[0].Name)
}

func TestRouteSegment_FilterVisibleMatches(t *testing.T) {
	owner := &User{UserData: UserData{Username: "owner"}}
	owner.ID = 1

	other := &User{UserData: UserData{Username: "other"}}
	other.ID = 2

	newSegment := func() *RouteSegment {
		return &RouteSegment{
			RouteSegmentMatches: []*RouteSegmentMatch{
				{Workout: &Workout{UserID: owner.ID, Name: "public"}},
				{Workout: &Workout{UserID: owner.ID, Name: "private", Private: true}},
				nil,
			},
		}
	}

	rs := newSegment()
	rs.FilterVisibleMatches(other)
	require.Len(t, rs.RouteSegmentMatches, 1)
	assert.Equal(t, "public", rs.RouteSegmentMatches[0].Workout.Name)

	rs = newSegment()
	rs.FilterVisibleMatches(owner)
	assert.Len(t, rs.RouteSegmentMatches, 2)
}

func TestStatistics_AreViewerAware(t *testing.T) {
	db := createMemoryDB(t)

	owner := privateTestUser(t, db, "owner", false)
	other := privateTestUser(t, db, "other", false)
	admin := privateTestUser(t, db, "boss", true)

	privateWorkout := privateTestWorkout(t, db, owner, "private workout", true, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	privateTestWorkout(t, db, owner, "public workout", false, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	require.NoError(t, db.Model(&MapData{}).Where("workout_id = ?", privateWorkout.ID).Update("total_distance", 42000).Error)

	owner.Profile.TotalsShow = WorkoutTypeList{WorkoutTypeRunning}

	totals := func(viewer *User) *Bucket {
		t.Helper()

		buckets, err := owner.GetDefaultTotals(viewer)
		require.NoError(t, err)
		require.Len(t, buckets, 1)

		return buckets[0]
	}

	assert.Equal(t, 2, totals(owner).Workouts)
	assert.Equal(t, 43000.0, totals(owner).Distance)

	assert.Equal(t, 2, totals(admin).Workouts)
	assert.Equal(t, 43000.0, totals(admin).Distance)

	assert.Equal(t, 1, totals(other).Workouts)
	assert.Equal(t, 1000.0, totals(other).Distance)

	record := func(viewer *User) *WorkoutRecord {
		t.Helper()

		r, err := owner.GetRecords(WorkoutTypeRunning, viewer)
		require.NoError(t, err)

		return r
	}

	assert.Equal(t, 42000.0, record(owner).Distance.Value)
	assert.Equal(t, 42000.0, record(admin).Distance.Value)
	assert.Equal(t, 1000.0, record(other).Distance.Value)
	assert.Equal(t, 1000.0, record(nil).Distance.Value)

	records, err := owner.GetAllRecords(other)
	require.NoError(t, err)
	assert.NotEmpty(t, records)

	for _, r := range records {
		assert.NotEqual(t, 42000.0, r.Distance.Value)
	}
}
