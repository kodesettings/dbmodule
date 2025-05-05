//go:build js && wasm
package database_repository

import (
	"database/sql"
	"time"

	database_model "github.com/kodesettings/dbmodule/v2/internal/database/model"
)

type refresh_token_handler struct {
	db *sql.DB
}

func (h *refresh_token_handler) FindAll() []database_model.RefreshToken {
	rows, err := h.db.Query("SELECT * FROM refresh_token ORDER BY created_at DESC;")
	if err != nil {
		return []database_model.RefreshToken{}
	}

	refresh_tokens := []database_model.RefreshToken{}
	for rows.Next() {
		var a database_model.RefreshToken
		err = rows.Scan(&a.Id, &a.User, &a.Token, &a.DeviceIdentifier,
						&a.CreatedAt, &a.ExpiresAt)

		if err != nil {
			return []database_model.RefreshToken{}
		}
		refresh_tokens = append(refresh_tokens, a)
	}

	return refresh_tokens
}

func (h *refresh_token_handler) Create(refresh_token database_model.RefreshToken) bool {
	now := time.Now().Unix()
	refresh_token.CreatedAt = uint64(now);
	refresh_token.ExpiresAt = uint64(now); // TODO: add validity period

	_, err := h.db.Exec("INSERT INTO refresh_token (createdAt, expiresAt) VALUES (?, ?, ?);",
						refresh_token.CreatedAt, refresh_token.ExpiresAt)

	if err != nil {
		return false;
	}

	return true;
}


func (h *refresh_token_handler) Update(refresh_token database_model.RefreshToken) bool {
	_, err := h.db.Exec("UPDATE refresh_token SET (user = ?, token = ?, " +
						"deviceIdentifier = ?) WHERE Id = ?;",
						refresh_token.User, refresh_token.Token,
						refresh_token.DeviceIdentifier, refresh_token.Id)

	if err != nil {
		return false;
	}

	return true;
}

func (h *refresh_token_handler) Remove(refresh_token database_model.RefreshToken) bool {
	_, err := h.db.Exec("DELETE FROM refresh_token WHERE Id = ?;", refresh_token.Id)
	if err != nil {
		return false;
	}

	return true;
}

func (h *refresh_token_handler) FindById(id uint64) database_model.RefreshToken {
	row := h.db.QueryRow("SELECT * FROM refresh_token WHERE Id = ?;", id)

	refresh_token := database_model.RefreshToken{}
	if err := row.Scan(&refresh_token.Id, &refresh_token.User, &refresh_token.Token,
					&refresh_token.DeviceIdentifier, &refresh_token.CreatedAt,
					&refresh_token.ExpiresAt); err != nil {
		return database_model.RefreshToken{}
	}

	return refresh_token;
}

func (h *refresh_token_handler) FindByDeviceIdentifier(deviceIdentifier string) database_model.RefreshToken {
	row := h.db.QueryRow("SELECT * FROM refresh_token WHERE deviceIdentifier = ?;", deviceIdentifier)

	refresh_token := database_model.RefreshToken{}
	if err := row.Scan(&refresh_token.Id, &refresh_token.User, &refresh_token.Token,
					&refresh_token.DeviceIdentifier, &refresh_token.CreatedAt,
					&refresh_token.ExpiresAt); err != nil {
		return database_model.RefreshToken{}
	}

	return refresh_token;
}
