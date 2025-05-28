#ifndef AUTH_API_KEY_API_KEY_HPP
#define AUTH_API_KEY_API_KEY_HPP

#include "../../database/model/api_key.hpp"
#include "../../core/api_error.hpp"
#include "../../helpers/context.hpp"
#include "../../helpers/utils.hpp"
#include "../../helpers/validator.hpp"

void superAdminApiKeyCheck(const std::string req, std::string *resp);
void apiKeyCheck(const std::string req, std::string *resp);

#endif // AUTH_API_KEY_API_KEY_HPP
