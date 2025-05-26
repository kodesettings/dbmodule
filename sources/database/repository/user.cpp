#include "../model/user.hpp"

User* findAllUsers() {
  // TODO: return all rows from the table
  // eg.: SELECT * FROM user ORDER BY created_at DESC;
  return new User[10];
}

bool createUser(User user) {
  auto timestamp = std::chrono::system_clock::now().time_since_epoch();
  auto localtime = std::chrono::duration_cast<std::chrono::seconds>(timestamp).count();
  user.createdAt = localtime;
  user.updatedAt = localtime;

  // TODO: create an entry in the table and return it's object
  // eg.: INSERT INTO user (createdAt, updatedAt) VALUES (?, ?, ?);

  return true;
}

bool updateUser(User user) {
  auto timestamp = std::chrono::system_clock::now().time_since_epoch();
  auto localtime = std::chrono::duration_cast<std::chrono::seconds>(timestamp).count();
  user.updatedAt = localtime;

  // TODO: update entry in table
  // eg.: UPDATE user SET (fullname = ?, email = ?, password = ?, isEmailVerified = ?, roles = ?, createdAt = ?, updatedAt = ?) WHERE Id = ?;

  return true;
}

bool removeUser(uint64 id) {
  // TODO: remove entry from table
  // eg.: DELETE FROM user WHERE Id = ?;

  return true;
}

User findUserById(uint64 id, bool *found) {
  // TODO: select row by id
  // eg.: SELECT * FROM user WHERE Id = ?;
  return User{};
}

User findUserByEmail(std::string email, bool *found) {
  // TODO: select row by email
  // eg.: SELECT * FROM user WHERE Email = ?;
  return User{};
}
