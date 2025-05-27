#include <iostream>
#ifdef __EMSCRIPTEN__
#include <emscripten.h>
#endif
#include "api_error.hpp"
#include "api_response.hpp"

// NOTE: this is for debugging only
#include "../database/repository/api_key.cpp"
#include "../database/repository/refresh_token.cpp"
#include "../database/repository/user.cpp"
#include "../controllers/api_key.cpp"
#include "../controllers/user.cpp"

#include "../helpers/context.hpp"
#include "../helpers/mimecontent.hpp"
#include "../helpers/utils.hpp"
#include "../helpers/validator.hpp"

#include "../auth/api_key/api_key.cpp"
#include "../auth/middlewares/authentication.cpp"
#include "../auth/middlewares/authorization.cpp"
// ---

// TODO: add javascript binding functions here
