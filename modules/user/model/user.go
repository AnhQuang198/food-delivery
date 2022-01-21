package model

import (
	"errors"
	"food-delivery/commons"
	"food-delivery/components/tokenprovider"
)

const EntityName = "user"

type User struct {
	commons.SQLModel `json:",inline"`
	Email            string         `json:"email" gorm:"column:email"`
	Password         string         `json:"password" gorm:"column:password"`
	LastName         string         `json:"last_name" gorm:"column:last_name"`
	FirstName        string         `json:"first_name" gorm:"column:first_name"`
	Phone            string         `json:"phone" gorm:"column:phone"`
	Role             string         `json:"role" gorm:"column:role"`
	Avatar           *commons.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	Email     string         `json:"email" gorm:"column:email"`
	Password  string         `json:"password" gorm:"column:password"`
	LastName  string         `json:"last_name" gorm:"column:last_name"`
	FirstName string         `json:"first_name" gorm:"column:first_name"`
	Phone     string         `json:"phone" gorm:"column:phone"`
	Role      string         `json:"-" gorm:"column:role"`
	Avatar    *commons.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUsernameOrPasswordInvalid = commons.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = commons.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
