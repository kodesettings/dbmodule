#ifndef HELPERS_UTILS_H
#define HELPERS_UTILS_H

#include <string>
#include <boost/json.hpp>

// TODO: implement this function
bool isArrayOfStrings(std::string permissions) { return false; }

struct User {
  std::string userId;
  std::string email;
  std::string fullname;
};

struct JwtToken {
  User user;
  std::string usage;
};

template<typename T>
static T json_object_to_object(const std::string &req) {
  boost::json::value object = boost::json::parse(req);
  return boost::json::value_to<T>(object);
}

template<typename T>
static std::string object_to_json_object(T object) {
  boost::json::object out = boost::json::value_from(object);
  return boost::json::serialize(out);
}

#endif // HELPERS_UTILS_H
