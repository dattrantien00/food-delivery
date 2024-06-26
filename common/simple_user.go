package common

import (
)

const EntityName = "User"

type SimpleUser struct {
	SQLModel
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"-" gorm:"column:role;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

// func (u *User) GetUserId() int {
// 	return u.Id
// }

// func (u *User) GetEmail() string {
// 	return u.Email
// }

// func (u *User) GetRole() string {
// 	return u.Role
// }

func (u *SimpleUser) Mask() {
	u.GenUID(DbTypeUser)
}
