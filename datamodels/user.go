package datamodels

import (
	"time"

	//"golang.org/x/crypto/bcrypt"
)

// User is our User example model.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/user.go"
// which could wrap by embedding the datamodels.User or
// define completely new fields instead but for the shake
// of the example, we will use this datamodel
// as the only one User model in our application.
type User struct {
	ID              int64     `json:"userId" form:"id"`
	Username        string    `json:"username" form:"username"`
	Password        string    `json:"password,omitempty" form:"password"`
	HashedPassword  []byte    `json:"-" form:"-"`
	VerificationCode string    `json:"verificationCode,omitempty" form:"verificationCode"`
	CreatedAt       time.Time `json:"-" form:"-"`
	Token           string    `json:"authToken" form:"authToken"`
}

// IsValid can do some very very simple "low-level" data validations.
func (u User) IsValid() bool {
	return u.ID > 0
}

// GeneratePassword will generate a hashed password for us based on the
// user's input.
func GeneratePassword(userPassword string) ([]byte, error) {
	//return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return []byte(userPassword), nil
}

// ValidatePassword will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	//if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
	//	return false, err
	//}
	return true, nil
}

type Credit struct {
	CreditCode string `json:"creditCode"`
}

type Remain struct {
	RemainCount int `json:"remainCount"`
}
type SimpleUser struct {
	UserId              int64     `json:"userId" form:"userId"`
}

type UserInfo struct{
	UserId int64 `json:"userId"`
	Username string `json:"username"`
	Level int `json:"level"`
	Credit float32 `json:"credit"`
	Tags []string `json:"tags"`
	Project int `json:"project"`
}

type Invite struct {
	InviteCode string `json:"inviteCode,omitempty"`
	IsValid int `json:"isValid,omitempty"`
}