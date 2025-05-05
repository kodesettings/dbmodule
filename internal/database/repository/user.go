//go:build js && wasm
package database_repository

import (
	"database/sql"
	"time"

	database_model "github.com/kodesettings/dbmodule/v2/internal/database/model"
)

type UserHandler struct {
	db *sql.DB
}

func (h *UserHandler) FindAll() []database_model.User {
	rows, err := h.db.Query("SELECT * FROM user ORDER BY created_at DESC;")
	if err != nil {
		return []database_model.User{}
	}

	users := []database_model.User{}
	for rows.Next() {
		var a database_model.User
		err = rows.Scan(&a.Id, &a.Fullname, &a.Email, &a.Password, &a.IsEmailVerified,
						&a.Roles, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return []database_model.User{}
		}
		users = append(users, a)
	}

	return users
}

func (h *UserHandler) Create(user database_model.User) bool {
	now := time.Now().Unix()
	user.CreatedAt = uint64(now);
	user.UpdatedAt = uint64(now);

	_, err := h.db.Exec("INSERT INTO user (createdAt, updatedAt) VALUES (?, ?, ?);",
						user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return false;
	}

	return true;
}


func (h *UserHandler) Update(user database_model.User) bool {
	now := time.Now().Unix()
	user.UpdatedAt = uint64(now);

	_, err := h.db.Exec("UPDATE user SET (fullname = ?, email = ?, password = ?, " +
						"isEmailVerified = ?, roles = ?, createdAt = ?, " +
						"updatedAt = ?) WHERE Id = ?;",
						user.Fullname, user.Email, user.Password, user.IsEmailVerified,
						user.Roles, user.CreatedAt, user.UpdatedAt, user.Id)
	if err != nil {
		return false;
	}

	return true;
}

func (h *UserHandler) Remove(user database_model.User) bool {
	_, err := h.db.Exec("DELETE FROM user WHERE Id = ?;", user.Id)
	if err != nil {
		return false;
	}

	return true;
}

func (h *UserHandler) FindById(id uint64) database_model.User {
	row := h.db.QueryRow("SELECT * FROM user WHERE Id = ?;", id)

	user := database_model.User{}
	if err := row.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password,
					&user.IsEmailVerified, &user.Roles, &user.CreatedAt,
					&user.UpdatedAt); err != nil {
		return database_model.User{}
	}

	return user;
}

func (h *UserHandler) FindByEmail(email string) database_model.User {
	row := h.db.QueryRow("SELECT * FROM user WHERE Email = ?;", email)

	user := database_model.User{}
	if err := row.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password,
					&user.IsEmailVerified, &user.Roles, &user.CreatedAt,
					&user.UpdatedAt); err != nil {
		return database_model.User{}
	}

	return user;
}
