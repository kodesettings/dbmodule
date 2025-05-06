//go:build js && wasm
package controllers

import (
	"net/http"
	"io"
	"encoding/json"
	jwt "github.com/golang-jwt/jwt/v5"
	bcrypt "golang.org/x/crypto/bcrypt"
	"fmt"
	"time"
	"strconv"
	. "github.com/kodesettings/dbmodule/v2/internal/core"
	. "github.com/kodesettings/dbmodule/v2/internal/config"
	helpers "github.com/kodesettings/dbmodule/v2/internal/helpers"
	database_model "github.com/kodesettings/dbmodule/v2/internal/database/model"
	database_repository "github.com/kodesettings/dbmodule/v2/internal/database/repository"
)

type user_identifiers struct {
	username         string `json:"username"`
	email            string `json:"email"`
	password         string `json:"password"`
	newPassword      string `json:"newPassword"`
	token            string `json:"token"`
	deviceIdentifier string `json:"deviceIdentifier"`
}

type jwt_claims struct {
    issuer    string `json:"iss"`
    subject   string `json:"sub"`
    expiresAt int64  `json:"exp"`
    issuedAt  int64  `json:"iat"`
    jwt.RegisteredClaims
}

type controller struct {
	api_error ApiError;
	api_response ApiResponse;
	handler database_repository.RefreshTokenRepo
	user_handler database_repository.UserRepo
}

func (c *controller) __parse_json_model(req *http.Request, user *database_model.User) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		c.api_error.InternalError(err.Error())
		return
	}

	json.Unmarshal([]byte(b), &user)
}

func (c *controller) __parse_json_id(req *http.Request, id *user_identifiers) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		c.api_error.InternalError(err.Error())
		return
	}

	json.Unmarshal([]byte(b), &id)
}

func (c *controller) __sign_jwt_token(username string, usage string) string {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"usage": usage,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	});

	tokenString, err := refreshToken.SignedString(/*secret jwt key*/TokenInfo[JwtSecret]);
	if err != nil {
		c.api_error.BadRequestError("token error: " + err.Error());
		return ""
	}

	return tokenString
}

func (c *controller) __decodeJWTToken(token_string string) string {
	_, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	});

	return err.Error()
}

func (c *controller) __verifyJWTToken(token_string string) bool {
    _, err := jwt.ParseWithClaims(token_string, &jwt_claims{}, func(token *jwt.Token) (interface{}, error) {
		// TODO: get usage information here
        return []byte(TokenInfo[JwtSecret]), nil // secret key inserted here
    })

	return err == nil
}


// POST - /auth/login - Login Handler
func (c *controller) Login(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("user email not found"); return; }

	isPasswordCorrect := bcrypt.CompareHashAndPassword([]byte(__id.password), []byte(user.Password))
	if isPasswordCorrect != nil {
		c.api_error.BadRequestError("wrong password");
	}

	accessToken := c.__sign_jwt_token(__id.username, "auth-access");

	isRefreshTokenAlreadyCreated, found := c.handler.FindByDeviceIdentifier(__id.deviceIdentifier);
	if found { c.handler.Remove(isRefreshTokenAlreadyCreated.Id); }
	validityPeriod, _ := strconv.Atoi(TokenInfo[RefreshTokenValidity]);
	refreshTokenExpirationTime := time.Now().Add(time.Duration(validityPeriod));

	refreshToken := c.__sign_jwt_token(__id.username, "auth-refresh");

	refreshTokenDBObj := database_model.RefreshToken{
		Id: user.Id,
		User: user.Fullname, // TODO: is this the right field?
		Token: refreshToken,
		DeviceIdentifier: __id.deviceIdentifier,
		CreatedAt: uint64(time.Now().Unix()),
		ExpiresAt: uint64(refreshTokenExpirationTime.Unix()),
	};

	isCreated := c.handler.Create(refreshTokenDBObj);
	if (!isCreated) {
		c.api_error.InternalError("error creating refreshToken db object");
	}

	type response struct {
		accessToken  string
		refreshToken string
		userId       uint64
	}

	obj, err := json.Marshal(response{accessToken, refreshToken, user.Id});
	if err != nil {
		c.api_response.SuccessResponse(obj);
	} else {
		c.api_error.BadRequestError(err.Error());
	}
}


// POST - /auth/register - Register Handler
func (c *controller) Register(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("user email not found"); return; }

	c.user_handler.Create(user);

	// Send Welcome Email
	helpers.MimeTextBuilder(user.Email, user.Fullname, "", "", "Welcome!", "Thank your for chosing us");
	// TODO: send email through cloudflare's smtp service

	c.api_response.SuccessResponse("registered");
}

