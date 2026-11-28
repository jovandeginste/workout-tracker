package auth

import "context"

// Identity is the information an external authentication provider returns
// after a successful login.
type Identity struct {
	Username string
	Name     string
	Admin    bool
}

// Authenticator verifies a username and password against an external source.
type Authenticator interface {
	Authenticate(ctx context.Context, username, password string) (*Identity, error)
}
