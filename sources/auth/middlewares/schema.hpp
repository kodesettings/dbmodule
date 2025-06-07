#ifndef AUTH_MIDDLEWARES_SCHEMA_HPP
#define AUTH_MIDDLEWARES_SCHEMA_HPP

#define AUTH_MIDDLEWARES_SCHEMA " \
{ \
    auth: { \
        type: \"object\", \
        required: [\"authorization\"], \
        properties: { \
            authorization: { type: \"string\" } \
        } \
    } \
} \
"

#endif // AUTH_MIDDLEWARES_SCHEMA_HPP
