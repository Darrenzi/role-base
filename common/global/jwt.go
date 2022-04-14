package global

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID       int
	Username string
	Roles    []string
	jwt.StandardClaims
}

func GenerateToken(id int, username string, roles []string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(Config.JWT.Expire) * time.Hour)
	claims := Claims{
		id,
		username,
		roles,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(Config.JWT.Secret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JWT.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
