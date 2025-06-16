#ifndef HELPERS_VALIDATOR_HPP
#define HELPERS_VALIDATOR_HPP

#include "../core/api_error.hpp"

#define BODY "body"
#define HEADER "headers"
#define QUERY "query"
#define PARAM "params"

bool isUrlEndPoint(const std::string &url) {
  if (url.find("://") != 0) return false;
  return true;
}

bool isAuthBearer(const std::string &value) {
  if (value.substr(0, 7).compare("Bearer ") != 0) return false;
  if (value.rfind(" ") != 0) return false;
  return true;
}

bool isJsonSchemaValid(const std::string &schema, const std::string &source, const std::string &req) {
  // TODO: validate schema with valijson project
  return true;
}

#endif // HELPERS_VALIDATOR_HPP
