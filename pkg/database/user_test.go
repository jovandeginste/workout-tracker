package database

import (
	"testing"

	"github.com/fsouza/slognil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func init() { //nolint:gochecknoinits
	online = false
}

func defaultUser() *User {
	return &User{
		UserData: UserData{
			Username: "my-username",
			Name:     "my-name",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}
}

func dummyMapData() *MapData {
	return &MapData{Creator: "tester"}
}

func createMemoryDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := Connect("memory", "", false, slognil.NewLogger())
	require.NoError(t, err)

	return db
}

func createDefaultUser(t *testing.T, db *gorm.DB) {
	t.Helper()

	require.NoError(t, defaultUser().Create(db))
}

func TestUser_IsValid(t *testing.T) {
	u := defaultUser()
	u.Active = true

	require.NoError(t, u.IsValid())
	assert.True(t, u.IsActive())
}

func TestUser_PasswordIsValid(t *testing.T) {
	pwd := "my-password"
	u := defaultUser()
	u.Active = true
	u.Password = ""

	require.Error(t, u.IsValid())
	assert.Empty(t, u.Salt)
	assert.Empty(t, u.Password)

	require.NoError(t, u.SetPassword(pwd))

	require.NoError(t, u.IsValid())
	assert.NotEmpty(t, u.Salt)
	assert.NotEmpty(t, u.Password)
	require.NotEqual(t, u.Password, pwd)

	assert.True(t, u.ValidLogin(pwd))
	assert.False(t, u.ValidLogin(pwd+pwd))
}

