package database

import (
	"errors"
	"fmt"
	"net/mail"

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
	ErrUsernameInvalid       = errors.New("username is not a mail address")
)

type User struct {
	gorm.Model
	Password string `form:"-"        gorm:"type:varchar(128);not null"`
	Salt     string `form:"-"        gorm:"type:varchar(16);not null"`
	Username string `form:"username" gorm:"uniqueIndex;not null;type:varchar(32)"`
	Name     string `form:"name"     gorm:"type:varchar(64);not null"`
	Active   bool   `form:"active"`
	Admin    bool   `form:"admin"`

	Profile  Profile
	Workouts []Workout
}

func (u *User) BeforeSave(_ *gorm.DB) error {
	if err := u.GenerateSalt(); err != nil {
		return err
	}

	return u.IsValid()
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, db.Error
	}

	return u, nil
}

func GetUserByID(db *gorm.DB, userID int) (*User, error) {
	var u User

	if err := db.First(&u, userID).Error; err != nil {
		return nil, db.Error
	}

	return &u, nil
}

func GetUser(db *gorm.DB, username string) (*User, error) {
	var u User

	if err := db.Preload("Profile").Where(&User{Username: username}).First(&u).Error; err != nil {
		return nil, db.Error
	}

	return &u, nil
}

type Profile struct {
	gorm.Model
	UserID   int
	Theme    string `form:"theme"`
	Language string `form:"language"`
}

func (p *Profile) Save(db *gorm.DB) error {
	return db.Save(p).Error
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

	if _, err := mail.ParseAddress(u.Username); err != nil {
		return fmt.Errorf("%w: %s", ErrUsernameInvalid, err)
	}

	return nil
}

func (u *User) SetPassword(password string) error {
	if len(password) < PasswordMinimumLength || len(password) > PasswordMaximumLength {
		return ErrPasswordInvalidLength
	}

	if err := u.GenerateSalt(); err != nil {
		return err
	}

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

func (u *User) GenerateSalt() error {
	if u.Salt != "" {
		return nil
	}

	u.Salt = rand.String(8, rand.GetAlphaNumericPool())

	return nil
}

func (u *User) Save(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(u).Error
}

func (u *User) GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w *Workout

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

	if err := db.Where(&Workout{UserID: u.ID}).Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) AddWorkout(db *gorm.DB, workoutType WorkoutType, notes string, content []byte) (*Workout, error) {
	w := NewWorkout(u, workoutType, notes, content)

	if err := w.Create(db); err != nil {
		return nil, err
	}

	return w, nil
}
