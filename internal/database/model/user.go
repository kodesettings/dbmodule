//go:build js && wasm
package database_model

import (
	bcrypt "golang.org/x/crypto/bcrypt"
	. "github.com/kodesettings/dbmodule/v2/internal/core"
)

type IRole struct {
	MEMBER string;
	ADMIN string;
	MODERATOR string;
}

const (
	ROLE_MEMBER = "member"
	ROLE_ADMIN = "admin"
	ROLE_MODERATOR = "superadmin"
)

type User struct {
	Id                      uint64  `json:"id"`
	Fullname                string  `json:"fullname"`
	Email                   string  `json:"email"`
	Password                string  `json:"password"`
	IsEmailVerified         bool    `json:"isEmailVerified"`
	Roles                   []Role  `json:"roles"`
	CreatedAt               uint64  `json:"createdAt"`
	UpdatedAt               uint64  `json:"updatedAt"`
}

type Role struct {
	Id                      uint64  `json:"id"`
	Role                    string  `json:"role"`
}

// Hashing Password Before Saving the User
func SavePassword(user User) User {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c := ApiError{/*empty struct, it is assigned in function call*/}
		c.InternalError(err.Error())
	}
	user.Password = string(hash)
	return user
}
