#ifndef CORE_API_RESPONSE_H
#define CORE_API_RESPONSE_H

// Helper code for the API consumer to understand the error and handle is accordingly
enum StatusCode {
  STATUS_SUCCESS = 10000,
  STATUS_FAILURE = 10001,
  STATUS_RETRY = 10002,
  STATUS_INVALID_ACCESS_TOKEN = 10003
};

enum ResponseStatus {
  HTTP_SUCCESS = 200,
  HTTP_BAD_REQUEST = 400,
  HTTP_UNAUTHORIZED = 401,
  HTTP_FORBIDDEN = 403,
  HTTP_NOT_FOUND = 404,
  HTTP_INTERNAL_ERROR = 500
};

#include <string>

class ApiResponse {
public:
  ApiResponse() {}
  ApiResponse(StatusCode statusCode, ResponseStatus responseStatus, const std::string &message) :
    statusCode_(statusCode), responseStatus_(responseStatus), message_(message) {}
public:
  std::string prepare() {
    char str[2000];
    sprintf(str, "{\"status\": \"%d\", \"statusText:\", \"%s\", \"headers:\" \"%s\"}", statusCode_, message_.c_str(), "{}");
    return str;
  }
private:
  StatusCode statusCode_;
  ResponseStatus responseStatus_;
  std::string message_;
};

class AuthFailureResponse: public ApiResponse {
public:
  AuthFailureResponse(std::string message = "Authentication Failure") {
    ApiResponse(StatusCode::STATUS_FAILURE, ResponseStatus::HTTP_UNAUTHORIZED, message);
  }
};

class NotFoundResponse: public ApiResponse {
public:
  NotFoundResponse(std::string message = "Not Found") {
    ApiResponse(StatusCode::STATUS_FAILURE, ResponseStatus::HTTP_NOT_FOUND, message);
  }
};

class ForbiddenResponse: public ApiResponse {
public:
  ForbiddenResponse(std::string message = "Forbidden") {
    ApiResponse(StatusCode::STATUS_FAILURE, ResponseStatus::HTTP_FORBIDDEN, message);
  }
};

class BadRequestResponse: public ApiResponse {
public:
  BadRequestResponse(std::string message = "Bad Parameters") {
    ApiResponse(StatusCode::STATUS_FAILURE, ResponseStatus::HTTP_BAD_REQUEST, message);
  }
};

class InternalErrorResponse: public ApiResponse {
public:
  InternalErrorResponse(std::string message = "Internal Error") {
    ApiResponse(StatusCode::STATUS_FAILURE, ResponseStatus::HTTP_INTERNAL_ERROR, message);
  }
};

template<typename T>
class SuccessResponse: public ApiResponse {
public:
  SuccessResponse(const std::string &message, T payload) {
    // TODO: handling template payload
    ApiResponse(StatusCode::STATUS_SUCCESS, ResponseStatus::HTTP_SUCCESS, message);
  }
};

class AccessTokenErrorResponse: public ApiResponse {
public:
  AccessTokenErrorResponse(std::string message = "Access token invalid") {
    ApiResponse(StatusCode::STATUS_INVALID_ACCESS_TOKEN, ResponseStatus::HTTP_UNAUTHORIZED, message);
  }
};

class RefreshTokenResponse: public ApiResponse {
public:
  RefreshTokenResponse(const std::string &message) {
    ApiResponse(StatusCode::STATUS_SUCCESS, ResponseStatus::HTTP_SUCCESS, message);
  }
};

#endif // CORE_API_RESPONSE_H
