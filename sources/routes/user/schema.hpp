#ifndef ROUTES_USER_H
#define ROUTES_USER_H

#define ROUTES_USER_SCHEMA " \
{ \
    login: { \
        type: \"object\", \
        required: [\"email\", \"password\", \"deviceIdentifier\"], \
        properties: { \
            email: { type: \"string\", format: \"email\" }, \
            password: { type: \"string\" }, \
            deviceIdentifier: { type: \"string\" } \
        } \
    }, \
    register: { \
        type: \"object\", \
        required: [\"email\", \"fullname\", \"password\", \"confirmPassword\"], \
        properties: { \
            email: { type: \"string\", format: \"email\" }, \
            fullname: { type: \"string\" }, \
            password: { type: \"string\" }, \
            confirmPassword: { type: \"string\" }, \
        } \
    }, \
    newAccessToken: { \
        type: \"object\", \
        required: [\"deviceIdentifier\"], \
        properties: { \
            deviceIdentifier: { type: \"string\" } \
        } \
    }, \
    verifyEmailRequest: { \
        type: \"object\", \
        required: [\"email\"], \
        properties: { \
            email: { type: \"string\", format: \"email\" } \
        } \
    }, \
    verifyEmail: { \
        type: \"object\", \
        required: [\"token\"], \
        properties: { \
            token: { type: \"token\" } \
        } \
    }, \
    forgetPassword: { \
        type: \"object\", \
        required: [\"email\"], \
        properties: { \
            email: { type: \"string\", format: \"email\" } \
        } \
    }, \
    resetForgettedPassword: { \
        type: \"object\", \
        required: [\"password\", \"confirmPassword\"], \
        properties: { \
            password: { type: \"string\" }, \
            confirmPassword: { type: \"string\" }, \
        } \
    }, \
    changePassword: { \
        type: \"object\", \
        required: [\"currentPassword\", \"newPassword\", \"confirmPassword\"], \
        properties: { \
            currentPassword: { type: \"string\" }, \
            newPassword: { type: \"string\" }, \
            confirmPassword: { type: \"string\" }, \
        } \
    } \
} \
"

#endif // ROUTES_USER_H
