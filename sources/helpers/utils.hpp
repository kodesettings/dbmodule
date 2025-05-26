#ifndef HELPERS_UTILS_H
#define HELPERS_UTILS_H

#include <string>
#include <boost/json.hpp>
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

#include "../database/model/api_key.hpp"
#include "../database/model/refresh_token.hpp"
#include "../database/model/user.hpp"
namespace boost::json {
  void tag_invoke(value_from_tag, value& jv, const ApiKey &apiKey) {
    jv = {
      {"id", apiKey._id},
      {"key", apiKey.key},
      {"permissions", value_from(apiKey.permissions)},
      {"status", apiKey.status},
      {"createdAt", apiKey.createdAt},
      {"updatedAt", apiKey.updatedAt}
    };
  }

  ApiKey tag_invoke(value_to_tag<ApiKey>, const value& jv) {
    const object& obj = jv.as_object();
    return ApiKey{
      value_to<uint64>(obj.at("id")),
      value_to<std::string>(obj.at("key")),
      value_to<std::vector<std::string>>(obj.at("permissions")),
      value_to<uint16>(obj.at("status")),
      value_to<uint64>(obj.at("createdAt")),
      value_to<uint64>(obj.at("updatedAt"))
    };
  }

  void tag_invoke(value_from_tag, value& jv, const RefreshToken &refreshToken) {
    jv = {
      {"id", refreshToken._id},
      {"user", refreshToken.user},
      {"token", refreshToken.token},
      {"deviceIdentifier", refreshToken.deviceIdentifier},
      {"createdAt", refreshToken.createdAt},
      {"expiresAt", refreshToken.expiresAt}
    };
  }

  RefreshToken tag_invoke(value_to_tag<RefreshToken>, const value& jv) {
    const object& obj = jv.as_object();
    return RefreshToken{
      value_to<uint64>(obj.at("id")),
      value_to<std::string>(obj.at("user")),
      value_to<std::string>(obj.at("token")),
      value_to<std::string>(obj.at("deviceIdentifier")),
      value_to<uint64>(obj.at("createdAt")),
      value_to<uint64>(obj.at("expiresAt"))
    };
  }

  void tag_invoke(value_from_tag, value& jv, const User &user) {
    jv = {
      {"id", user._id},
      {"fullname", user.fullname},
      {"email", user.email},
      {"password", user.password},
      {"isEmailVerified", user.isEmailVerified},
      {"roles", value_from(user.roles)},
      {"createdAt", user.createdAt},
      {"updatedAt", user.updatedAt}
    };
  }

  User tag_invoke(value_to_tag<User>, const value& jv) {
    const object& obj = jv.as_object();
    return User{
      value_to<uint64>(obj.at("id")),
      value_to<std::string>(obj.at("fullname")),
      value_to<std::string>(obj.at("email")),
      value_to<std::string>(obj.at("password")),
      value_to<bool>(obj.at("isEmailVerified")),
      value_to<std::vector<std::string>>(obj.at("roles")),
      value_to<uint64>(obj.at("createdAt")),
      value_to<uint64>(obj.at("updatedAt"))
    };
  }
}

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
