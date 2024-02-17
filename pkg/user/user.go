package user

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUsernameInvalidLength = errors.New("username has invalid length")
	ErrUsernameInvalidChars  = errors.New("username contains invalid characters")

	UsernameValidRegex = regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
)

type User struct {
	gorm.Model
	Password string `form:"password" json:"password"`
	Username string `form:"username" gorm:"uniqueIndex;not null" json:"username"`
	Name     string `form:"name"     json:"name"`
	Active   bool   `form:"active"   json:"active"`
	Admin    bool   `form:"admin"    json:"admin"`

	Profile  Profile
	Workouts []Workout
}

func GetUser(db *gorm.DB, username string) (*User, error) {
	var u *User

	if err := db.Where(&User{Username: username}).First(&u).Error; err != nil {
		return nil, db.Error
	}

	return u, nil
}

type Profile struct {
	gorm.Model
	UserID int
	Theme  ThemePreference
}

type ThemePreference string

func (u *User) ValidLogin(password string) bool {
	if u == nil {
		return false
	}

	if !u.Active || u.Password == "" || u.Username == "" {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) IsValid() error {
	if len(u.Username) < 3 || len(u.Username) > 20 {
		return ErrUsernameInvalidLength
	}

	if !UsernameValidRegex.MatchString(u.Username) {
		return ErrUsernameInvalidChars
	}

	return nil
}

func (u *User) CryptPassword() error {
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(cryptedPassword)

	return nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) UpdateUser(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w *Workout

	if err := db.Where(&Workout{UserID: u.ID}).First(&w, id).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) GetWorkouts(db *gorm.DB) ([]Workout, error) {
	var w []Workout

	if err := db.Where(&Workout{UserID: u.ID}).Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) AddWorkout(db *gorm.DB, workoutType, notes string, content []byte) (*Workout, error) {
	w := NewWorkout(u, workoutType, notes, content)

	if err := w.Create(db); err != nil {
		return nil, err
	}

	return w, nil
}
