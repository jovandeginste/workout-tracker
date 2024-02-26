package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserIsValid(t *testing.T) {
	u := User{
		Username: "my-username",
		Password: "my-password",
		Active:   true,
	}

	require.NoError(t, u.IsValid())
	assert.True(t, u.IsActive())
}

func TestUserPasswordIsValid(t *testing.T) {
	pwd := "my-password"
	u := User{
		Username: "my-username",
		Name:     "my-name",
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
		Username: "my-username",
		Password: "my-password",
		Active:   false,
	}

	require.NoError(t, u.IsValid())
	assert.False(t, u.IsActive())
}

func TestUserUsernameIsEmail(t *testing.T) {
	u := User{
		Username: "my-username@localhost",
		Password: "my-password",
	}

	require.NoError(t, u.IsValid())
	assert.False(t, u.IsActive())
}

func TestUserUsernameIsNotValid(t *testing.T) {
	for _, username := range []string{
		"invalid-char-;",
		"invalid-char-@",
		"invalid-char-<script>",
		"invalid-char space",
	} {
		u := User{
			Username: username,
			Password: "my-password",
			Name:     "my-name",
		}

		require.ErrorIs(t, u.IsValid(), ErrUsernameInvalid)
		assert.False(t, u.IsActive())
	}
}

func TestUserUsernameIsTooLong(t *testing.T) {
	u := User{
		Username: "too-long-too-long-too-long-too-long-too-long-too-long-too-long-too-long",
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
