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
namespace boost::json {
  void tag_invoke(value_from_tag, value& jv, const ApiKey &apiKey) {
    jv = {
      {"id", apiKey._id},
      {"key", apiKey.key},
      {"permissions", apiKey.permissions[0]}, // TODO: add array
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
      value_to<std::string>(obj.at("permissions")),
      value_to<std::string>(obj.at("permissions")),
      value_to<uint16>(obj.at("status")),
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
