#ifndef AUTH_MIDDLEWARES_H
#define AUTH_MIDDLEWARES_H

#define AUTH_MIDDLEWARES_SCHEMA " \
{ \
    auth: { \
        type: \"object\", \
        required: [\"authorization\"], \
        properties: { \
            deviceIdentifier: { type: \"string\" } \
        } \
    } \
} \
"

#endif // AUTH_MIDDLEWARES_SCHEMA
