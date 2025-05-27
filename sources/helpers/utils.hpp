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

namespace boost {
namespace json {
  void tag_invoke(value_from_tag, value& jv, const JwtUser &jwtUser) {
    jv = {
      {"userId", jwtUser.userId},
      {"email", jwtUser.email},
      {"fullname", jwtUser.fullname}
    };
  }

  JwtUser tag_invoke(value_to_tag<JwtUser>, const value& jv) {
    const object& obj = jv.as_object();
    return JwtUser{
      value_to<std::string>(obj.at("userId")),
      value_to<std::string>(obj.at("email")),
      value_to<std::string>(obj.at("fullname"))
    };
  }

  void tag_invoke(value_from_tag, value& jv, const JwtToken &jwtToken) {
    jv = {
      {"user", value_from(jwtToken.user)},
      {"usage", jwtToken.usage}
    };
  }

  JwtToken tag_invoke(value_to_tag<JwtToken>, const value& jv) {
    const object& obj = jv.as_object();
    return JwtToken{
      value_to<JwtUser>(obj.at("user")),
      value_to<std::string>(obj.at("usage"))
    };
  }
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

#include "../database/model/user.hpp"

std::string generateClaim(User user, const std::string &usage) {
  return object_to_json_object<struct JwtToken>(JwtToken{
    .user = JwtUser{
      .userId = std::to_string(user._id),
      .email = user.email,
      .fullname = user.fullname
    },
    .usage = usage
  });
}

std::string createJwtToken(std::string claim, uint64 validity) {
  return jwt::create()
    .set_type("JWS")
    .set_issuer("auth0")
    .set_payload_claim("sample", jwt::claim(claim))
    .set_expires_in(std::chrono::seconds{validity})
    .sign(jwt::algorithm::hs256{JWT_SECRET});
}

#endif // HELPERS_UTILS_H
