#include "user.hpp"

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

  auto claim = generateClaim(user, "auth-access");
  auto accessToken = createJwtToken(claim, ACCESS_TOKEN_VALIDITY);

  bool isRefreshTokenAlreadyCreated;
  auto __refreshToken = findRefreshTokenByDeviceIdentifier(deviceIdentifier, &isRefreshTokenAlreadyCreated);
  if (isRefreshTokenAlreadyCreated) {
    removeRefreshToken(__refreshToken._id);
  }

  claim = generateClaim(user, "auth-refresh");
  auto refreshToken = createJwtToken(claim, REFRESH_TOKEN_VALIDITY);

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

  std::string jsonObj = std::string("{\"accessToken\": \"")
    .append(accessToken)
    .append("\", \"refreshToken\": \"")
    .append(refreshToken)
    .append("\", \"userId\": \"")
    .append(std::to_string(user._id))
    .append("\"");

  *resp = SuccessResponse<std::string>("success", jsonObj).prepare();
}

// POST - /auth/register - Register Handler
void registering(const std::string req, std::string *resp) {
  auto user = json_object_to_object<struct User>(req);

  bool isEmailInUse;
  findUserByEmail(user.email, &isEmailInUse);
  if (isEmailInUse) {
    *resp = BadRequestError("Email is Already Taken").prepare();
    return;
  }

  bool success = createUser(user);
  if (success) {
    //Send Welcome Email
    //sendEmail(email, fullname, "Welcome!", "Thank you for chosing us");
  }

  *resp = SuccessResponse<std::string>("success", "{}").prepare();
}

// POST - /auth/access-token - Create new access token
void newAccessToken(const std::string req, std::string *resp) {
  std::string deviceIdentifier, authHeader, token;

  auto decodedToken = jwt::decode(token);
  if (decodedToken.get_payload_json()["usage"].to_str() != "auth-refresh") {
    *resp = BadRequestError("Token is not defined").prepare();
    return;
  }

  bool isUserLoggedIn;
  findRefreshTokenByDeviceIdentifier(deviceIdentifier, &isUserLoggedIn);

  if (!isUserLoggedIn) {
    *resp = ForbiddenError("Login Again").prepare();
    return;
  }

  try {
    auto verify = jwt::verify();
    verify.verify(decodedToken);
  } catch (...) {
    bool found;
    auto refreshToken = findRefreshTokenByDeviceIdentifier(deviceIdentifier, &found);
    removeRefreshToken(refreshToken._id);
    *resp = ForbiddenError("Login Again").prepare();
    return;
  }

  std::string id = decodedToken.get_payload_json()["user"].get("id").to_str();

  bool found;
  User user = findUserById((uint64)std::stoi(id), &found);

  auto claim = generateClaim(user, "auth-access");
  auto accessToken = createJwtToken(claim, ACCESS_TOKEN_VALIDITY);

  *resp = SuccessResponse<std::string>("success", accessToken).prepare();
}

// POST - /auth/verify-email-request - Emails a link that verifies the email
void verifyEmailRequest(const std::string req, std::string *resp) {
  std::string email;

  bool exist;
  auto user = findUserByEmail(email, &exist);
  if (!exist) {
    *resp = NotFoundError("User not found").prepare();
    return;
  }

  if (user.isEmailVerified) {
    *resp = BadRequestError("the email is already verified").prepare();
    return;
  }

  auto claim = generateClaim(user, "emailVerify");
  auto token = createJwtToken(claim, ACCESS_TOKEN_VALIDITY);

  std::string verifyLink = std::string("/auth/verify-email/").append(token); // Replace your domain
  std::string emailBody = std::string("<a href=\"").append(verifyLink).append("\">click here for verifying email</a>");
  std::string mimetext = mimeTextBuilder(user.email, user.fullname, "", "", "verify email", emailBody);

  // TODO: send email

  *resp = SuccessResponse<std::string>("success", "{}").prepare();
}

// GET - /auth/verify-email/:token - Verifies user's email
void verifyEmail(const std::string req, std::string *resp) {
  std::string token;

  auto decodedToken = jwt::decode(token);

  try {
    auto verify = jwt::verify();
    verify.verify(decodedToken);
  } catch (...) {
    *resp = ForbiddenError("Permission denied").prepare();
    return;
  }

  if (decodedToken.get_payload_json()["usage"].to_str() != "emailVerify") {
    *resp = BadRequestError("Token is not valid").prepare();
    return;
  }

  std::string id = decodedToken.get_payload_json()["user"].get("id").to_str();

  bool found;
  User user = findUserById((uint64)std::stoi(id), &found);

  if (!found) {
    *resp = NotFoundError("User not found").prepare();
    return;
  }

  user.isEmailVerified = true;
  updateUser(user);

  *resp = SuccessResponse<std::string>("success", "{}").prepare();
}

// POST - /auth/forget-password - Forget Password Handler
void forgetPassword(const std::string req, std::string *resp) {
  std::string email;

  bool exist;
  auto user = findUserByEmail(email, &exist);
  if (!exist) {
    *resp = NotFoundError("User not found").prepare();
    return;
  }

  auto claim = generateClaim(user, "forgetPassword");
  auto token = createJwtToken(claim, ACCESS_TOKEN_VALIDITY);

  std::string verifyLink = std::string("/auth/reset-password/").append(token); // Replace your domain
  std::string emailBody = std::string("<a href=\"").append(verifyLink).append("\">click here for changing password</a>");
  std::string mimetext = mimeTextBuilder(user.email, user.fullname, "", "", "forget password", emailBody);

  // TODO: send email

  *resp = SuccessResponse<std::string>("success", "{}").prepare();
}

// POST - /reset-password/:token - Reset Password Handler
void resetForgottenPassword(const std::string req, std::string *resp) {
  std::string token, password, confirmPassword;

  auto decodedToken = jwt::decode(token);

  try {
    auto verify = jwt::verify();
    verify.verify(decodedToken);
  } catch (...) {
    *resp = ForbiddenError("Permission denied").prepare();
    return;
  }

  if (decodedToken.get_payload_json()["usage"].to_str() != "forgetPassword") {
    *resp = BadRequestError("Invalid token").prepare();
    return;
  }

  std::string id = decodedToken.get_payload_json()["user"].get("id").to_str();

  bool found;
  User user = findUserById((uint64)std::stoi(id), &found);

  if (!found) {
    *resp = NotFoundError("User not found").prepare();
    return;
  }

  user.password = password;
  updateUser(user);

  *resp = SuccessResponse<std::string>("success", "{}").prepare();
}

// PUT - /change-password/:id - Change Password Handler
void changePassword(const std::string req, std::string *resp) {
  std::string userId, currentPassword, newPassword, confirmPassword;

  bool found;
  User user = findUserById((uint64)std::stoi(userId), &found);

  if (!found) {
    *resp = NotFoundError("User not found").prepare();
    return;
  }

  const bool isPasswordCorrect = std::string(currentPassword).compare(user.password);

  if (!isPasswordCorrect) {
    *resp = BadRequestError("wrong password").prepare();
    return;
  }

  user.password = newPassword;
  updateUser(user);

  *resp = SuccessResponse<std::string>("success", "{}").prepare();
}
