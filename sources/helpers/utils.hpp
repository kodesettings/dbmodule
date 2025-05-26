#ifndef HELPERS_UTILS_H
#define HELPERS_UTILS_H

#include "context.hpp"

// TODO: implement this function
bool isArrayOfStrings(std::string permissions) { return false; }

struct JwtUser {
  std::string userId;
  std::string email;
  std::string fullname;
};

struct JwtToken {
  JwtUser user;
  std::string usage;
};

template<typename T>
static T json_object_to_object(const std::string &req) {
  boost::json::value object = boost::json::parse(req);
  return boost::json::value_to<T>(object);
}

template<typename T>
static std::string object_to_json_object(T object) {
  boost::json::value out = boost::json::value_from(object);
  return boost::json::serialize(out);
}

#endif // HELPERS_UTILS_H
