package helpers

import (
	"fmt"
	"reflect"
	//permission "dbmodule/database/model/apiKey"
	jwt "github.com/golang-jwt/jwt/v5"
)

func is_array_of_strings(value any) bool {
	rt := reflect.TypeOf(value)
	if rt.Kind() == reflect.Array { if reflect.TypeOf(rt.Elem()).String() == "string" { return true } }
	return false
}

// func add_permissions(stringArray []string) []permission {
//  return stringArray.filter((value): value is Permission => Object.values(Permission).includes(value as Permission));
//}

var secretKey = []byte("secret-key")

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}
  
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

type User struct {
	userId string
	email string
	fullname string
}

type JwtToken struct {
	user User
	usage string
}
