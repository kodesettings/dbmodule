cmake_minimum_required(VERSION 3.17)

project(dbmodule)

add_library(server STATIC "sources/config/env_configs.hpp" "sources/core/api_error.hpp" "sources/core/api_response.hpp" "sources/core/worker.cpp")
