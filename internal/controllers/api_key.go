//go:build js && wasm
package controllers

import (
	"net/http"
	"io"
	"encoding/json"
	. "github.com/kodesettings/dbmodule/v2/internal/core"
	_ "github.com/kodesettings/dbmodule/v2/internal/helpers"
	database_model "github.com/kodesettings/dbmodule/v2/internal/database/model"
	database_repository "github.com/kodesettings/dbmodule/v2/internal/database/repository"
)

type ak_identifiers struct {
	id     uint64 `json:"id"`
	key    string `json:"key"`
	status uint16 `json:"status"`
}

type ak_controller struct {
	api_error ApiError;
	api_response ApiResponse;
	handler database_repository.ApiKeyRepo
}

func (c *ak_controller) __parse_json_model(req *http.Request, apiKey *database_model.ApiKey) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		c.api_error.InternalError(err.Error())
		return
	}

	json.Unmarshal([]byte(b), &apiKey)
}

func (c *ak_controller) __parse_json_id(req *http.Request, id *ak_identifiers) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		c.api_error.InternalError(err.Error())
		return
	}

	json.Unmarshal([]byte(b), &id)
}

// GET - /apikey/find-all
func (c *ak_controller) FindAll(w http.ResponseWriter, req *http.Request) {
	var apiKeys []database_model.ApiKey = c.handler.FindAll();
	obj, err := json.Marshal(apiKeys);

	if err != nil {
		c.api_error.BadRequestError(err.Error());
	} else {
		c.api_response.SuccessResponse(obj);
	}
}

// POST - /apikey/create
func (c *ak_controller) Create(w http.ResponseWriter, req *http.Request) {
	var apiKey database_model.ApiKey
	c.__parse_json_model(req, &apiKey)

	_, found := c.handler.FindByKey(apiKey.Key);
	if !found { c.api_error.BadRequestError("key is already taken"); return; }
/*
    const permissions = apiKey.permissions;
    if (!isArrayOfStrings(permissions)) apiKey.permissions = ""; // TODO: fix this
*/
	var result bool = c.handler.Create(apiKey);

	if result {
		c.api_response.SuccessResponse("successfully created apiKey");
	} else {
		c.api_error.InternalError("database problem");
	}
}

// PUT - /apikey/update
func (c *ak_controller) Update(w http.ResponseWriter, req *http.Request) {
	var __id ak_identifiers;
	c.__parse_json_id(req, &__id);

	apiKey, found := c.handler.FindById(__id.id);
	if !found { c.api_error.NotFoundError("api key not found"); return; }

	apiKey, found = c.handler.FindByKeyButId(__id.id, __id.key);
	if !found { c.api_error.BadRequestError("key already taken"); return; }
/*
    if (ctx.req.query.permissions && isArrayOfStrings(ctx.req.query.permissions)) {
      // TODO: fix this
      // apiKey.permissions = addPermissions(ctx.req.query.permissions);
    }
*/
	if __id.status == 1 {
		apiKey.Status = database_model.STATUS_ACTIVE
	} else {
		apiKey.Status = database_model.STATUS_INACTIVE
	}

	var result bool = c.handler.Create(apiKey);

	if result {
		c.api_response.SuccessResponse("successfully updated apiKey");
	} else {
		c.api_error.InternalError("database problem");
	}
}

// DELETE - /apikey/remove
func (c *ak_controller) Remove(w http.ResponseWriter, req *http.Request) {
	var __id ak_identifiers;
	c.__parse_json_id(req, &__id);

	_, found := c.handler.FindById(__id.id);
	if !found { c.api_error.BadRequestError("api key not found"); return; }

	var result bool = c.handler.Remove(__id.id);

	if result {
		c.api_response.SuccessResponse("successfully deleted apiKey");
	} else {
		c.api_error.InternalError("database problem");
	}
}

// GET - /apikey/find-by-id
func (c *ak_controller) FindById(w http.ResponseWriter, req *http.Request) {
	var __id ak_identifiers;
	c.__parse_json_id(req, &__id);

	apiKey, found := c.handler.FindById(__id.id);
	if !found { c.api_error.BadRequestError("api key not found"); return; }

	c.api_response.SuccessResponse(apiKey);
}

// GET - /apikey/find-by-id
func (c *ak_controller) FindByKey(w http.ResponseWriter, req *http.Request) {
	var __id ak_identifiers;
	c.__parse_json_id(req, &__id);

	apiKey, found := c.handler.FindByKey(__id.key);
	if !found { c.api_error.BadRequestError("api key not found"); return; }

	if found {
		c.api_response.SuccessResponse(apiKey);
	} else {
		c.api_error.InternalError("database problem");
	}
}
