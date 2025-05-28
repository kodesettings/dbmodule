#ifndef CONTROLLERS_USER_HPP
#define CONTROLLERS_USER_HPP

#include "../database/model/refresh_token.hpp"
#include "../database/model/user.hpp"
#include "../core/api_error.hpp"
#include "../core/api_response.hpp"
#include "../helpers/context.hpp"
#include "../helpers/mimecontent.hpp"
#include "../helpers/utils.hpp"

void login(const std::string req, std::string *resp);
void registering(const std::string req, std::string *resp);
void newAccessToken(const std::string req, std::string *resp);
void verifyEmailRequest(const std::string req, std::string *resp);
void verifyEmail(const std::string req, std::string *resp);
void forgetPassword(const std::string req, std::string *resp);
void resetForgottedPassword(const std::string req, std::string *resp);
void changePassword(const std::string req, std::string *resp);

#endif // CONTROLLERS_USER_HPP
