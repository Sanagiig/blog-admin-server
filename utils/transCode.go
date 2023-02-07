package utils

import "encoding/base64"

func EncodeBase64(str string) string {
	res := base64.URLEncoding.EncodeToString([]byte(str))
	return res
}

func DecodeBase64(str string) (string, error) {
	res, err := base64.URLEncoding.DecodeString(str)
	return string(res), err
}
