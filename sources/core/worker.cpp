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
// ---

// TODO: add javascript binding functions here
