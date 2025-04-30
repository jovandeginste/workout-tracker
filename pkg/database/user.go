package database

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/cat-dealer/go-rand/v2"
	"github.com/invopop/ctxi18n"
	"github.com/invopop/ctxi18n/i18n"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

type UserSecrets struct {
	Password string `gorm:"type:varchar(128);not null"` // The user's password as bcrypt hash
	Salt     string `gorm:"type:varchar(16);not null"`  // The salt used to hash the user's password
	APIKey   string `gorm:"type:varchar(32)"`           // The user's API key
}

type UserData struct {
	Model
	LastVersion string `gorm:"last_version" json:"lastVersion"` // Which version of the app the user has last seen and acknowledged

	Username string `form:"username" gorm:"uniqueIndex;not null;type:varchar(32)" json:"username"` // The user's username
	Name     string `form:"name" gorm:"type:varchar(64);not null" json:"name"`                     // The user's name

	Active bool `form:"active" json:"active"` // Whether the user is active
	Admin  bool `form:"admin" json:"admin"`   // Whether the user is an admin
}

type User struct {
	db      *gorm.DB
	context context.Context

	UserData
	UserSecrets `swaggerignore:"true"`

	Workouts     []Workout     `gorm:"constraint:OnDelete:CASCADE" json:"-"` // The user's workouts
	Equipment    []Equipment   `gorm:"constraint:OnDelete:CASCADE" json:"-"` // The user's equipment
	Measurements []Measurement `gorm:"constraint:OnDelete:CASCADE" json:"-"` // The user's measurements

	Profile Profile `gorm:"constraint:OnDelete:CASCADE" json:"profile"` // The user's profile settings

	anonymous bool // Whether we have an actual user or not
}

func (u *User) GetContext() context.Context {
	return u.context
}

func (u *User) SetContext(ctx context.Context) {
	u.context = ctx
}

func (u *User) I18n(message string, vars ...any) string {
	return u.GetTranslator().T(message, vars...)
}

func (u *User) GetTranslator() *i18n.Locale {
	return ctxi18n.Locale(u.context)
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

func GetUsers(db *gorm.DB) ([]*User, error) {
	var u []*User

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

	if err := currentUserQuery(db).Where(&UserSecrets{APIKey: key}).First(&u).Error; err != nil {
		return nil, db.Error
	}

	u.SetDB(db)

	return &u, nil
}

func GetUserByID(db *gorm.DB, userID uint64) (*User, error) {
	var u User

	if err := currentUserQuery(db).First(&u, userID).Error; err != nil {
		return nil, db.Error
	}

	u.SetDB(db)

	return &u, nil
}

func GetUser(db *gorm.DB, username string) (*User, error) {
	var u User

	if err := currentUserQuery(db).Where(&UserData{Username: username}).First(&u).Error; err != nil {
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
	return db.Select(clause.Associations).Delete(u).Error
}

func (u *User) GetMeasurementForDate(date time.Time) (*Measurement, error) {
	var m *Measurement

	if err := u.db.Where(&Measurement{UserID: u.ID}).Where("date = ?", datatypes.Date(date.UTC())).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u.NewMeasurement(date), nil
		}

		return nil, err
	}

	return m, nil
}

func (u *User) GetLatestMeasurementForDate(date time.Time) (Measurement, error) {
	var m Measurement

	if err := u.db.Where(&Measurement{UserID: u.ID}).Where("date <= ?", datatypes.Date(date)).Order("date DESC").First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *u.NewMeasurement(date), nil
		}

		return *u.NewMeasurement(date), err
	}

	return m, nil
}

func (u *User) GetLatestMeasurements(c int) ([]*Measurement, error) {
	var m []*Measurement

	if err := u.db.Where(&Measurement{UserID: u.ID}).Order("date DESC").Limit(c).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (u *User) GetCurrentMeasurement() (*Measurement, error) {
	var m *Measurement

	d := time.Now().In(u.Timezone())

	if err := u.db.Where(&Measurement{UserID: u.ID}).Where("date = ?", datatypes.Date(d)).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u.NewMeasurement(d), nil
		}

		return nil, err
	}

	return m, nil
}

func (u *User) GetLatestMeasurement() (Measurement, error) {
	return u.GetLatestMeasurementForDate(time.Now())
}

func (u *User) GetWorkout(db *gorm.DB, id uint64) (*Workout, error) {
	var w *Workout

	db = db.Preload("Data").Preload("Data.Details").Preload("GPX").Preload("Equipment")

	if err := db.Where(&Workout{UserID: u.ID}).First(&w, id).Error; err != nil {
		return nil, err
	}

	w.User = u

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

	for _, wo := range w {
		wo.User = u
	}

	return w, nil
}

func (u *User) AddWorkout(db *gorm.DB, workoutType WorkoutType, notes string, filename string, content []byte) ([]*Workout, []error) {
	if u == nil {
		return nil, []error{ErrNoUser}
	}

	ws, err := NewWorkout(u, workoutType, notes, filename, content)
	if err != nil {
		return nil, []error{fmt.Errorf("%w: %s", ErrInvalidData, err)}
	}

	errs := []error{}

	for _, w := range ws {
		if err := w.Create(db); err != nil {
			errs = append(errs, err)
		}

		var equipment []*Equipment

		for i, e := range u.Equipment {
			if e.ValidFor(&w.Type) {
				equipment = append(equipment, &u.Equipment[i])
			}
		}

		if err := db.Model(&w).Association("Equipment").Replace(equipment); err != nil {
			errs = append(errs, err)
		}
	}

	return ws, errs
}

func (u *User) GetAllEquipment(db *gorm.DB) ([]*Equipment, error) {
	var w []*Equipment

	if err := db.Preload("Workouts").Where(&Equipment{UserID: u.ID}).Order("name DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	for _, e := range w {
		e.User = *u
	}

	return w, nil
}

func (u *User) GetEquipment(db *gorm.DB, id uint64) (*Equipment, error) {
	var w *Equipment

	if err := db.Where(&Equipment{UserID: u.ID}).First(&w, id).Error; err != nil {
		return nil, err
	}

	w.User = *u

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

func (u *User) HeightAt(d time.Time) float64 {
	w := u.measurementAt("height", d)
	if w == 0 {
		return 165
	}

	return w
}

func (u *User) WeightAt(d time.Time) float64 {
	w := u.measurementAt("weight", d)
	if w == 0 {
		return 70
	}

	return w
}

func (u *User) measurementAt(key string, d time.Time) float64 {
	var w float64

	q := u.db.
		Model(&Measurement{}).
		Where(&Measurement{UserID: u.ID}).
		Where("date <= ?", datatypes.Date(d)).
		Where("? > ?", key, 0).
		Order("date DESC").
		Pluck(key, &w)

	if err := q.Error; err != nil {
		return 0
	}

	return w
}
