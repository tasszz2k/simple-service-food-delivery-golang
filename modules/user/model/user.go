package usermodel

import (
	"errors"
	"simple-service-food-delivery-golang/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `sql:", inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	Salt            string `json:"-" gorm:"column:salt"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Role            string `json:"role" gorm:"column:role"`
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetSalt() string {
	return u.Salt
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetPhone() string {
	return u.Phone
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `sql:", inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	Role            string `json:"-" gorm:"column:role"`
	Salt            string `json:"-" gorm:"column:salt"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordIncorrect = common.NewCustomError(
		errors.New("Username or password incorrect"),
		"Username or password incorrect",
		"ERR_USERNAME_OR_PASSWORD_INCORRECT",
	)
)
