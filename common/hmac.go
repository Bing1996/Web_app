package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = 2 * time.Hour

var (
	signingKey        = []byte("MySecret")
	ErrorInvalidToken = errors.New("token无效")
)

type CustomClaims struct {
	UserID   int64  `json:"user_id,omitempty"`
	UserName string `json:"username,omitempty"`
	jwt.RegisteredClaims
}

// GenToken 基于HMAC签名方法生成Token
func GenToken(userId int64, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userId,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "bluebell",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		},
	}

	// 创建签名对象
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(signingKey)
}

// ParseToken HMAC签名的解析
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 通过jwt内置函数解析得到一个*jwt.Token对象并赋值
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claims进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, err
	}

	return nil, ErrorInvalidToken

}
