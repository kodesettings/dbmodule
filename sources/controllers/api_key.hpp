#ifndef CONTROLLERS_API_KEY_HPP
#define CONTROLLERS_API_KEY_HPP

#include "../database/model/api_key.hpp"
#include "../core/api_error.hpp"
#include "../core/api_response.hpp"
#include "../helpers/context.hpp"
#include "../helpers/utils.hpp"

void findAll(const std::string req, std::string *resp);
void create(const std::string req, std::string *resp);
void update(const std::string req, std::string *resp);
void remove(const std::string req, std::string *resp);
void findById(const std::string req, std::string *resp);
void findByKey(const std::string req, std::string *resp);


#endif // CONTROLLERS_API_KEY_HPP
