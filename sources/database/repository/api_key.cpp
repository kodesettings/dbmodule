#include "../model/api_key.hpp"

ApiKey* findAllApiKeys() {
  // TODO: return all rows from the table
  // eg.: SELECT * FROM api_key ORDER BY created_at DESC;
  return new ApiKey[10];
}

bool createApiKey(ApiKey apiKey) {
  createSuperAdminApiKey();

  auto timestamp = std::chrono::system_clock::now().time_since_epoch();
  auto localtime = std::chrono::duration_cast<std::chrono::seconds>(timestamp).count();
  apiKey.createdAt = localtime;
  apiKey.updatedAt = localtime;

  // TODO: create an entry in the table and return it's object
  // eg.: INSERT INTO api_key (createdAt, updatedAt) VALUES (?, ?, ?);

  return true;
}

bool updateApiKey(ApiKey apiKey) {
  createSuperAdminApiKey();

  auto timestamp = std::chrono::system_clock::now().time_since_epoch();
  auto localtime = std::chrono::duration_cast<std::chrono::seconds>(timestamp).count();
  apiKey.updatedAt = localtime;

  // TODO: update entry in table
  // eg.: UPDATE api_key SET (key = ?, permission = ?, updatedAt = ?) WHERE Id = ?;

  return true;
}

bool removeApiKey(uint64 id) {
  createSuperAdminApiKey();

  // TODO: remove entry from table
  // eg.: DELETE FROM api_key WHERE Id = ?;

  return true;
}

ApiKey findApiKeyById(uint64 id, bool *found) {
  createSuperAdminApiKey();

  // TODO: select row by id
  // eg.: SELECT * FROM api_key WHERE Id = ?;
  return ApiKey{};
}

ApiKey findApiKeyByKey(std::string key, bool *found) {
  createSuperAdminApiKey();

  // TODO: select row by key
  // eg.: SELECT * FROM api_key WHERE Key = ?;
  return ApiKey{};
}

ApiKey findApiKeyByKeyButId(uint64 id, std::string key, bool *found) {
  createSuperAdminApiKey();

  // TODO: select row by id and key
  // eg.: SELECT * FROM api_key WHERE (Id = ?, Key = ?);
  return ApiKey{};
}
