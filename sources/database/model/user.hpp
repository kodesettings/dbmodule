#ifndef DATABASE_MODEL_USER_HPP
#define DATABASE_MODEL_USER_HPP

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
  std::vector<std::string> roles;
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

namespace boost {
namespace json {
  void tag_invoke(value_from_tag, value& jv, const User &user) {
    jv = {
      {"id", user._id},
      {"fullname", user.fullname},
      {"email", user.email},
      {"password", user.password},
      {"isEmailVerified", user.isEmailVerified},
      {"roles", value_from(user.roles)},
      {"createdAt", user.createdAt},
      {"updatedAt", user.updatedAt}
    };
  }

  User tag_invoke(value_to_tag<User>, const value& jv) {
    const object& obj = jv.as_object();
    return User{
      value_to<uint64>(obj.at("id")),
      value_to<std::string>(obj.at("fullname")),
      value_to<std::string>(obj.at("email")),
      value_to<std::string>(obj.at("password")),
      value_to<bool>(obj.at("isEmailVerified")),
      value_to<std::vector<std::string>>(obj.at("roles")),
      value_to<uint64>(obj.at("createdAt")),
      value_to<uint64>(obj.at("updatedAt"))
    };
  }
}
}

#endif // DATABASE_MODEL_USER_HPP
