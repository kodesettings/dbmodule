package core

import (
	. "github.com/kodesettings/dbmodule/v2/internal/config"
)

type ApiError struct {
	errorType ErrorState
	message   string
}

type ErrorState int32
type HeaderFields map[string]string

const (
   BAD_TOKEN ErrorState = iota
   TOKEN_EXPIRED
   UNAUTHORIZED
   INVALID_ACCESS_TOKEN
   INTERNAL_ERROR
   NOT_FOUND
   NO_ENTRY
   NO_DATA
   BAD_REQUEST
   FORBIDDEN
)

func (a *ApiResponse) handle(err ApiError, headers HeaderFields) func() {
	switch (err.errorType) {
		case BAD_TOKEN:
		case TOKEN_EXPIRED:
		case UNAUTHORIZED:
			return a.AuthFailureResponse(err.message).send(headers);
		case INVALID_ACCESS_TOKEN:
			return a.AccessTokenErrorResponse(err.message).send(headers);
		case INTERNAL_ERROR:
			return a.InternalErrorResponse(err.message).send(headers);
		case NOT_FOUND:
		case NO_ENTRY:
		case NO_DATA:
			return a.NotFoundResponse(err.message).send(headers);
		case BAD_REQUEST:
			return a.BadRequestResponse(err.message).send(headers);
		case FORBIDDEN:
			return a.ForbiddenResponse(err.message).send(headers);
		default: {
			message := err.message;
			// Do not send failure message in production as it may send sensitive data
			if (environment == "production") { message = "Something wrong happened."; }
			return a.InternalErrorResponse(message).send(headers);
		}
	}
}

func (a *ApiError) AuthFailureError(headers HeaderFields) func() {
	a.message = "Invalid Credentials"
	a.errorType = UNAUTHORIZED
	return handle(*a, headers)
}

func (a *ApiError) InternalError(headers HeaderFields) func() {
	a.message = "Internal error"
	a.errorType = INTERNAL_ERROR
	return handle(*a, headers)
}

func (a *ApiError) BadRequestError(headers HeaderFields) func() {
	a.message = "Bad Request"
	a.errorType = BAD_REQUEST
	return handle(*a, headers)
}

func (a *ApiError) NotFoundError(headers HeaderFields) func() {
	a.message = "Not Found"
	a.errorType = NOT_FOUND
	return handle(*a, headers)
}

func (a *ApiError) ForbiddenError(headers HeaderFields) func() {
	a.message = "Permission denied"
	a.errorType = FORBIDDEN
	return handle(*a, headers)
}

func (a *ApiError) NoEntryError(headers HeaderFields) func() {
	a.message = "Entry don't exists"
	a.errorType = NO_ENTRY
	return handle(*a, headers)
}

func (a *ApiError) BadTokenError(headers HeaderFields) func() {
	a.message = "Token is not valid"
	a.errorType = BAD_TOKEN
	return handle(*a, headers)
}

func (a *ApiError) TokenExpiredError(headers HeaderFields) func() {
	a.message = "Token is expired"
	a.errorType = TOKEN_EXPIRED
	return handle(*a, headers)
}

func (a *ApiError) NoDataError(headers HeaderFields) func() {
	a.message = "No data available"
	a.errorType = NO_DATA
	return handle(*a, headers)
}

func (a *ApiError) AccessTokenError(headers HeaderFields) func() {
	a.message = "Invalid access token"
	a.errorType = INVALID_ACCESS_TOKEN
	return handle(*a, headers)
}
