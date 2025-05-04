package helpers

import (
	"net/http"

	_ "github.com/syumai/workers"
	_ "github.com/syumai/workers/cloudflare/d1" // register driver

)

type LocalUserResponse http.ResponseWriter
type LocalUserRequest http.Request

