package auth

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrLDAPUnavailable    = errors.New("LDAP unavailable")
)

const (
	defaultLDAPTimeout           = 10 * time.Second
	defaultLDAPUsernameAttribute = "uid"
	defaultLDAPNameAttribute     = "cn"
	ldapUsernamePlaceholder      = "{username}"
	ldapNameMaximumLength        = 64
)

// LDAPConfig contains the connection and lookup settings for an LDAP
// directory. Use either BindDN (search mode) or UserDN (direct-bind mode),
// never both.
type LDAPConfig struct {
	URL                string
	StartTLS           bool
	InsecureSkipVerify bool
	CAFile             string
	BindDN             string
	BindPassword       string
	UserDN             string
	BaseDN             string
	UserFilter         string
	UsernameAttribute  string
	NameAttribute      string
	RequiredGroup      string
	AdminGroup         string
	Timeout            time.Duration
}

// LDAP authenticates users against an LDAP directory.
type LDAP struct {
	cfg    LDAPConfig
	roots  *x509.CertPool
	server string
	dial   dialFunc
}

type conn interface {
	Bind(username, password string) error
	Search(*ldap.SearchRequest) (*ldap.SearchResult, error)
	StartTLS(*tls.Config) error
	Close() error
}

type dialFunc func(cfg LDAPConfig, roots *x509.CertPool, server string) (conn, error)

// NewLDAP validates the configuration and returns an LDAP authenticator.
func NewLDAP(cfg LDAPConfig) (*LDAP, error) {
	return newLDAP(cfg, dialLDAP)
}

func newLDAP(cfg LDAPConfig, dial dialFunc) (*LDAP, error) {
	server, err := validateLDAPURL(cfg)
	if err != nil {
		return nil, err
	}

	if err := validateLDAPMode(cfg); err != nil {
		return nil, err
	}

	cfg = normalizeLDAPConfig(cfg)

	if err := validateLDAPGroups(cfg); err != nil {
		return nil, err
	}

	roots, err := loadRootCAs(cfg.CAFile)
	if err != nil {
		return nil, err
	}

	return &LDAP{
		cfg:    cfg,
		roots:  roots,
		server: server,
		dial:   dial,
	}, nil
}

func validateLDAPURL(cfg LDAPConfig) (string, error) {
	if cfg.URL == "" {
		return "", errors.New("ldap_url is required")
	}

	u, err := url.Parse(cfg.URL)
	if err != nil || (u.Scheme != "ldap" && u.Scheme != "ldaps") || u.Host == "" {
		return "", errors.New("ldap_url must be a valid ldap:// or ldaps:// URL")
	}

	if cfg.StartTLS && u.Scheme != "ldap" {
		return "", errors.New("ldap_start_tls requires an ldap:// URL")
	}

	return u.Hostname(), nil
}

func validateLDAPMode(cfg LDAPConfig) error {
	if (cfg.BindDN == "") == (cfg.UserDN == "") {
		return errors.New("set exactly one of ldap_bind_dn or ldap_user_dn")
	}

	if cfg.BindDN == "" {
		if strings.Count(cfg.UserDN, ldapUsernamePlaceholder) != 1 {
			return errors.New("ldap_user_dn must contain exactly one {username} placeholder")
		}

		return nil
	}

	if cfg.BindPassword == "" {
		return errors.New("ldap_bind_password or ldap_bind_password_file is required with ldap_bind_dn")
	}

	if cfg.BaseDN == "" || cfg.UserFilter == "" {
		return errors.New("ldap_base_dn and ldap_user_filter are required with ldap_bind_dn")
	}

	if strings.Count(cfg.UserFilter, ldapUsernamePlaceholder) != 1 {
		return errors.New("ldap_user_filter must contain exactly one {username} placeholder")
	}

	return nil
}

func normalizeLDAPConfig(cfg LDAPConfig) LDAPConfig {
	if cfg.UsernameAttribute == "" {
		cfg.UsernameAttribute = defaultLDAPUsernameAttribute
	}

	if cfg.NameAttribute == "" {
		cfg.NameAttribute = defaultLDAPNameAttribute
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = defaultLDAPTimeout
	}

	return cfg
}

func validateLDAPGroups(cfg LDAPConfig) error {
	if err := validateOptionalDN(cfg.RequiredGroup); err != nil {
		return fmt.Errorf("ldap_required_group: %w", err)
	}

	if err := validateOptionalDN(cfg.AdminGroup); err != nil {
		return fmt.Errorf("ldap_admin_group: %w", err)
	}

	return nil
}

func validateOptionalDN(value string) error {
	if value == "" {
		return nil
	}

	_, err := ldap.ParseDN(value)

	return err
}

func loadRootCAs(filename string) (*x509.CertPool, error) {
	if filename == "" {
		return nil, nil
	}

	pem, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read ldap_ca_file: %w", err)
	}

	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(pem) {
		return nil, errors.New("ldap_ca_file contains no certificates")
	}

	return roots, nil
}

