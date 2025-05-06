//go:build js && wasm
package auth_middlewares

import (
	"net/http"
	"encoding/json"
	"io"
	"strings"
	. "github.com/kodesettings/dbmodule/v2/internal/core"
	database_model "github.com/kodesettings/dbmodule/v2/internal/database/model"
	database_repository "github.com/kodesettings/dbmodule/v2/internal/database/repository"
)

type auth struct {
	userId string `json:"userId"`
}

func (a *auth) __getUserId(req *http.Request) uint64 {
	user := database_model.User{}
	b, err := io.ReadAll(req.Body)
	if err != nil {
		c := ApiError{/*empty struct, it is assigned in function call*/}
		c.InternalError(err.Error())
		return 0
	}

	json.Unmarshal([]byte(b), &user)
	return user.Id
}

func (a *auth) AuthorizationValidator(w http.ResponseWriter, req *http.Request) {
	var userId uint64 = a.__getUserId(req);

	handler := database_repository.UserRepo{}
	user, found := handler.FindById(userId);
	if !found {
		c := ApiError{/*empty struct, it is assigned in function call*/}
		c.ForbiddenError("cannot find user id")
		return
	}

	userRoleNames := []string{database_model.ROLE_MEMBER,
		database_model.ROLE_ADMIN, database_model.ROLE_MODERATOR }

	var hasRequiredRoles bool

	for _, s := range user.Roles {
		var i int = 0
role_check:
		found := strings.Contains(s.Role, userRoleNames[i])
		i++
		if !found { goto role_check; }
		hasRequiredRoles = hasRequiredRoles && found;
	}

	if hasRequiredRoles {
		r := ApiResponse{};
		r.SuccessResponse("authorized");
	} else {
		c := ApiError{/*empty struct, it is assigned in function call*/}
		c.BadRequestError("permission error - don't have enough roles")
	}
}
