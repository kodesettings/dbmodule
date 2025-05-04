package core

import (
	"fmt"
	. "github.com/kodesettings/dbmodule/v2/internal/config"
)

type ApiResponse struct {
	statusCode string
	status     ResponseStatus
	message    string
	data       any
	headers    HeaderFields
}

type ResponseStatus int
type HeaderFields map[string]string

var StatusCode = map[ResponseStatus]string{
	STATUS_SUCCESS: "10000",
	STATUS_FAILURE: "10001",
	STATUS_RETRY: "10002",
	STATUS_INVALID_ACCESS_TOKEN: "10003",
}

const (
	STATUS_SUCCESS ResponseStatus = iota
	STATUS_FAILURE
	STATUS_RETRY
	STATUS_INVALID_ACCESS_TOKEN
)

const (
	SUCCESS ResponseStatus= 200
	BAD_REQUEST ResponseStatus = 400
	UNAUTHORIZED ResponseStatus = 401
	FORBIDDEN ResponseStatus = 403
	NOT_FOUND ResponseStatus = 404
	INTERNAL_ERROR ResponseStatus  = 500
)

/*
abstract class ApiResponse {
  constructor(protected statusCode: StatusCode, protected status: ResponseStatus, protected message: string) {}

  protected async prepare(response: any, headers: { [key: string]: string } = {}): Promise<Response> {
    let options = {status: this.status, statusText: this.message, headers: headers};
    return new Response(response, options);
  }

  public async send(response?: any, headers?: { [key: string]: string }): Promise<Response> {
    return this.prepare(response, headers);
  }
}
*/

func (a *ApiResponse) ApiResponseBuilder(apiResponse ApiResponse, headers HeaderFields) string {
	prepare := func(apiResponse ApiResponse) string {
		options := fmt.Sprint("{\"status\": \"%d\", \"statusText:\", \"%s\", \"headers:\" \"%s\"}",
			apiResponse.status, apiResponse.message, apiResponse.headers)
		return options
	}

	return prepare(*a)
}

func (a *ApiResponse) AuthFailureResponse(message string) func() {
	if (message == "") { a.message = "Authentication Failure" }
	a.statusCode = StatusCode[STATUS_FAILURE]
	a.status = UNAUTHORIZED
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) NotFoundResponse(message string) func() {
	if (message == "") { a.message = "Not Found" }
	a.statusCode = StatusCode[STATUS_FAILURE]
	a.status = NOT_FOUND
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) ForbiddenResponse(message string) func() {
	if (message == "") { a.message = "Forbidden" }
	a.statusCode = StatusCode[STATUS_FAILURE]
	a.status = FORBIDDEN
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) BadRequestResponse(message string) func() {
	if (message == "") { a.message = "Bad Parameters" }
	a.statusCode = StatusCode[STATUS_FAILURE]
	a.status = BAD_REQUEST
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) InternalErrorResponse(message string) func() {
	if (message == "") { a.message = "Internal Error" }
	a.statusCode = StatusCode[STATUS_FAILURE]
	a.status = INTERNAL_ERROR
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) SuccessResponse(data any) func() {
	a.message = "Success"
	a.data = data
	a.statusCode = StatusCode[STATUS_SUCCESS]
	a.status = SUCCESS
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) AccessTokenErrorResponse(data any) func() {
	a.message = "Access token invalid"

	var instruction string = "refresh_token"
	a.headers["instruction"] = instruction

	a.statusCode = StatusCode[STATUS_INVALID_ACCESS_TOKEN]
	a.status = UNAUTHORIZED
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}

func (a *ApiResponse) TokenRefreshResponse(accessToken string, refreshToken string) func() {
	a.message = "Access token invalid"
	a.statusCode = StatusCode[STATUS_SUCCESS]
	a.status = SUCCESS
	return func() { a.ApiResponseBuilder(*a, a.headers) }
}