// POST - /auth/access-token - Create new access token
func (c *controller) NewAccessToken(w http.ResponseWriter, req *http.Request) {
	deviceIdentifier := req.Header.Get("deviceIdentifier")
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		c := ApiError{/*empty struct, it is assigned in function call*/}
		c.BadRequestError("no authorization token provided")
		return
	}

	var token string = authHeader[:7]; // remove Bearer prefix from token
	var decodedToken string = c.__decodeJWTToken(token);

	if decodedToken == "" { c.api_error.BadRequestError("token is not defined"); return; }

	user, isUserLoggedIn := c.handler.FindByDeviceIdentifier(deviceIdentifier);
	if !isUserLoggedIn { c.api_error.ForbiddenError("login again"); }

	success := c.__verifyJWTToken(token);
	if !success {
		c.handler.Remove(user.Id);
		c.api_error.ForbiddenError("login again");
		return;
	}
/*
    if (decodedToken.usage !== "auth-refresh") {
      throw new BadRequestError("Invalid Token");
    }
*/
	accessToken := c.__sign_jwt_token(user.User, "auth-access");

	obj, err := json.Marshal(accessToken);
	if err != nil {
		c.api_response.SuccessResponse(obj);
	} else {
		c.api_error.BadRequestError(err.Error());
	}
}

// POST - /auth/verify-email-request - Emails a link that verifies the email
func (c *controller) VerifyEmailRequest(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("user email not found"); return; }

	isEmailAlreadyVerified := user.IsEmailVerified;
	if isEmailAlreadyVerified {
		c.api_error.BadRequestError("the email is already verified");
	}

	token := c.__sign_jwt_token(__id.username, "emailVerify");

	var verifyLink string = fmt.Sprintf("/auth/verify-email/%s", token);
	var body string = fmt.Sprintf("<a href=\"%s\">click here for verifying email/a>", verifyLink);

	helpers.MimeTextBuilder(user.Email, user.Fullname, "", "", "verify email", body);
	// TODO: send email through cloudflare's smtp service

	c.api_response.SuccessResponse("success");
}

// GET - /auth/verify-email/:token - Verifies user's email
func (c *controller) VerifyEmail(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	success := c.__verifyJWTToken(__id.token);
	if !success { c.api_error.BadRequestError("permission denied"); return; }
/*
    if (decodedToken.usage !== "emailVerify") {
      throw new BadRequestError("Invalid Token");
    }
*/
	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("user email not found"); return; }

	user.IsEmailVerified = true;
	c.user_handler.Update(user);

	c.api_response.SuccessResponse("succesfully verified email address");
}


// POST - /auth/forget-password - Forget Password Handler
func (c *controller) ForgetPassword(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("user email not found"); return; }

	token := c.__sign_jwt_token(__id.username, "forgetPassword");

	var resetLink string = fmt.Sprintf("/auth/reset-password/%s", token);
	var body string = fmt.Sprintf("<a href=\"%s\">click here for changing password/a>", resetLink);

	helpers.MimeTextBuilder(user.Email, user.Fullname, "", "", "forget password", body);
	// TODO: send email through cloudflare's smtp service

	c.api_response.SuccessResponse("succesfully verified email address");
}

// POST - /reset-password/:token - Reset Password Handler
func (c *controller) ResetForgettedPassword(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	success := c.__verifyJWTToken(__id.token);
	if !success { c.api_error.BadRequestError("permission denied"); return; }
/*
    if (decodedToken.usage !== "forgetPassword") {
      throw new BadRequestError("Invalid Token");
    }
*/
	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("no user found"); return; }

	user.Password = __id.password;
	c.user_handler.Update(user);

	c.api_response.SuccessResponse("reset password done");
}


// PUT - /change-password/:id - Change Password Handler
func (c *controller) ChangePassword(w http.ResponseWriter, req *http.Request) {
	var __id user_identifiers;
	c.__parse_json_id(req, &__id);

	user, found := c.user_handler.FindByEmail(__id.email);
	if !found { c.api_error.NotFoundError("no user found"); return; }

	isPasswordCorrect := bcrypt.CompareHashAndPassword([]byte(__id.password), []byte(user.Password))
	if isPasswordCorrect != nil {
		c.api_error.BadRequestError("wrong password");
	}

	user.Password = __id.newPassword;
	c.user_handler.Update(user);

	c.api_response.SuccessResponse("changing password done");
}
