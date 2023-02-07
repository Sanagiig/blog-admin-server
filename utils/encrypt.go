package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSha256(data string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)

	//由于是十六进制表示，因此需要转换
	s := hex.EncodeToString(sum)
	return string(s), nil
}
