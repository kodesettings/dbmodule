#ifndef AUTH_API_KEY_SCHEMA_HPP
#define AUTH_API_KEY_SCHEMA_HPP

#define AUTH_API_KEY_SCHEMA " \
{ \
    apiKey: { \
        type: \"object\", \
        required: [\"x-api-key\"], \
        properties: { \
            [\"x-api-key\"]: { type: \"string\" } \
        } \
    }, \
    auth: { \
        type: \"object\", \
        required: [\"authorization\"], \
        properties: { \
            deviceIdentifier: { type: \"string\" } \
        } \
    } \
} \
"

#endif // AUTH_API_KEY_SCHEMA_HPP
