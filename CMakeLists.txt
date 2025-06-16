cmake_minimum_required(VERSION 3.17)

project(dbmodule)

add_library(server STATIC
  "sources/auth/api_key/api_key.cpp"
  "sources/auth/api_key/api_key.hpp"
  "sources/auth/api_key/schema.hpp"
  "sources/auth/middlewares/authentication.cpp"
  "sources/auth/middlewares/authorization.cpp"
  "sources/auth/middlewares/middlewares.hpp"
  "sources/auth/middlewares/schema.hpp"
  "sources/config/env_configs.hpp"
  "sources/controllers/api_key.cpp"
  "sources/controllers/user.cpp"
  "sources/core/api_error.hpp"
  "sources/core/api_response.hpp"
  "sources/core/worker.cpp"
  "sources/database/model/api_key.hpp"
  "sources/database/model/refresh_token.hpp"
  "sources/database/model/user.hpp"
  "sources/database/repository/api_key.cpp"
  "sources/database/repository/refresh_token.cpp"
  "sources/database/repository/user.cpp"
  "sources/helpers/context.hpp"
  "sources/helpers/mimecontent.hpp"
  "sources/helpers/utils.hpp"
  "sources/helpers/validator.hpp"
  "sources/routes/api_key/router.cpp"
  "sources/routes/api_key/schema.hpp"
  "sources/routes/user/router.cpp"
  "sources/routes/user/schema.hpp"
)

# this is required for 3rd party dependencies
# NOTE: set your own path to the folder
target_include_directories(server PUBLIC "/usr/local/include/")
target_include_directories(server PUBLIC "/usr/local/include/picojson/")
