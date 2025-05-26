#ifndef ROUTES_API_KEY_H
#define ROUTES_API_KEY_H

#define ROUTES_API_KEY_SCHEMA " \
{ \
    id: { \
        type: \"object\", \
        required: [\"id\"], \
        properties: { \
            id: { type: \"string\" }, \
        } \
    }, \
    create: { \
        type: \"object\", \
        required: [\"key\", \"status\", \"permission\"], \
        properties: { \
            key: { type: \"string\" }, \
            status: { type: \"boolean\" }, \
            permission: { type: \"array\" }, \
        } \
    }, \
    update: { \
        type: \"object\", \
        required: [\"key\", \"status\", \"permission\"], \
        properties: { \
            key: { type: \"string\" }, \
            status: { type: \"boolean\" }, \
            permission: { type: \"array\" }, \
        } \
    }, \
    findByKey: { \
        type: \"object\", \
        required: [\"key\"], \
        properties: { \
            key: { type: \"string\" } \
        } \
    } \
} \
"

#endif // ROUTES_API_KEY_H
