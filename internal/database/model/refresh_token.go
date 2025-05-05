//go:build js && wasm
package database_model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"fmt"
	. "github.com/kodesettings/dbmodule/v2/internal/config"
)

type RefreshToken struct {
	_id                     uint64 `json:"id"`
	user                    string `json:"user"`
	token                   string `json:"token"`
	deviceIdentifier        string `json:"deviceIdentifier"`
	createdAt               uint64 `json:"createdAt"`
	expiresAt               uint64 `json:"expiresAt"`
}

// https://aran.dev/posts/using-golang-crypto-aes-and-crypto-cipher-packages/
func encrypt(value string) (string, error) {
	block, err := aes.NewCipher([]byte(TokenInfo[JwtSecret]))
	if err != nil {
		return "", err
	}

	plainText := []byte(value)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(value string) (string, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(value)
	if err != nil {
		return "", fmt.Errorf("decoding base64: %w", err)
	}

	block, err := aes.NewCipher([]byte(TokenInfo[JwtSecret]))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
