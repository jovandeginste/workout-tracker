package auth

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"testing"

	"github.com/go-ldap/ldap/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type bindCall struct {
	username string
	password string
}

type fakeConn struct {
	bindCalls  []bindCall
	bindFunc   func(username, password string) error
	searchFunc func(request *ldap.SearchRequest) (*ldap.SearchResult, error)
}

func (c *fakeConn) Bind(username, password string) error {
	c.bindCalls = append(c.bindCalls, bindCall{username: username, password: password})
	if c.bindFunc != nil {
		return c.bindFunc(username, password)
	}

	return nil
}

func (c *fakeConn) Search(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
	if c.searchFunc != nil {
		return c.searchFunc(request)
	}

	return &ldap.SearchResult{}, nil
}

func (c *fakeConn) StartTLS(*tls.Config) error {
	return nil
}

func (c *fakeConn) Close() error {
	return nil
}

func searchResult(entries ...*ldap.Entry) *ldap.SearchResult {
	return &ldap.SearchResult{Entries: entries}
}

func ldapEntry(dn string, attributes map[string][]string) *ldap.Entry {
	entry := &ldap.Entry{DN: dn}

	for name, values := range attributes {
		entry.Attributes = append(entry.Attributes, &ldap.EntryAttribute{Name: name, Values: values})
	}

	return entry
}

func newTestLDAP(t *testing.T, cfg LDAPConfig, c *fakeConn) *LDAP {
	t.Helper()

	cfg.URL = "ldaps://ldap.example.com:636"

	l, err := newLDAP(cfg, func(LDAPConfig, *x509.CertPool, string) (conn, error) {
		return c, nil
	})
	require.NoError(t, err)

	return l
}

func TestNewLDAPValidation(t *testing.T) {
	validDirect := LDAPConfig{
		URL:    "ldaps://ldap.example.com:636",
		UserDN: "uid={username},ou=people,dc=example,dc=com",
	}

	validSearch := LDAPConfig{
		URL:          "ldaps://ldap.example.com:636",
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
	}

	tests := []struct {
		name string
		cfg  LDAPConfig
	}{
		{name: "missing URL", cfg: LDAPConfig{UserDN: validDirect.UserDN}},
		{name: "bad scheme", cfg: LDAPConfig{URL: "http://ldap.example.com", UserDN: validDirect.UserDN}},
		{name: "start TLS on ldaps", cfg: LDAPConfig{URL: "ldaps://ldap.example.com:636", StartTLS: true, UserDN: validDirect.UserDN}},
		{name: "no mode", cfg: LDAPConfig{URL: validDirect.URL}},
		{name: "both modes", cfg: LDAPConfig{URL: validDirect.URL, UserDN: validDirect.UserDN, BindDN: validSearch.BindDN, BindPassword: validSearch.BindPassword}},
		{name: "search without password", cfg: LDAPConfig{URL: validSearch.URL, BindDN: validSearch.BindDN, BaseDN: validSearch.BaseDN, UserFilter: validSearch.UserFilter}},
		{name: "search without base", cfg: LDAPConfig{URL: validSearch.URL, BindDN: validSearch.BindDN, BindPassword: validSearch.BindPassword, UserFilter: validSearch.UserFilter}},
		{name: "search without placeholder", cfg: LDAPConfig{URL: validSearch.URL, BindDN: validSearch.BindDN, BindPassword: validSearch.BindPassword, BaseDN: validSearch.BaseDN, UserFilter: "(uid=jo)"}},
		{name: "search with two placeholders", cfg: LDAPConfig{URL: validSearch.URL, BindDN: validSearch.BindDN, BindPassword: validSearch.BindPassword, BaseDN: validSearch.BaseDN, UserFilter: "(uid={username})(cn={username})"}},
		{name: "direct without placeholder", cfg: LDAPConfig{URL: validDirect.URL, UserDN: "uid=jo,ou=people,dc=example,dc=com"}},
		{name: "invalid required group", cfg: LDAPConfig{URL: validDirect.URL, UserDN: validDirect.UserDN, RequiredGroup: "not a dn"}},
		{name: "invalid admin group", cfg: LDAPConfig{URL: validDirect.URL, UserDN: validDirect.UserDN, AdminGroup: "not a dn"}},
		{name: "missing CA file", cfg: LDAPConfig{URL: validDirect.URL, UserDN: validDirect.UserDN, CAFile: "/does/not/exist"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLDAP(tt.cfg)
			require.Error(t, err)
		})
	}

	t.Run("valid direct mode", func(t *testing.T) {
		_, err := NewLDAP(validDirect)
		require.NoError(t, err)
	})

	t.Run("valid search mode", func(t *testing.T) {
		_, err := NewLDAP(validSearch)
		require.NoError(t, err)
	})
}

