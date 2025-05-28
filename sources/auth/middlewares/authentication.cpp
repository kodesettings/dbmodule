#include "middlewares.hpp"

void authentication(const std::string req, std::string *resp) {
  std::string authHeader, deviceIdentifier; // TODO: get these from header

  if (authHeader.empty()) {
    *resp = BadRequestError("No authorization token provided").prepare();
    return;
  }

  auto token = authHeader.substr(0, 7); // remove Bearer prefix from token
  auto decodedToken = jwt::decode(token);

  if (decodedToken.get_payload_json().empty()) {
    *resp = BadRequestError("Invalid Token").prepare();
    return;
  }

  // TODO: overwrite request
  std::string userId = decodedToken.get_payload_json()["user"].get("id").to_str();

  bool isUserLoggedIn;
  findRefreshTokenByDeviceIdentifier(deviceIdentifier, &isUserLoggedIn);

  if (!isUserLoggedIn) {
    *resp = ForbiddenError("Login Again").prepare();
    return;
  }

  if (decodedToken.get_payload_json()["user"].get("usage").to_str() == "auth-access") {
    try {
      auto verify = jwt::verify();
      verify.verify(decodedToken);
    } catch (...) {
      *resp = BadRequestError("send refresh token").prepare();
    }
  } else {
    *resp = BadRequestError("token is not valid").prepare();
  }
}
