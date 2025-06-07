#include "middlewares.hpp"
#include "schema.hpp"

void authorization(const std::string req, std::string *resp) {
  std::string userId; // TODO: get this from json input

  bool exist;
  auto user = findUserById((uint64)std::stoi(userId), &exist);
  if (!exist) { return; }

  if (user.roles.size() == 0) {
failure:
    *resp = BadRequestError("permission error - don't have enough roles").prepare();
    return;
  }

  std::vector<std::string> roles = user.roles;
  std::vector<std::string> allRoles{ROLE_MEMBER, ROLE_ADMIN, ROLE_MODERATOR};
  // TODO: write an efficient algorithm to check if all roles are present in the input vector
}
