package database

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

type Config struct {
	Model
	EnvConfig  `mapstructure:",squash"`
	UserConfig `mapstructure:",squash"`

	// Overrides tracks which user settings are managed by the configuration
	// file or environment. They cannot be changed through the web interface.
	Overrides ConfigOverrides `mapstructure:"-" gorm:"-"`
}

// ConfigOverrides marks the user settings that are overridden by the
// configuration file or environment.
type ConfigOverrides struct {
	RegistrationDisabled bool
	SocialsDisabled      bool
	LoginCustomText      bool
}

// UserConfig are options that can be changed at runtime or configured through the web interface
// If they are set through the environment to a non-default value, that will
// take precedence
type UserConfig struct {
	RegistrationDisabled bool   `mapstructure:"registration_disabled" form:"registration_disabled"`
	SocialsDisabled      bool   `mapstructure:"socials_disabled" form:"socials_disabled"`
	LoginCustomText      string `mapstructure:"login_custom_text" form:"login_custom_text"` // Markdown text shown on the login screen
}

// EnvConfig are options that are read from the config file or environment only
type EnvConfig struct {
	Bind               string `mapstructure:"bind" gorm:"-"`                 // Which address to bind to
	WebRoot            string `mapstructure:"web_root" gorm:"-"`             // The web root path (relative to the bind address)
	JWTEncryptionKey   string `mapstructure:"jwt_encryption_key" gorm:"-"`   // Encryption key for JWT
	Dev                bool   `mapstructure:"dev" gorm:"-"`                  // Development mode
	DatabaseDriver     string `mapstructure:"database_driver" gorm:"-"`      // Which database driver to use
	DSN                string `mapstructure:"dsn" gorm:"-"`                  // Database DSN
	Logging            bool   `mapstructure:"logging" gorm:"-"`              // Enable logging
	Debug              bool   `mapstructure:"debug" gorm:"-"`                // Debug logging mode
	Offline            bool   `mapstructure:"offline" gorm:"-"`              // Disable calls to external services
	Compress           bool   `mapstructure:"compress" gorm:"-"`             // Enable gzip compression
	WorkerDelaySeconds int    `mapstructure:"worker_delay_seconds" gorm:"-"` // Time in seconds between worker runs

	JWTEncryptionKeyFile string `mapstructure:"jwt_encryption_key_file" gorm:"-"` // File containing the encryption key for JWT
	DSNFile              string `mapstructure:"dsn_file" gorm:"-"`                // File containing the database DSN

	LDAPEnabled            bool   `mapstructure:"ldap_enabled" gorm:"-"`              // Enable LDAP authentication
	LDAPURL                string `mapstructure:"ldap_url" gorm:"-"`                  // LDAP server URL
	LDAPStartTLS           bool   `mapstructure:"ldap_start_tls" gorm:"-"`            // Use StartTLS on an ldap:// URL
	LDAPInsecureSkipVerify bool   `mapstructure:"ldap_insecure_skip_verify" gorm:"-"` // Skip LDAP TLS certificate verification
	LDAPCAFile             string `mapstructure:"ldap_ca_file" gorm:"-"`              // PEM file with a private CA
	LDAPBindDN             string `mapstructure:"ldap_bind_dn" gorm:"-"`              // Service account DN for search mode
	LDAPBindPassword       string `mapstructure:"ldap_bind_password" gorm:"-"`        // Service account password
	LDAPBindPasswordFile   string `mapstructure:"ldap_bind_password_file" gorm:"-"`   // File with the service account password
	LDAPUserDN             string `mapstructure:"ldap_user_dn" gorm:"-"`              // DN template for direct-bind mode
	LDAPBaseDN             string `mapstructure:"ldap_base_dn" gorm:"-"`              // Search base for search mode
	LDAPUserFilter         string `mapstructure:"ldap_user_filter" gorm:"-"`          // Search filter with a {username} placeholder
	LDAPUsernameAttribute  string `mapstructure:"ldap_username_attribute" gorm:"-"`   // Attribute used as the app username
	LDAPNameAttribute      string `mapstructure:"ldap_name_attribute" gorm:"-"`       // Attribute used as the display name
	LDAPRequiredGroup      string `mapstructure:"ldap_required_group" gorm:"-"`       // DN of a group the user must be a member of
	LDAPAdminGroup         string `mapstructure:"ldap_admin_group" gorm:"-"`          // DN of a group whose members become admins
	LDAPAutoCreate         bool   `mapstructure:"ldap_auto_create" gorm:"-"`          // Create unknown LDAP users on login
	LDAPDefaultActive      bool   `mapstructure:"ldap_default_active" gorm:"-"`       // Active value for newly created LDAP users
	LDAPTimeoutSeconds     int    `mapstructure:"ldap_timeout_seconds" gorm:"-"`      // LDAP dial and operation timeout
}

func GetConfig(db *gorm.DB) (*Config, error) {
	var c Config

	if err := db.Model(&Config{}).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) UpdateFromDatabase(db *gorm.DB) error {
	dbConfig, err := GetConfig(db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	c.copy(dbConfig)

	return nil
}

// copy copies all fields from one Config to another, if the destination field
// is the zero value (false, 0, "")
func (c *Config) copy(from *Config) {
	copyZeroFields(&c.Model, &from.Model)
	copyZeroFields(&c.EnvConfig, &from.EnvConfig)
	copyZeroFields(&c.UserConfig, &from.UserConfig)
}

// copyZeroFields copies the fields of src to dst, if the destination field is
// the zero value.
func copyZeroFields[T any](dst, src *T) {
	v := reflect.ValueOf(dst).Elem()
	vFrom := reflect.ValueOf(src).Elem()

	for i := range v.NumField() {
		if !v.Field(i).CanSet() || !v.Field(i).IsZero() {
			continue
		}

		v.Field(i).Set(vFrom.Field(i))
	}
}

func (c *Config) Save(db *gorm.DB) error {
	realCnf, err := GetConfig(db)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		c.ID = realCnf.ID
		c.CreatedAt = realCnf.CreatedAt
	}

	return db.Save(c).Error
}