func TestLDAPAuthenticateSearchMode(t *testing.T) {
	c := &fakeConn{
		searchFunc: func(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
			assert.Equal(t, "ou=people,dc=example,dc=com", request.BaseDN)
			assert.Equal(t, "(uid=jo)", request.Filter)

			return searchResult(ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{
				"uid":      {"jo"},
				"cn":       {"Jo Vandeginste"},
				"memberOf": {"cn=workout-tracker-admins,ou=groups,dc=example,dc=com"},
			})), nil
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		BindDN:            "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword:      "service-password",
		BaseDN:            "ou=people,dc=example,dc=com",
		UserFilter:        "(uid={username})",
		UsernameAttribute: "uid",
		NameAttribute:     "cn",
		AdminGroup:        "cn=workout-tracker-admins,ou=groups,dc=example,dc=com",
	}, c)

	identity, err := l.Authenticate(context.Background(), "jo", "user-password")
	require.NoError(t, err)
	assert.Equal(t, &Identity{Username: "jo", Name: "Jo Vandeginste", Admin: true}, identity)
	assert.Equal(t, []bindCall{
		{username: "cn=workout-tracker,ou=services,dc=example,dc=com", password: "service-password"},
		{username: "uid=jo,ou=people,dc=example,dc=com", password: "user-password"},
	}, c.bindCalls)
}

func TestLDAPAuthenticateEmptyPassword(t *testing.T) {
	c := &fakeConn{}

	l := newTestLDAP(t, LDAPConfig{
		UserDN: "uid={username},ou=people,dc=example,dc=com",
	}, c)

	_, err := l.Authenticate(context.Background(), "jo", "")
	require.ErrorIs(t, err, ErrInvalidCredentials)
	assert.Empty(t, c.bindCalls)
}

func TestLDAPAuthenticateUnknownUser(t *testing.T) {
	c := &fakeConn{
		searchFunc: func(*ldap.SearchRequest) (*ldap.SearchResult, error) {
			return searchResult(), nil
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
	}, c)

	_, err := l.Authenticate(context.Background(), "nobody", "password")
	require.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLDAPAuthenticateMultipleUsers(t *testing.T) {
	c := &fakeConn{
		searchFunc: func(*ldap.SearchRequest) (*ldap.SearchResult, error) {
			return searchResult(
				ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{"uid": {"jo"}}),
				ldapEntry("uid=jo,ou=other,dc=example,dc=com", map[string][]string{"uid": {"jo"}}),
			), nil
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
	}, c)

	_, err := l.Authenticate(context.Background(), "jo", "password")
	require.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLDAPAuthenticateInvalidUserPassword(t *testing.T) {
	c := &fakeConn{
		bindFunc: func(username, _ string) error {
			if username == "uid=jo,ou=people,dc=example,dc=com" {
				return ldap.NewError(ldap.LDAPResultInvalidCredentials, errors.New("invalid credentials"))
			}

			return nil
		},
		searchFunc: func(*ldap.SearchRequest) (*ldap.SearchResult, error) {
			return searchResult(ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{"uid": {"jo"}})), nil
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
	}, c)

	_, err := l.Authenticate(context.Background(), "jo", "wrong-password")
	require.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLDAPAuthenticateFilterEscaping(t *testing.T) {
	var filter string

	c := &fakeConn{
		searchFunc: func(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
			filter = request.Filter
			return searchResult(), nil
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
	}, c)

	_, err := l.Authenticate(context.Background(), "*)(uid=*", "password")
	require.ErrorIs(t, err, ErrInvalidCredentials)
	assert.Equal(t, "(uid="+ldap.EscapeFilter("*)(uid=*")+")", filter)
}

func TestLDAPAuthenticateRequiredGroup(t *testing.T) {
	cfg := LDAPConfig{
		BindDN:        "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword:  "service-password",
		BaseDN:        "ou=people,dc=example,dc=com",
		UserFilter:    "(uid={username})",
		RequiredGroup: "cn=workout-tracker-users,ou=groups,dc=example,dc=com",
	}

	t.Run("member", func(t *testing.T) {
		c := &fakeConn{
			searchFunc: func(*ldap.SearchRequest) (*ldap.SearchResult, error) {
				return searchResult(ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{
					"uid":      {"jo"},
					"memberOf": {"cn=workout-tracker-users,ou=groups,dc=example,dc=com"},
				})), nil
			},
		}

		_, err := newTestLDAP(t, cfg, c).Authenticate(context.Background(), "jo", "password")
		require.NoError(t, err)
	})

	t.Run("not a member", func(t *testing.T) {
		c := &fakeConn{
			searchFunc: func(*ldap.SearchRequest) (*ldap.SearchResult, error) {
				return searchResult(ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{
					"uid": {"jo"},
				})), nil
			},
		}

		_, err := newTestLDAP(t, cfg, c).Authenticate(context.Background(), "jo", "password")
		require.ErrorIs(t, err, ErrInvalidCredentials)
	})
}

