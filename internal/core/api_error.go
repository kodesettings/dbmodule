//go:build js && wasm
package core

import (
	_ "github.com/kodesettings/dbmodule/v2/internal/config"
)

type ApiError struct {
	errorType     ErrorState
	message       string
	response     *ApiResponse
}

type ErrorState int32
type __errhandle func()

const (
	ERR_CODE_BAD_TOKEN ErrorState = iota
	ERR_CODE_TOKEN_EXPIRED
	ERR_CODE_UNAUTHORIZED
	ERR_CODE_INVALID_ACCESS_TOKEN
	ERR_CODE_INTERNAL_ERROR
	ERR_CODE_NOT_FOUND
	ERR_CODE_NO_ENTRY
	ERR_CODE_NO_DATA
	ERR_CODE_BAD_REQUEST
	ERR_CODE_FORBIDDEN
)

func (a *ApiError) handle(err ApiError, error_msg string) func() {
	switch (err.errorType) {
	case ERR_CODE_BAD_TOKEN:
	case ERR_CODE_TOKEN_EXPIRED:
	case ERR_CODE_UNAUTHORIZED:
		return a.response.AuthFailureResponse(a.message);
	case ERR_CODE_INVALID_ACCESS_TOKEN:
		return a.response.AccessTokenErrorResponse(a.message);
	case ERR_CODE_INTERNAL_ERROR:
		return a.response.InternalErrorResponse(a.message);
	case ERR_CODE_NOT_FOUND:
	case ERR_CODE_NO_ENTRY:
	case ERR_CODE_NO_DATA:
		return a.response.NotFoundResponse(a.message);
	case ERR_CODE_BAD_REQUEST:
		return a.response.BadRequestResponse(a.message);
	case ERR_CODE_FORBIDDEN:
		return a.response.ForbiddenResponse(a.message);
	default: {
		// Do not send failure message in production as it may send sensitive data
		// if (environment == "production") { a.message = "Something wrong happened."; }
		return a.response.InternalErrorResponse(a.message);
	}
	}
	return nil
}

func (a *ApiError) AuthFailureError(error_msg string) __errhandle {
	a.message = "Invalid Credentials"
	a.errorType = ERR_CODE_UNAUTHORIZED
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) InternalError(error_msg string) __errhandle {
	a.message = "Internal error"
	a.errorType = ERR_CODE_INTERNAL_ERROR
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) BadRequestError(error_msg string) __errhandle {
	a.message = "Bad Request"
	a.errorType = ERR_CODE_BAD_REQUEST
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) NotFoundError(error_msg string) __errhandle {
	a.message = "Not Found"
	a.errorType = ERR_CODE_NOT_FOUND
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) ForbiddenError(error_msg string) __errhandle {
	a.message = "Permission denied"
	a.errorType = ERR_CODE_FORBIDDEN
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) NoEntryError(error_msg string) __errhandle {
	a.message = "Entry don't exists"
	a.errorType = ERR_CODE_NO_ENTRY
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) BadTokenError(error_msg string) __errhandle {
	a.message = "Token is not valid"
	a.errorType = ERR_CODE_BAD_TOKEN
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) TokenExpiredError(error_msg string) __errhandle {
	a.message = "Token is expired"
	a.errorType = ERR_CODE_TOKEN_EXPIRED
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) NoDataError(error_msg string) __errhandle {
	a.message = "No data available"
	a.errorType = ERR_CODE_NO_DATA
	return func() { a.handle(*a, error_msg) }
}

func (a *ApiError) AccessTokenError(error_msg string) __errhandle {
	a.message = "Invalid access token"
	a.errorType = ERR_CODE_INVALID_ACCESS_TOKEN
	return func() { a.handle(*a, error_msg) }
}
