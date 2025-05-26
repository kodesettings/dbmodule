#ifndef CONFIG_ENV_CONFIGS_H
#define CONFIG_ENV_CONFIGS_H

/**
 * database variables from wrangler.toml
 */

#define DB_NAME ""
#define DB_HOST ""
#define DB_PORT 9999
#define MIN_POOL_SIZE 1
#define MAX_POOL_SIZE 100

/**
 * token info env variables from wrangler.toml
 */
#define JWT_SECRET ""
#define ACCESS_TOKEN_VALIDITY 100000
#define REFRESH_TOKEN_VALIDITY 100000
#define ISSUER "cloudflare workers"
#define AUDIENCE "public"


/**
 * Environment variables for API access
 */
#define ENVIRONMENT "dev"
#define SUPER_ADMIN_API_KEY "941820ce-39fc-11f0-bb76-2ba6e0b4a99c"

#endif // CONFIG_ENV_CONFIGS_H
