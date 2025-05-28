#include "api_key.hpp"

// GET - /apikey/find-all
void findAll(const std::string req, std::string *resp) {
  auto apiKeys = findAllApiKeys();

  *resp = SuccessResponse<struct ApiKey*>("success", apiKeys).prepare();
}

// POST - /apikey/create
void create(const std::string req, std::string *resp) {
  auto payload = json_object_to_object<struct ApiKey>(req);

  bool exist;
  auto apiKey = findApiKeyByKey(payload.key, &exist);
  if (!exist) { *resp = BadRequestError("Key is already taken").prepare(); return; }
/*
  const permissions = apiKey.permissions;
  if (!isArrayOfStrings(permissions)) apiKey.permissions = [];
*/
  auto result = createApiKey(apiKey);
  if (!result) { *resp = InternalError("database problem").prepare(); return; }
  std::string jsonObj = object_to_json_object<struct ApiKey>(apiKey);
  *resp = SuccessResponse<std::string>("success", jsonObj).prepare();
}

// PUT - /apikey/update
void update(const std::string req, std::string *resp) {
  auto payload = json_object_to_object<struct ApiKey>(req);
  auto id = payload._id;
  if (!id) { *resp = NotFoundError("ApiKey not found").prepare(); return; }

  bool exist;
  auto apiKey = findApiKeyById(id, &exist);
  if (exist) { *resp = BadRequestError("Key is already taken").prepare(); return; }
/*
  if (req.body.permissions && isArrayOfStrings(req.body.permissions))
    apiKey.permissions = req.body.permissions;
  if (req.body.status) apiKey.status = req.body.status;
  if (req.body.key) apiKey.key = req.body.key;
*/
  bool result = updateApiKey(apiKey);
  if (result) { SuccessResponse<struct ApiKey>("success", apiKey).prepare(); return; }
  else { InternalError("database problem").prepare(); return; }
}

// DELETE - /apikey/remove
void remove(const std::string req, std::string *resp) {
  auto payload = json_object_to_object<struct ApiKey>(req);
  auto id = payload._id;
  if (!id) { *resp = NotFoundError("ApiKey not found").prepare(); return; }

  bool exist;
  auto apiKey = findApiKeyById(id, &exist);
  if (exist) { *resp = BadRequestError("Key is already taken").prepare(); return; }

  bool result = removeApiKey(id);
  if (result) { SuccessResponse<struct ApiKey>("success", apiKey).prepare(); return; }
  else { InternalError("database problem").prepare(); return; }
}

// GET - /apikey/find-by-id
void findById(const std::string req, std::string *resp) {
  auto payload = json_object_to_object<struct ApiKey>(req);
  auto id = payload._id;
  if (!id) { *resp = NotFoundError("No ApiKey found").prepare(); return; }

  bool exist;
  auto apiKey = findApiKeyById(id, &exist);
  if (!exist) { *resp = InternalError("database problem").prepare(); return; }
  else { SuccessResponse<struct ApiKey>("success", apiKey).prepare(); return; }
}


// GET - /apikey/find-by-key
void findByKey(const std::string req, std::string *resp) {
  auto payload = json_object_to_object<struct ApiKey>(req);
  auto key = payload.key;
  if (key.empty()) { *resp = NotFoundError("No ApiKey Found").prepare(); return; }

  bool exist;
  auto apiKey = findApiKeyByKey(key, &exist);
  if (!exist) { *resp = InternalError("database problem").prepare(); return; }
  else { SuccessResponse<struct ApiKey>("success", apiKey).prepare(); return; }
}
