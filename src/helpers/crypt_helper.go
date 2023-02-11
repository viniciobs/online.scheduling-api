package helpers

import (
	"encoding/base64"
)

func Crypt(value string) string {
	for i := 0; i < 3; i++ {
		value = base64.StdEncoding.EncodeToString([]byte(value))
	}

	return value
}

func Decrypt(value string) string {
	for i := 0; i < 3; i++ {
		res, _ := base64.StdEncoding.DecodeString(value)
		value = string(res)
	}

	return value
}
