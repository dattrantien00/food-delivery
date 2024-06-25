package usermodel

import (
	"errors"
	"food-delivery/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"-" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask() {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Salt      string        `json:"salt" gorm:"column:salt;"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"-" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return "users"
}
func (u *UserCreate) Mask(bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return "users"
}

var (
	ErrEmailExist = common.NewErrorResponse(
		errors.New("Email existed"),
		"Email existed",
		"Email existed",
		"ErrEmailExisted",
	)

	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("Username or password invalid"),
		"Username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)
