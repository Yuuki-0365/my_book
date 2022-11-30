package tool

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("SmallRedBook")

type Claims struct {
	UserId    string `json:"user_id"` // 唯一标识
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(userId string, authority int) (string, error) {
	now := time.Now()
	expireTime := now.Add(2 * time.Hour)
	claims := Claims{
		UserId:    userId,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "my_book", // 签发者
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SignedString 方法根据传入的空接口类型参数 key，返回完整的签名令牌。
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
