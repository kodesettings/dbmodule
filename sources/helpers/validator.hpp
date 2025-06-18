#ifndef HELPERS_VALIDATOR_HPP
#define HELPERS_VALIDATOR_HPP

#include "../core/api_error.hpp"

#include "valijson/adapters/boost_json_adapter.hpp"
#include "valijson/schema.hpp"
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

bool isJsonSchemaValid(const std::string &schemaJson, const std::string &source, const std::string &req) {
  boost::system::error_code ec;
  auto schemaDoc = boost::json::parse(schemaJson, ec);
  if (ec) {
    std::cerr << "Error parsing schema json: " << ec.message() << std::endl;
    return false;
  }

  valijson::Schema schema;
  valijson::SchemaParser schemaParser;

  valijson::adapters::BoostJsonAdapter schemaAdapter(schemaDoc);
  schemaParser.populateSchema(schemaAdapter, schema);

  auto targetDoc = boost::json::parse(req, ec);
  if (ec) {
    std::cerr << "Error parsing target json: " << ec.message() << std::endl;
    return false;
  }

  valijson::Validator validator;
  valijson::ValidationResults results;
  valijson::adapters::BoostJsonAdapter targetAdapter(targetDoc);
  return validator.validate(schema, targetAdapter, &results);
}

#endif // HELPERS_VALIDATOR_HPP
