#ifndef HELPERS_VALIDATOR_HPP
#define HELPERS_VALIDATOR_HPP

#include "../core/api_error.hpp"

#include "valijson/adapters/picojson_adapter.hpp"
#include "valijson/validation_results.hpp"
#include "valijson/schema_parser.hpp"
#include "valijson/validator.hpp"

#define BODY "body"
#define HEADER "headers"
#define QUERY "query"
#define PARAM "params"

bool isUrlEndPoint(const std::string &url) {
  if (url.find("://") != 0) return false;
  return true;
}

bool isAuthBearer(const std::string &value) {
  if (value.substr(0, 7).compare("Bearer ") != 0) return false;
  if (value.rfind(" ") != 0) return false;
  return true;
}

#include "picojson/picojson.h"

picojson::value jsonParse(std::string serialized, picojson::value def) {
  picojson::value v;
  std::string err;
  picojson::parse(v, serialized.data(), serialized.data() + serialized.size(), &err);

  if (!err.empty()) {
    return def;
  }

  return v;
}

bool isJsonSchemaValid(const std::string &schemaStr, const std::string &source, const std::string &req) {
  auto validatorSchema = std::make_shared<valijson::Schema>();
  auto schemaJson = jsonParse(schemaStr, picojson::value{});
  auto schemaAdapter = valijson::adapters::PicoJsonAdapter(schemaJson);
  valijson::SchemaParser parser;
  parser.populateSchema(schemaAdapter, *validatorSchema);

  auto targetJson = jsonParse(req, picojson::value{});
  auto targetAdapter = valijson::adapters::PicoJsonAdapter(targetJson);

  valijson::ValidationResults results;
  auto validator = valijson::Validator();
  return validator.validate(*validatorSchema, targetAdapter, &results);
}

#endif // HELPERS_VALIDATOR_HPP
