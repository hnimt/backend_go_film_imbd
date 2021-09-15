package security

import (
	"micro_backend_film/pkg/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWT_KEY = "secret"

func GenToken(user entity.User) (string, error) {

	claims := JwtCustomClaims{
		UserId: user.UserID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
