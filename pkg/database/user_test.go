package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserIsValid(t *testing.T) {
	u := User{
		Username: "my-username@localhost",
		Password: "my-password",
		Active:   true,
	}

	require.NoError(t, u.IsValid())
	assert.True(t, u.IsActive())
}

func TestUserPasswordIsValid(t *testing.T) {
	pwd := "my-password"
	u := User{
		Username: "my-username@localhost",
		Active:   true,
	}

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

func TestUserIsNotActive(t *testing.T) {
	u := User{
		Username: "my-username@localhost",
		Password: "my-password",
		Active:   false,
	}

	require.NoError(t, u.IsValid())
	assert.False(t, u.IsActive())
}

func TestUserUsernameIsNotEmail(t *testing.T) {
	u := User{
		Username: "my-username",
		Password: "my-password",
	}

	require.ErrorIs(t, u.IsValid(), ErrUsernameInvalid)
	assert.False(t, u.IsActive())
}

func TestUserUsernameIsTooLong(t *testing.T) {
	u := User{
		Username: "diezi6moo1oogo9kohth9Zu3aethahF4@localhost",
		Password: "my-password",
	}

	require.ErrorIs(t, u.IsValid(), ErrUsernameInvalidLength)
	assert.False(t, u.IsActive())
}

func TestUserPasswordNotSet(t *testing.T) {
	u := User{
		Username: "username@localhost",
		Password: "",
	}

	require.ErrorIs(t, u.IsValid(), ErrPasswordInvalidLength)
	assert.False(t, u.IsActive())
}
