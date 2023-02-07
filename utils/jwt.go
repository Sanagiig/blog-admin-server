package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-blog/global/settings"
	"time"
)

const formatTimeStr = "2006-01-02 15:04:05"

type TokenInfo struct {
	ID       string
	Username string
	CreateAt time.Time
}

func GenerateToken(id string, username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"createAt": time.Now().Format(formatTimeStr),
	})
	res, err := t.SignedString([]byte(settings.JWTKey))
	if err != nil {
		return "", err
	}
	return res, nil
}

func secret() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.JWTKey), nil
	}
}
func ParseToken(token string) (*TokenInfo, error) {
	tokenInfo := &TokenInfo{}
	t, err := jwt.Parse(token, secret())
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("解析到非法Token")
	}

	uc, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token 解析数据错误")
	}

	time, err := time.Parse(formatTimeStr, uc["createAt"].(string))
	if err != nil {
		return nil, err
	}

	tokenInfo.Username = uc["username"].(string)
	tokenInfo.ID = uc["id"].(string)
	tokenInfo.CreateAt = time
	return tokenInfo, err
}