func dialLDAP(cfg LDAPConfig, roots *x509.CertPool, server string) (conn, error) {
	tlsConfig := &tls.Config{
		ServerName:         server,
		RootCAs:            roots,
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: cfg.InsecureSkipVerify, //nolint:gosec // explicit opt-in
	}

	c, err := ldap.DialURL(
		cfg.URL,
		ldap.DialWithDialer(&net.Dialer{Timeout: cfg.Timeout}),
		ldap.DialWithTLSConfig(tlsConfig),
	)
	if err != nil {
		return nil, err
	}

	c.SetTimeout(cfg.Timeout)

	if cfg.StartTLS {
		if err := c.StartTLS(tlsConfig); err != nil {
			_ = c.Close()
			return nil, err
		}
	}

	return c, nil
}

// Authenticate verifies the user's password against the directory and returns
// the attributes the app needs.
func (l *LDAP) Authenticate(ctx context.Context, username, password string) (*Identity, error) {
	if username == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	c, err := l.dial(l.cfg, l.roots, l.server)
	if err != nil {
		return nil, fmt.Errorf("%w: dial: %v", ErrLDAPUnavailable, err)
	}
	defer c.Close()

	var entry *ldap.Entry

	if l.cfg.BindDN != "" {
		if err := c.Bind(l.cfg.BindDN, l.cfg.BindPassword); err != nil {
			return nil, fmt.Errorf("%w: service bind: %v", ErrLDAPUnavailable, err)
		}

		entry, err = l.searchUser(c, username)
		if err != nil {
			return nil, err
		}

		if entry == nil {
			return nil, ErrInvalidCredentials
		}

		if err := c.Bind(entry.DN, password); err != nil {
			return nil, userBindError(err)
		}
	} else {
		dn := strings.ReplaceAll(l.cfg.UserDN, ldapUsernamePlaceholder, ldap.EscapeDN(username))
		if err := c.Bind(dn, password); err != nil {
			return nil, userBindError(err)
		}

		entry, err = l.readEntry(c, dn)
		if err != nil {
			return nil, err
		}

		if entry == nil {
			return nil, ErrInvalidCredentials
		}
	}

	return l.identity(username, entry)
}

func userBindError(err error) error {
	if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
		return ErrInvalidCredentials
	}

	return fmt.Errorf("%w: user bind: %v", ErrLDAPUnavailable, err)
}

func (l *LDAP) searchUser(c conn, username string) (*ldap.Entry, error) {
	filter := strings.ReplaceAll(l.cfg.UserFilter, ldapUsernamePlaceholder, ldap.EscapeFilter(username))

	result, err := c.Search(ldap.NewSearchRequest(
		l.cfg.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		2,
		0,
		false,
		filter,
		l.attributes(),
		nil,
	))
	if err != nil {
		return nil, fmt.Errorf("%w: search: %v", ErrLDAPUnavailable, err)
	}

	switch len(result.Entries) {
	case 0:
		return nil, nil
	case 1:
		return result.Entries[0], nil
	default:
		return nil, ErrInvalidCredentials
	}
}

func (l *LDAP) readEntry(c conn, dn string) (*ldap.Entry, error) {
	result, err := c.Search(ldap.NewSearchRequest(
		dn,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		1,
		0,
		false,
		"(objectClass=*)",
		l.attributes(),
		nil,
	))
	if err != nil {
		return nil, fmt.Errorf("%w: read entry: %v", ErrLDAPUnavailable, err)
	}

	if len(result.Entries) == 0 {
		return nil, nil
	}

	return result.Entries[0], nil
}

func (l *LDAP) attributes() []string {
	attributes := make([]string, 0, 3)

	for _, attribute := range []string{l.cfg.UsernameAttribute, l.cfg.NameAttribute, "memberOf"} {
		if !slices.Contains(attributes, attribute) {
			attributes = append(attributes, attribute)
		}
	}

	return attributes
}

func (l *LDAP) identity(login string, entry *ldap.Entry) (*Identity, error) {
	if l.cfg.RequiredGroup != "" && !l.isMember(entry, l.cfg.RequiredGroup) {
		return nil, ErrInvalidCredentials
	}

	username := strings.TrimSpace(entry.GetAttributeValue(l.cfg.UsernameAttribute))
	if username == "" {
		username = login
	}

	name := strings.TrimSpace(entry.GetAttributeValue(l.cfg.NameAttribute))
	if name == "" {
		name = username
	}

	if len([]rune(name)) > ldapNameMaximumLength {
		name = string([]rune(name)[:ldapNameMaximumLength])
	}

	return &Identity{
		Username: username,
		Name:     name,
		Admin:    l.isMember(entry, l.cfg.AdminGroup),
	}, nil
}

func (l *LDAP) isMember(entry *ldap.Entry, group string) bool {
	if group == "" {
		return false
	}

	want, err := ldap.ParseDN(group)
	if err != nil {
		return false
	}

	for _, value := range entry.GetAttributeValues("memberOf") {
		got, err := ldap.ParseDN(value)
		if err == nil && got.EqualFold(want) {
			return true
		}
	}

	return false
}
