//go:build js && wasm
package auth_api_key

import (
	"net/http"
	"io"
	"encoding/json"
	. "github.com/kodesettings/dbmodule/v2/internal/core"
	database_repository "github.com/kodesettings/dbmodule/v2/internal/database/repository"
)

type apikey struct {
	api_key        _apiKey         `json:"apiKey"`
	authentication _authentication `json:"auth"`
}

type apikey_controller struct {
	api_error      ApiError
	api_response   ApiResponse
}

type _apiKey struct {
	x_api_key string `json:"x-api-key"`
}

type _authentication struct {
	device_identifier string `json:"deviceIdentifier"`
}

func (a *apikey_controller) SuperAdminApiKeyCheck(w http.ResponseWriter, req *http.Request) {
	apiKey := apikey{}
	b, err := io.ReadAll(req.Body)
	if err != nil {
		a.api_error.InternalError(err.Error())
		return
	}

	json.Unmarshal([]byte(b), &apiKey)
	// TODO: get user from db and check its permission
	permission := false

	if !permission {
		r := ApiResponse{}
		r.ForbiddenResponse("Permission Denied");
	}
}

func (a *apikey_controller) ApiKeyCheck(w http.ResponseWriter, req *http.Request) {
	key := req.Header.Get("x-api-key")
	if key == "" {
		a.api_error.ForbiddenError("api key was not found")
		return
	}

	handler := database_repository.ApiKeyRepo{}
	_, found := handler.FindByKey(key);
	if !found {
		a.api_error.ForbiddenError("no api key was found")
		return
	}

	// TODO: get user from db and check its permission
	permission := false

	if !permission {
		a.api_response.ForbiddenResponse("Permission Denied");
	}

	// const exists = apiKey?.permissions.includes(Permission.GENERAL, 0);
	// if (!exists) throw new ForbiddenError("Permission Denied");
}
