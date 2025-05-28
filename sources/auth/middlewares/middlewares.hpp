#ifndef AUTH_MIDDLEWARES_MIDDLEWARES_HPP
#define AUTH_MIDDLEWARES_MIDDLEWARES_HPP

#include "../../database/model/refresh_token.hpp"
#include "../../core/api_error.hpp"
#include "../../core/api_response.hpp"
#include "../../helpers/context.hpp"
#include "../../helpers/utils.hpp"

void authentication(const std::string req, std::string *resp);
void authorization(const std::string req, std::string *resp);

#endif // AUTH_MIDDLEWARES_MIDDLEWARES_HPP
