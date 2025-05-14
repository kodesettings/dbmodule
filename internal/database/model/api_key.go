//go:build js && wasm
package database_model

import (
	"time"
	. "github.com/kodesettings/dbmodule/v2/internal/config"
)

type ApiKeyRepo interface {
	FindAll() []ApiKey
	Create(api_key ApiKey) bool
}

const (
	PERMISSION_GENERAL = "GENERAL"
	PERMISSION_SUPER_ADMIN = "SUPER_ADMIN"
)

const (
	STATUS_ACTIVE = 1
	STATUS_INACTIVE = 0
)

type ApiKey struct {
	Id           uint64 `json:"id"`
	Key          string `json:"key"`
	Permissions  string `json:"permissions"`
	Status       uint16 `json:"status"`
	CreatedAt    uint64 `json:"createdAt"`
	UpdatedAt    uint64 `json:"updatedAt"`
}

func CreateSuperAdminApiKey() bool {
	var handler ApiKeyRepo;

	// check if there are any documents in the collection
	var apiKeys []ApiKey = handler.FindAll();

	if len(apiKeys) != 0 {
		return true;
	}

	// if no documents exist, add a default document
	var apiKey = ApiKey{
		Id: 0,
		Key: SuperAdminApiKey,
		Permissions: PERMISSION_SUPER_ADMIN,
		Status: STATUS_ACTIVE,
		CreatedAt: uint64(time.Now().Unix()),
		UpdatedAt: uint64(time.Now().Unix()),
	};

	// save the default document
	return handler.Create(apiKey);
}
