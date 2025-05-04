package database_model

import (
	. "golang.org/x/crypto/bcrypt"
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

type RefreshToken struct {
	_id                     uint64  `json:"id"`
	fullname                string  `json:"fullname"`
	email                   string  `json:"email"`
	password                string  `json:"password"`
	isEmailVerified         bool    `json:"isEmailVerified"`
	roles                   []Role  `json:"roles"`
	createdAt               uint64  `json:"createdAt"`
	updatedAt               uint64  `json:"updatedAt"`
}

type Role struct {
	_id                     uint64  `json:"id"`
	role                    string  `json:"role"`
}

// Hashing Password Before Saving the User
func savePassword(user User) User {
	hash, err := bcrypt.GenerateFromPassword(user.password, bcrypt.DefaultCost)
	if err != nil {
		return InternalError(err.message)
	}
	user.password = hash
	return user
}
