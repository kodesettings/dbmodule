//go:build js && wasm
package auth_middlewares

import (
	"net/http"
	jwt "github.com/golang-jwt/jwt/v5"
	. "github.com/kodesettings/dbmodule/v2/internal/core"
	database_repository "github.com/kodesettings/dbmodule/v2/internal/database/repository"
)

type auth struct {
	authentication _authentication `json:"auth"`
}

type _authentication struct {
	device_identifier string `json:"deviceIdentifier"`
}

func (a *auth) __decodeJWTToken(token_string string) string {
	_, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	});

	return err.Error()
}

func (a *auth) AuthenticationValidator(w http.ResponseWriter, req *http.Request) {
	deviceIdentifier := req.Header.Get("deviceIdentifier")
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		c := ApiError{/*empty struct, it is assigned in function call*/}
		c.BadRequestError("no authorization token provided")
		return
	}

	var token string = authHeader[:7]; // remove Bearer prefix from token
	var decodedToken string = a.__decodeJWTToken(token);

	if decodedToken == "" { c := ApiError{}; c.BadRequestError("invalid tokn"); return; }

	handler := database_repository.RefreshTokenRepo{}
	_, isUserLoggedIn := handler.FindByDeviceIdentifier(deviceIdentifier);

	if !isUserLoggedIn { c:= ApiError{}; c.InternalError("login again"); return; }
/*
    if (decodedToken.usage === "auth-access") {
      const finalDecodedToken = jwt.verify(
        token,
        tokenInfo.jwtSecret
      ) as JwtPayload;

      if (!finalDecodedToken) {
        throw new BadRequestError("send refresh token");
      }

      return new SuccessResponse("successful authentication").send();
    } else {
      throw new BadRequestError("token is not valid");
    }
*/
}
