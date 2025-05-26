#ifndef CORE_API_ERROR_H
#define CORE_API_ERROR_H

#include "../config/env_configs.hpp"
#include "api_response.hpp"

#define BAD_TOKEN "BadTokenError"
#define TOKEN_EXPIRED "TokenExpiredError"
#define UNAUTHORIZED "AuthFailureError"
#define ACCESS_TOKEN "AccessTokenError"
#define INTERNAL_ERROR "InternalError"
#define NOT_FOUND "NotFoundError"
#define NO_ENTRY "NoEntryError"
#define NO_DATA "NoDataError"
#define BAD_REQUEST "BadRequestError"
#define FORBIDDEN "ForbiddenError"

typedef const std::string ERRTYPE;

class ApiError: public ApiResponse {
public:
  ApiError() {/* default constructor */}
  ApiError(ERRTYPE type, std::string message = "error") {}

  static void handle(const std::string &err) {
    if (err == UNAUTHORIZED) { AuthFailureResponse(err).prepare(); return; }
    if (err == ACCESS_TOKEN) { AccessTokenErrorResponse(err).prepare(); return; }
    if (err == INTERNAL_ERROR) { InternalErrorResponse(err).prepare(); return; }
    if (err == NO_DATA) { NotFoundResponse(err).prepare(); return; }
    if (err == BAD_REQUEST) { BadRequestResponse(err).prepare(); return; }
    if (err == FORBIDDEN) { ForbiddenResponse(err).prepare(); return; }
default_error:
    // Do not send failure message in production as it may send sensitive data
    // if (environment === "production") message = "Something wrong happened.";
    InternalErrorResponse(std::string(err));
  }
};

class AuthFailureError: public ApiError {
public:
  AuthFailureError(std::string message = "Invalid Credentials") {
    ApiError(UNAUTHORIZED, message);
  }
};

class InternalError: public ApiError {
public:
  InternalError(std::string message = "Internal error") {
    ApiError(INTERNAL_ERROR, message);
  }
};

class BadRequestError: public ApiError {
public:
  BadRequestError(std::string message = "Bad Request") {
    ApiError(BAD_REQUEST, message);
  }
};

class NotFoundError: public ApiError {
public:
  NotFoundError(std::string message = "Not Found") {
    ApiError(NOT_FOUND, message);
  }
};

class ForbiddenError: public ApiError {
public:
  ForbiddenError(std::string message = "Permission denied") {
    ApiError(FORBIDDEN, message);
  }
};

class NoEntryError: public ApiError {
public:
  NoEntryError(std::string message = "Entry don't exists") {
    ApiError(NO_ENTRY, message);
  }
};

class BadTokenError: public ApiError {
public:
  BadTokenError(std::string message = "Token is not valid") {
    ApiError(BAD_TOKEN, message);
  }
};

class TokenExpiredError: public ApiError {
public:
  TokenExpiredError(std::string message = "Token is expired") {
    ApiError(TOKEN_EXPIRED, message);
  }
};

class NoDataError: public ApiError {
public:
  NoDataError(std::string message = "No data available") {
    ApiError(NO_DATA, message);
  }
};

class AccessTokenError: public ApiError {
public:
  AccessTokenError(std::string message = "Invalid access token") {
    ApiError(ACCESS_TOKEN, message);
  }
};

#endif // CORE_API_ERROR_H
