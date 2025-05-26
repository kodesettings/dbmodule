#ifndef DATABASE_MODEL_REFRESH_TOKEN_H
#define DATABASE_MODEL_REFRESH_TOKEN_H

#include "../../config/env_configs.hpp"
#include "../../helpers/context.hpp"

struct RefreshToken {
  uint64 _id;
  std::string user;
  std::string token;
  std::string deviceIdentifier;
  uint64 createdAt;
  uint64 expiresAt;
};

RefreshToken* findAllRefreshTokens();
bool createRefreshToken(RefreshToken refreshToken);
bool updateRefreshToken(RefreshToken refreshToken);
bool removeRefreshToken(uint64 id);
RefreshToken findRefreshTokenById(uint64 id, bool *found);
RefreshToken findRefreshTokenByDeviceIdentifier(std::string deviceIdentifier, bool *found);

std::string encrypt(const std::string &value) {
  // TODO: accept password, encrypt and output base64 of encrypted pwd
  return "";
}

std::string decrypt(const std::string &value) {
  // TODO: return ciphered text of decoded base64 string
  return "";
}

namespace boost::json {
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
}

#endif // DATABASE_MODEL_REFRESH_TOKEN_H
