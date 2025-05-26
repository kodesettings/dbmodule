#ifndef DATABASE_MODEL_API_KEY_H
#define DATABASE_MODEL_API_KEY_H

#include "../../config/env_configs.hpp"
#include "../../helpers/context.hpp"

#define PERMISSION_GENERAL "GENERAL"
#define PERMISSION_SUPER_ADMIN "SUPER_ADMIN"

enum Status {
  STATUS_ACTIVE = 1,
  STATUS_INACTIVE = 0
};

struct ApiKey {
  uint64 _id;
  std::string key;
  std::vector<std::string> permissions;
  uint16 status;
  uint64 createdAt;
  uint64 updatedAt;
};

ApiKey* findAllApiKeys();
bool createApiKey(ApiKey apiKey);
bool updateApiKey(ApiKey apiKey);
bool removeApiKey(uint64 id);
ApiKey findApiKeyById(uint64 id, bool *found);
ApiKey findApiKeyByKey(std::string key, bool *found);
ApiKey findApiKeyByKeyButId(uint64 id, std::string key, bool *found);

bool createSuperAdminApiKey(void) {
  ApiKey *apiKeys = findAllApiKeys();

  if (apiKeys != nullptr) {
    return true;
  }

  auto apiKey = ApiKey{
    _id: 0,
    key: SUPER_ADMIN_API_KEY,
    permissions: {PERMISSION_SUPER_ADMIN},
    status: STATUS_ACTIVE,
    createdAt: 0, // TODO: add timestamp from helper method
    updatedAt: 0  // TODO: add timestamp from helper method
  };

  // save the default document
  return createApiKey(apiKey);
}

#endif // DATABASE_MODEL_API_KEY_H
