#include "../database/model/refresh_token.hpp"
#include "../database/model/user.hpp"
#include "../core/api_error.hpp"
#include "../core/api_response.hpp"
#include "../helpers/context.hpp"
#include "../helpers/mimecontent.hpp"
#include "../helpers/utils.hpp"

// POST - /auth/login - Login Handler
void login(const std::string req, std::string *resp) {
  std::string email, password, deviceIdentifier; // TODO: get this from req

  bool exist;
  auto user = findUserByEmail(email, &exist);
  if (!exist) {
    *resp = NotFoundError("User not found").prepare();
    return;
  }

  const bool isPasswordCorrect = std::string(password).compare(user.password);

  if (!isPasswordCorrect) {
    *resp = BadRequestError("wrong password").prepare();
    return;
  }

  std::string claim = object_to_json_object<struct JwtToken>(JwtToken{
    .user = JwtUser{
      .userId = std::to_string(user._id),
      .email = user.email,
      .fullname = user.fullname
    },
    .usage = "auth-access"
  });

  auto accessToken = jwt::create()
    .set_type("JWS")
    .set_issuer("auth0")
    .set_payload_claim("sample", jwt::claim(claim))
    .set_expires_in(std::chrono::seconds{ACCESS_TOKEN_VALIDITY})
    .sign(jwt::algorithm::hs256{JWT_SECRET});

  bool isRefreshTokenAlreadyCreated;
  auto __refreshToken = findRefreshTokenByDeviceIdentifier(deviceIdentifier, &isRefreshTokenAlreadyCreated);
  if (isRefreshTokenAlreadyCreated) {
    removeRefreshToken(__refreshToken._id);
  }

  claim = object_to_json_object<struct JwtToken>(JwtToken{
    .user = JwtUser{
      .userId = std::to_string(user._id),
      .email = user.email,
      .fullname = user.fullname
    },
    .usage = "auth-refresh"
  });

  auto refreshToken = jwt::create()
    .set_type("JWS")
    .set_issuer("auth0")
    .set_payload_claim("sample", jwt::claim(claim))
    .set_expires_in(std::chrono::seconds{REFRESH_TOKEN_VALIDITY})
    .sign(jwt::algorithm::hs256{JWT_SECRET});


  RefreshToken refreshTokenDBObj{
    ._id = 0,
    .user = std::to_string(user._id),
    .token = std::string(refreshToken),
    .deviceIdentifier = deviceIdentifier,
    .createdAt = 0, // TODO: set creation time
    .expiresAt = 0 // TODO: set expiration time
  };

  if (!createRefreshToken(refreshTokenDBObj)) {
    *resp = InternalError("database problem").prepare();
    return;
  }

  std::string jsonObj; // TODO: return accessToken, refreshToken and userId
  *resp = SuccessResponse<std::string>("success", jsonObj).prepare();
}
