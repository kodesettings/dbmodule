package config

import (
	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
)

var db = map[string]string{
    name:            cloudflare.Getenv("DB_NAME"),
    host:            cloudflare.Getenv("DB_HOST"),
    port:            cloudflare.Getenv("DB_PORT"),
    minPoolSize:     cloudflare.Getenv("DB_MIN_POOL_SIZE"),
    maxPoolSize:     cloudflare.Getenv("DB_MAX_POOL_SIZE"),
}

var tokenInfo = map[string]string{
    jwtSecret:            cloudflare.Getenv("JWT_SECRET"),
    accessTokenValidity:  cloudflare.Getenv("ACCESS_TOKEN_VALIDITY_SEC"),
    refreshTokenValidity: cloudflare.Getenv("JWT_SREFRESH_TOKEN_VALIDITY_SECECRET"),
    issuer:               cloudflare.Getenv("TOKEN_ISSUER"),
    audience:             cloudflare.Getenv("TOKEN_AUDIENCE"),
}

const environment = cloudflare.Getenv("NODE_ENV");
const superAdminApiKey = cloudflare.Getenv("SUPER_ADMIN_API_KEY");