func TestLDAPAuthenticateAdminGroup(t *testing.T) {
	cfg := LDAPConfig{
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
		AdminGroup:   "CN=Workout Admins,OU=Groups,DC=example,DC=com",
	}

	c := &fakeConn{
		searchFunc: func(*ldap.SearchRequest) (*ldap.SearchResult, error) {
			return searchResult(ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{
				"uid":      {"jo"},
				"memberOf": {"cn=workout admins,ou=groups,dc=example,dc=com"},
			})), nil
		},
	}

	identity, err := newTestLDAP(t, cfg, c).Authenticate(context.Background(), "jo", "password")
	require.NoError(t, err)
	assert.True(t, identity.Admin)
}

func TestLDAPAuthenticateDirectBind(t *testing.T) {
	c := &fakeConn{
		searchFunc: func(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
			assert.Equal(t, "uid=jo,ou=people,dc=example,dc=com", request.BaseDN)
			assert.Equal(t, ldap.ScopeBaseObject, request.Scope)

			return searchResult(ldapEntry("uid=jo,ou=people,dc=example,dc=com", map[string][]string{
				"uid": {"jo"},
				"cn":  {"Jo"},
			})), nil
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		UserDN:            "uid={username},ou=people,dc=example,dc=com",
		UsernameAttribute: "uid",
		NameAttribute:     "cn",
	}, c)

	identity, err := l.Authenticate(context.Background(), "jo", "password")
	require.NoError(t, err)
	assert.Equal(t, &Identity{Username: "jo", Name: "Jo"}, identity)
	assert.Equal(t, []bindCall{
		{username: "uid=jo,ou=people,dc=example,dc=com", password: "password"},
	}, c.bindCalls)
}

func TestLDAPAuthenticateServiceBindFailure(t *testing.T) {
	c := &fakeConn{
		bindFunc: func(string, string) error {
			return errors.New("service bind failed")
		},
	}

	l := newTestLDAP(t, LDAPConfig{
		BindDN:       "cn=workout-tracker,ou=services,dc=example,dc=com",
		BindPassword: "wrong-service-password",
		BaseDN:       "ou=people,dc=example,dc=com",
		UserFilter:   "(uid={username})",
	}, c)

	_, err := l.Authenticate(context.Background(), "jo", "password")
	require.ErrorIs(t, err, ErrLDAPUnavailable)
}

func TestLDAPAuthenticateNetworkFailure(t *testing.T) {
	l := newTestLDAP(t, LDAPConfig{
		UserDN: "uid={username},ou=people,dc=example,dc=com",
	}, &fakeConn{})
	l.dial = func(LDAPConfig, *x509.CertPool, string) (conn, error) {
		return nil, errors.New("network failure")
	}

	_, err := l.Authenticate(context.Background(), "jo", "password")
	require.ErrorIs(t, err, ErrLDAPUnavailable)
}
