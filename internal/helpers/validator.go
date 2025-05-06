//go:build js && wasm
package helpers

import (
	"fmt"
	"strings"

	. "github.com/kodesettings/dbmodule/v2/internal/core"
	"github.com/goccy/go-json"
	"github.com/kaptinlin/jsonschema"
)

const (
	BODY = "body"
	HEADER = "headers"
	QUERY = "query"
	PARAM = "params"
)

func ValidateUrlEndpoint(value string) string {
	if (strings.Contains(value, "://")) { return "malformed URL"; }
	return value;
}

func ValidateAuthBearerToken(value string) string {
	if (!strings.HasPrefix(value, "Bearer ")) { return "invalid Bearer token"; }
	if (strings.Split(value, " ")[1] == "") { return "invalid Bearer token"; }
	return value;
}

const draft string = "2019-09" // json schema draft

func Validator(schema string, validationSource string, jsonText string, __execute func()) any {
		compiler := jsonschema.NewCompiler()
		output_schema, err := compiler.Compile([]byte(schema))
		if err != nil {
			r := ApiResponse{/*empty struct, it is assigned in function call*/}
			details := fmt.Sprint("Failed to compile schema: %v", err);
			return r.BadRequestResponse(details);
		}

		instance := map[string]interface{}{
			"name": "John Doe",
			"age":  19,
		}
		result := output_schema.Validate(instance)
		if !result.IsValid() {
			r := ApiResponse{/*empty struct, it is assigned in function call*/}
			details, _ := json.MarshalIndent(result.ToList(), "", "  ");
			return r.BadRequestResponse(string(details));
		}

		return func() { __execute(); }
	}
