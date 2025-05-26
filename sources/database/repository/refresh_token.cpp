#include "../model/refresh_token.hpp"

RefreshToken* findAllRefreshTokens() {
  // TODO: return all rows from the table
  // eg.: SELECT * FROM refresh_token ORDER BY created_at DESC;
  return new RefreshToken[10];
}

bool createRefreshToken(RefreshToken refreshToken) {
  auto timestamp = std::chrono::system_clock::now().time_since_epoch();
  auto localtime = std::chrono::duration_cast<std::chrono::seconds>(timestamp).count();
  refreshToken.createdAt = localtime;
  refreshToken.expiresAt = localtime;

  // TODO: create an entry in the table and return it's object
  // eg.: INSERT INTO refresh_token (createdAt, expiresAt) VALUES (?, ?, ?);

  return true;
}

bool updateRefreshToken(RefreshToken refreshToken) {
  // TODO: update entry in table
  // eg.: UPDATE refresh_token SET (user = ?, token = ?, deviceIdentifier = ?) WHERE Id = ?;

  return true;
}

bool removeRefreshToken(uint64 id) {
  // TODO: remove entry from table
  // eg.: DELETE FROM refresh_token WHERE Id = ?;

  return true;
}

RefreshToken findRefreshTokenById(uint64 id, bool *found) {
  // TODO: select row by id
  // eg.: SELECT * FROM refresh_token WHERE Id = ?;
  return RefreshToken{};
}

RefreshToken findRefreshTokenByDeviceIdentifier(std::string deviceIdentifier, bool *found) {
  // TODO: select row by device identifier
  // eg.: SELECT * FROM refresh_token WHERE deviceIdentifier = ?;
  return RefreshToken{};
}
