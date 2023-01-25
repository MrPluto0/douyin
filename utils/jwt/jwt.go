package jwt

import (
	"douyin/app/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.RegisteredClaims
	models.User
}

var (
	secret = []byte("16849841325189456f487")
	expire = 2 * time.Hour
)

func GenerateToken(user models.User) (string, error) {
	claims := &UserClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),             // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),             // 生效时间
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}

func ParseToken(token string) (*UserClaims, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}
	jwtToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	return jwtToken.Claims.(*UserClaims), nil
}
