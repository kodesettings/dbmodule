//go:build js && wasm
package config

import (
	_ "github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
)

const (
	Name int = iota
	Host
	Port
	MinPoolSize
	MaxPoolSize
)

/**
 * database variables from wrangler.toml
 */
var DB = map[int]string{
	Name:            cloudflare.Getenv("DB_NAME"),
	Host:            cloudflare.Getenv("DB_HOST"),
	Port:            cloudflare.Getenv("DB_PORT"),
	MinPoolSize:     cloudflare.Getenv("DB_MIN_POOL_SIZE"),
	MaxPoolSize:     cloudflare.Getenv("DB_MAX_POOL_SIZE"),
}

const (
	JwtSecret int = iota
	AccessTokenValidity
	RefreshTokenValidity
	Issuer
	Audience
)

/**
 * token info env variables from wrangler.toml
 */
var TokenInfo = map[int]string{
	JwtSecret:            cloudflare.Getenv("JWT_SECRET"),
	AccessTokenValidity:  cloudflare.Getenv("ACCESS_TOKEN_VALIDITY_SEC"),
	RefreshTokenValidity: cloudflare.Getenv("JWT_SREFRESH_TOKEN_VALIDITY_SECECRET"),
	Issuer:               cloudflare.Getenv("TOKEN_ISSUER"),
	Audience:             cloudflare.Getenv("TOKEN_AUDIENCE"),
}

/**
 * Environment variables for API access
 */
var Environment string = cloudflare.Getenv("NODE_ENV");
var SuperAdminApiKey string = cloudflare.Getenv("SUPER_ADMIN_API_KEY");
