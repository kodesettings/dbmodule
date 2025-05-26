#ifndef DATABASE_MODEL_USER_H
#define DATABASE_MODEL_USER_H

#include "../../config/env_configs.hpp"
#include "../../helpers/context.hpp"

struct IRole {
  std::string MEMBER;
  std::string ADMIN;
  std::string MODERATOR;
};

#define ROLE_MEMBER "member"
#define ROLE_ADMIN "admin"
#define ROLE_MODERATOR "superadmin"

struct Role {
  uint64 _id;
  std::string role;
};

struct User {
  uint64 _id;
  std::string fullname;
  std::string email;
  std::string password;
  bool isEmailVerified;
  Role roles[3];
  uint64 createdAt;
  uint64 updatedAt;
};

User* findAllUsers();
bool createUser(User user);
bool updateUser(User user);
bool removeUser(uint64 id);
User findUserById(uint64 id, bool *found);
User findUserByEmail(std::string email, bool *found);

// Hashing Password Before Saving the User
User savePassword(User user) {
  // TODO: overwriting password by generating SHA-256 representation from plain text
  user.password = std::string("");
  return user;
}

#endif // DATABASE_MODEL_USER_H
