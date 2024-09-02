package database

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/cat-dealer/go-rand/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PasswordMinimumLength = 4
	PasswordMaximumLength = 128
	UsernameMinimumLength = 1
	UsernameMaximumLength = 32
)

var (
	ErrPasswordInvalidLength = errors.New("password has invalid length")
	ErrUsernameInvalidLength = errors.New("username has invalid length")
	ErrUsernameInvalid       = errors.New("username is not valid")
	ErrNoUser                = errors.New("no user attached")
)

type User struct {
	gorm.Model

	LastVersion string `gorm:"last_version"` // Which version of the app the user has last seen and acknowledged

	Password string `form:"-"        gorm:"type:varchar(128);not null"`            // The user's password as bcrypt hash
	Salt     string `form:"-"        gorm:"type:varchar(16);not null"`             // The salt used to hash the user's password
	Username string `form:"username" gorm:"uniqueIndex;not null;type:varchar(32)"` // The user's username
	Name     string `form:"name"     gorm:"type:varchar(64);not null"`             // The user's name
	APIKey   string `gorm:"type:varchar(32)"`                                      // The user's API key
	Active   bool   `form:"active"`                                                // Whether the user is active
	Admin    bool   `form:"admin"`                                                 // Whether the user is an admin

	Profile   Profile     // The user's profile settings
	Workouts  []Workout   `json:"-"` // The user's workouts
	Equipment []Equipment `json:"-"` // The user's equipment

	anonymous bool // Whether we have an actual user or not

	db *gorm.DB
}

func AnonymousUser() *User {
	return &User{anonymous: true}
}

func (u *User) IsAnonymous() bool {
	if u == nil {
		return true
	}

	return u.anonymous
}

func (u *User) ShowFullDate() bool {
	if u == nil {
		return false
	}

	return u.Profile.PreferFullDate
}

func (u *User) PreferredUnits() *UserPreferredUnits {
	if u == nil {
		return &UserPreferredUnits{}
	}

	return &u.Profile.PreferredUnits
}

func (u *User) Timezone() *time.Location {
	if u == nil || u.Profile.Timezone == "" {
		return time.UTC
	}

	loc, err := time.LoadLocation(u.Profile.Timezone)
	if err != nil {
		return time.UTC
	}

	return loc
}

func (u *User) BeforeSave(_ *gorm.DB) error {
	u.GenerateAPIKey(false)
	u.GenerateSalt()

	return u.IsValid()
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, db.Error
	}

	return u, nil
}

func currentUserQuery(db *gorm.DB) *gorm.DB {
	return db.Preload("Profile").Preload("Equipment")
}

func GetUserByAPIKey(db *gorm.DB, key string) (*User, error) {
	var u User

	if err := currentUserQuery(db).Where(&User{APIKey: key}).First(&u).Error; err != nil {
		return nil, db.Error
	}

	u.SetDB(db)

	return &u, nil
}

func GetUserByID(db *gorm.DB, userID int) (*User, error) {
	var u User

	if err := currentUserQuery(db).First(&u, userID).Error; err != nil {
		return nil, db.Error
	}

	u.SetDB(db)

	return &u, nil
}

func GetUser(db *gorm.DB, username string) (*User, error) {
	var u User

	if err := currentUserQuery(db).Where(&User{Username: username}).First(&u).Error; err != nil {
		return nil, db.Error
	}

	if u.ID == u.Profile.UserID {
		u.Profile.User = &u
	}

	u.SetDB(db)

	return &u, nil
}

func (u *User) SetDB(db *gorm.DB) {
	u.db = db
}

func (u *User) APIActive() bool {
	if u == nil {
		return false
	}

	if !u.Profile.APIActive {
		return false
	}

	return u.APIKey != ""
}

func (u *User) IsActive() bool {
	if u == nil {
		return false
	}

	if !u.Active || u.Password == "" || u.Username == "" {
		return false
	}

	return true
}

func (u *User) ValidLogin(password string) bool {
	if !u.IsActive() {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.AddSalt(password))) == nil
}

func (u *User) AddSalt(password string) string {
	return u.Salt + password
}

func (u *User) IsValid() error {
	if u.Password == "" {
		return ErrPasswordInvalidLength
	}

	if len(u.Username) < UsernameMinimumLength || len(u.Username) > UsernameMaximumLength {
		return ErrUsernameInvalidLength
	}

	// Validate whether the username is a valid complete email address,
	// or a local part of a valid email address
	if _, err := mail.ParseAddress(u.Username); err != nil {
		if _, err := mail.ParseAddress(u.Username + "@localhost"); err != nil {
			return ErrUsernameInvalid
		}
	}

	return nil
}

func (u *User) SetPassword(password string) error {
	if len(password) < PasswordMinimumLength || len(password) > PasswordMaximumLength {
		return ErrPasswordInvalidLength
	}

	u.GenerateSalt()

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.AddSalt(password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(cryptedPassword)

	return nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) GenerateAPIKey(force bool) {
	if !force && u.APIKey != "" {
		return
	}

	u.APIKey = rand.String(32, rand.GetAlphaNumericPool())
}

func (u *User) GenerateSalt() {
	if u.Salt != "" {
		return
	}

	u.Salt = rand.String(8, rand.GetAlphaNumericPool())
}

func (u *User) Save(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(u).Error
}

func (u *User) GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w *Workout

	db = db.Preload("Data").Preload("Data.Details").Preload("GPX").Preload("Equipment")

	if err := db.Where(&Workout{UserID: u.ID}).First(&w, id).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) MarkWorkoutsDirty(db *gorm.DB) error {
	return db.Model(&Workout{}).Where(&Workout{UserID: u.ID}).Update("dirty", true).Error
}

func (u *User) GetWorkouts(db *gorm.DB) ([]*Workout, error) {
	var w []*Workout

	if err := db.Preload("Data").Where(&Workout{UserID: u.ID}).Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) AddWorkout(db *gorm.DB, workoutType WorkoutType, notes string, filename string, content []byte) (*Workout, error) {
	if u == nil {
		return nil, ErrNoUser
	}

	w, err := NewWorkout(u, workoutType, notes, filename, content)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	if err := w.Create(db); err != nil {
		return nil, err
	}

	var equipment []*Equipment

	for i, e := range u.Equipment {
		if e.ValidFor(&w.Type) {
			equipment = append(equipment, &u.Equipment[i])
		}
	}

	if err := db.Model(&w).Association("Equipment").Replace(equipment); err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) GetAllEquipment(db *gorm.DB) ([]*Equipment, error) {
	var w []*Equipment

	if err := db.Preload("Workouts").Where(&Equipment{UserID: u.ID}).Order("name DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) GetEquipment(db *gorm.DB, id int) (*Equipment, error) {
	var w *Equipment

	if err := db.Where(&Equipment{UserID: u.ID}).First(&w, id).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func AddRouteSegment(db *gorm.DB, notes string, filename string, content []byte) (*RouteSegment, error) {
	rs, err := NewRouteSegment(notes, filename, content)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	if err := rs.Create(db); err != nil {
		return nil, err
	}

	return rs, nil
}