func TestUser_IsNotActive(t *testing.T) {
	u := User{
		UserData: UserData{
			Username: "my-username",
			Active:   false,
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}

	require.NoError(t, u.IsValid())
	assert.False(t, u.IsActive())
}

func TestUser_UsernameIsEmail(t *testing.T) {
	u := User{
		UserData: UserData{
			Username: "my-username@localhost",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}

	require.NoError(t, u.IsValid())
	assert.False(t, u.IsActive())
}

func TestUser_UsernameIsNotValid(t *testing.T) {
	for _, username := range []string{
		"invalid-char-;",
		"invalid-char-@",
		"invalid-char-<script>",
		"invalid-char space",
	} {
		u := User{
			UserData: UserData{
				Username: username,
				Name:     "my-name",
			},
			UserSecrets: UserSecrets{
				Password: "my-password",
			},
		}

		require.ErrorIs(t, u.IsValid(), ErrUsernameInvalid)
		assert.False(t, u.IsActive())
	}
}

func TestUser_UsernameIsTooLong(t *testing.T) {
	u := User{
		UserData: UserData{
			Username: "too-long-too-long-too-long-too-long-too-long-too-long-too-long-too-long",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}

	require.ErrorIs(t, u.IsValid(), ErrUsernameInvalidLength)
	assert.False(t, u.IsActive())
}

func TestUser_PasswordNotSet(t *testing.T) {
	u := User{
		UserData: UserData{
			Username: "username",
		},
		UserSecrets: UserSecrets{
			Password: "",
		},
	}

	require.ErrorIs(t, u.IsValid(), ErrPasswordInvalidLength)
	assert.False(t, u.IsActive())
}

func TestUser_BeforeCreateNoPassword(t *testing.T) {
	db := createMemoryDB(t)
	u := &User{
		UserData: UserData{
			Username: "username",
		},
		UserSecrets: UserSecrets{
			Password: "",
		},
	}

	require.Error(t, u.Create(db))
	assert.NotEmpty(t, u.Salt)

	u, err := GetUser(db, "other-username")
	require.NoError(t, err)
	require.Nil(t, u)
}

func TestDatabaseUserCreate(t *testing.T) {
	db := createMemoryDB(t)
	u := &User{
		UserData: UserData{
			Username: "username",
			Name:     "my-name",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}

	require.NoError(t, u.Create(db))
	require.NoError(t, u.IsValid())
	assert.False(t, u.IsActive())
	assert.NotEmpty(t, u.Salt)
	assert.NotEmpty(t, u.ID)

	u, err := GetUser(db, "username")
	require.NoError(t, err)
	assert.Equal(t, "my-name", u.Name)

	u, err = GetUserByID(db, u.ID)
	require.NoError(t, err)
	assert.Equal(t, "my-name", u.Name)
}

func TestDatabaseUsers(t *testing.T) {
	db := createMemoryDB(t)

	u1 := User{
		UserData: UserData{
			Username: "username1",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}
	require.NoError(t, u1.Create(db))

	users, err := GetUsers(db)
	require.NoError(t, err)
	assert.Len(t, users, 1)

	u2 := User{
		UserData: UserData{
			Username: "username2",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}
	require.NoError(t, u2.Create(db))

	users, err = GetUsers(db)
	require.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestDatabaseUserSave(t *testing.T) {
	db := createMemoryDB(t)
	u := defaultUser()

	require.NoError(t, u.Create(db))

	u, err := GetUser(db, "my-username")
	require.NoError(t, err)
	assert.Equal(t, "my-name", u.Name)

	u.Name = "other-name"
	require.NoError(t, u.Save(db))

	u, err = GetUser(db, "my-username")
	require.NoError(t, err)
	assert.Equal(t, "other-name", u.Name)
}

func TestDatabaseUserCreateDoubleUsername(t *testing.T) {
	db := createMemoryDB(t)
	createDefaultUser(t, db)

	err := defaultUser().Create(db)
	require.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrDuplicatedKey)

	users, err := GetUsers(db)
	require.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestDatabaseUserDeleteUser(t *testing.T) {
	db := createMemoryDB(t)

	u := defaultUser()
	require.NoError(t, u.Create(db))

	users, err := GetUsers(db)
	require.NoError(t, err)
	assert.Len(t, users, 1)

	require.NoError(t, u.Delete(db))

	users, err = GetUsers(db)
	require.NoError(t, err)
	assert.Empty(t, users)
}

func TestDatabaseProfileSave(t *testing.T) {
	db := createMemoryDB(t)
	u := &User{
		UserData: UserData{
			Username: "username",
			Name:     "my-name",
		},
		UserSecrets: UserSecrets{
			Password: "my-password",
		},
	}
	u.Profile.Language = "en"

	require.NoError(t, u.Create(db))
	assert.NotEmpty(t, u.Profile.ID)

	u, err := GetUser(db, "username")
	require.NoError(t, err)
	assert.Equal(t, "en", u.Profile.Language)

	u.Profile.Language = "de"
	require.NoError(t, u.Profile.Save(db))
	u, err = GetUser(db, "username")
	require.NoError(t, err)
	assert.Equal(t, "de", u.Profile.Language)
}

func TestDatabaseUserWorkouts(t *testing.T) {
	populateGPXFS()

	db := createMemoryDB(t)

	u := defaultUser()
	require.NoError(t, u.Create(db))

	workouts, err := u.GetWorkouts(db)
	require.NoError(t, err)
	assert.Empty(t, workouts)

	w1, addErr := u.AddWorkout(
		db,
		WorkoutTypeAutoDetect,
		"some notes",
		"file.gpx",
		[]byte("invalid content"),
	)
	require.NotEmpty(t, addErr)
	require.ErrorIs(t, addErr[0], ErrInvalidData)
	assert.Nil(t, w1)

	workouts, err = u.GetWorkouts(db)
	require.NoError(t, err)
	assert.Empty(t, workouts)

	f1, err := gpxFS.ReadFile("sample1.gpx")
	require.NoError(t, err)

	w2, addErr := u.AddWorkout(
		db,
		WorkoutTypeAutoDetect,
		"some notes",
		"file.gpx",
		f1,
	)
	require.Empty(t, addErr)
	assert.Len(t, w2, 1)
	w2_1 := w2[0]

	workouts, err = u.GetWorkouts(db)
	require.NoError(t, err)
	assert.Len(t, workouts, 1)

	assert.True(t, w2_1.HasElevation())
	assert.True(t, w2_1.HasHeartRate())

	w2_1.Type = WorkoutTypeWalking
	require.NoError(t, w2_1.Save(db))
}
