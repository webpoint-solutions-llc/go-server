package hash

import (
	"github.com/matthewhartstonge/argon2"
)

func Generate(str string) (string, error) {
	argon := argon2.DefaultConfig()
	hash, err := argon.HashEncoded([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Verify(str string, hash []byte) bool {
	ok, err := argon2.VerifyEncoded([]byte(str), hash)
	if !ok || err != nil {
		return false
	}
	return true
}
