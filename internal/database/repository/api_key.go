//go:build js && wasm
package database_repository

import (
	"database/sql"
	"time"

	database_model "github.com/kodesettings/dbmodule/v2/internal/database/model"
)

type ApiKeyRepo struct {
	db *sql.DB
}

func (h *ApiKeyRepo) FindAll() []database_model.ApiKey {
	rows, err := h.db.Query("SELECT * FROM api_key ORDER BY created_at DESC;")
	if err != nil {
		return []database_model.ApiKey{}
	}

	api_keys := []database_model.ApiKey{}
	for rows.Next() {
		var a database_model.ApiKey
		err = rows.Scan(&a.Id, &a.Key, &a.Permissions, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return []database_model.ApiKey{}
		}
		api_keys = append(api_keys, a)
	}

	return api_keys
}

func (h *ApiKeyRepo) Create(api_key database_model.ApiKey) bool {
	database_model.CreateSuperAdminApiKey();

	now := time.Now().Unix()
	api_key.CreatedAt = uint64(now);
	api_key.UpdatedAt = uint64(now);

	_, err := h.db.Exec("INSERT INTO api_key (createdAt, updatedAt) VALUES (?, ?, ?);",
						api_key.CreatedAt, api_key.UpdatedAt)
	if err != nil {
		return false;
	}

	return true;
}


func (h *ApiKeyRepo) Update(api_key database_model.ApiKey) bool {
	database_model.CreateSuperAdminApiKey();

	now := time.Now().Unix()
	api_key.UpdatedAt = uint64(now);

	_, err := h.db.Exec("UPDATE api_key SET (key = ?, permission = ?, updatedAt = ?) WHERE Id = ?;",
						api_key.Key, api_key.Permissions, api_key.UpdatedAt, api_key.Id)
	if err != nil {
		return false;
	}

	return true;
}

func (h *ApiKeyRepo) Remove(api_key database_model.ApiKey) bool {
	database_model.CreateSuperAdminApiKey();

	_, err := h.db.Exec("DELETE FROM api_key WHERE Id = ?;", api_key.Id)
	if err != nil {
		return false;
	}

	return true;
}

func (h *ApiKeyRepo) FindById(id uint64) (database_model.ApiKey, bool) {
	database_model.CreateSuperAdminApiKey();

	row := h.db.QueryRow("SELECT * FROM api_key WHERE Id = ?;", id)

	api_key := database_model.ApiKey{}
	if err := row.Scan(&api_key.Id, &api_key.Key, &api_key.Permissions,
					&api_key.CreatedAt, &api_key.UpdatedAt); err != nil {
		return database_model.ApiKey{}, false;
	}

	return api_key, true;
}

func (h *ApiKeyRepo) FindByKey(key string) (database_model.ApiKey, bool) {
	database_model.CreateSuperAdminApiKey();

	row := h.db.QueryRow("SELECT * FROM api_key WHERE Key = ?;", key)

	api_key := database_model.ApiKey{}
	if err := row.Scan(&api_key.Id, &api_key.Key, &api_key.Permissions,
					&api_key.CreatedAt, &api_key.UpdatedAt); err != nil {
		return database_model.ApiKey{}, false;
	}

	return api_key, true;
}

func (h *ApiKeyRepo) FindByKeyButId(id uint64, key string) (database_model.ApiKey, bool) {
	database_model.CreateSuperAdminApiKey();

	row := h.db.QueryRow("SELECT * FROM api_key WHERE (Id = ?, Key = ?);", id, key)

	api_key := database_model.ApiKey{}
	if err := row.Scan(&api_key.Id, &api_key.Key, &api_key.Permissions,
				&api_key.CreatedAt, &api_key.UpdatedAt); err != nil {
		return database_model.ApiKey{}, false;
	}

	return api_key, true;
}
