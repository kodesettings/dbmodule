#include "api_key.hpp"
#include "schema.hpp"

void superAdminApiKeyCheck(const std::string req, std::string *resp) {
  auto apiKey = json_object_to_object<struct ApiKey>(req);

  auto permissions = apiKey.permissions;
  if (std::find(permissions.begin(), permissions.end(), PERMISSION_SUPER_ADMIN) != permissions.end()) {
    return;
  }

  *resp = ForbiddenError("Permission Denied").prepare();
}

void apiKeyCheck(const std::string req, std::string *resp) {
  std::string key = "x-api-key"; // TODO: get this value from header

  bool exist;
  auto apiKey = findApiKeyByKey(key, &exist);
  if (!exist) {
    *resp = ForbiddenError("No Api Key Found").prepare();
    return;
  }

  if (apiKey.permissions.size() == 0) {
    *resp = ForbiddenError("Permission Denied").prepare();
    return;
  }

  auto permissions = apiKey.permissions;
  if (std::find(permissions.begin(), permissions.end(), PERMISSION_GENERAL) != permissions.end()) {
    *resp = ForbiddenError("Permission Denied").prepare();
  }

  // TODO: add this serialized apiKey object into request
  std::string req_apiKey = object_to_json_object<ApiKey>(apiKey);
}
