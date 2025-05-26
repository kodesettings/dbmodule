#include <iostream>
#ifdef __EMSCRIPTEN__
#include <emscripten.h>
#endif
#include "api_error.hpp"
#include "api_response.hpp"
#include "../database/repository/api_key.cpp"
#include "../database/repository/refresh_token.cpp"
#include "../database/repository/user.cpp"
// TODO: add javascript binding functions here
