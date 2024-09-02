package database

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

type Config struct {
	gorm.Model

	// These options can be changed at runtime, configured through the database
	// If they are set through the environment to a non-default value, that will
	// take precedence
	RegistrationDisabled bool `mapstructure:"registration_disabled" form:"registration_disabled"`
	SocialsDisabled      bool `mapstructure:"socials_disabled" form:"socials_disabled"`

	// These options are read from the config file or environment only
	Logging          bool   `mapstructure:"logging" gorm:"-"`
	Debug            bool   `mapstructure:"debug" gorm:"-"`
	Bind             string `mapstructure:"bind" gorm:"-"`
	JWTEncryptionKey string `mapstructure:"jwt_encryption_key" gorm:"-"`
	DatabaseDriver   string `mapstructure:"database_driver" gorm:"-"`
	DSN              string `mapstructure:"dsn" gorm:"-"`
}

func getConfig(db *gorm.DB) (*Config, error) {
	var c Config

	if err := db.Model(&Config{}).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) UpdateFromDatabase(db *gorm.DB) error {
	dbConfig, err := getConfig(db)
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
	v := reflect.ValueOf(c).Elem()
	vFrom := reflect.ValueOf(from).Elem()
	n := v.Type().NumField()

	for i := range n {
		if v.Field(i).IsZero() {
			v.Field(i).Set(vFrom.Field(i))
		}
	}
}

func (c *Config) Save(db *gorm.DB) error {
	realCnf, err := getConfig(db)
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
