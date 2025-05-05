//go:build js && wasm
package config

import (
	_ "github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
)

const (
	name int = iota
	host
	port
	minPoolSize
	maxPoolSize
)

/**
 * database variables from wrangler.toml
 */
var db = map[int]string{
	name:            cloudflare.Getenv("DB_NAME"),
	host:            cloudflare.Getenv("DB_HOST"),
	port:            cloudflare.Getenv("DB_PORT"),
	minPoolSize:     cloudflare.Getenv("DB_MIN_POOL_SIZE"),
	maxPoolSize:     cloudflare.Getenv("DB_MAX_POOL_SIZE"),
}

const (
	jwtSecret int = iota
	accessTokenValidity
	refreshTokenValidity
	issuer
	audience
)

/**
 * token info env variables from wrangler.toml
 */
var tokenInfo = map[int]string{
	jwtSecret:            cloudflare.Getenv("JWT_SECRET"),
	accessTokenValidity:  cloudflare.Getenv("ACCESS_TOKEN_VALIDITY_SEC"),
	refreshTokenValidity: cloudflare.Getenv("JWT_SREFRESH_TOKEN_VALIDITY_SECECRET"),
	issuer:               cloudflare.Getenv("TOKEN_ISSUER"),
	audience:             cloudflare.Getenv("TOKEN_AUDIENCE"),
}

/**
 * Environment variables for API access
 */
var environment string = cloudflare.Getenv("NODE_ENV");
var superAdminApiKey string = cloudflare.Getenv("SUPER_ADMIN_API_KEY");
