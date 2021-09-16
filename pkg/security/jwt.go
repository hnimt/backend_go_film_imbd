package security

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWT_KEY = "asdasdasdasdasd"
)

func GenToken(email string, role string) (string, error) {

	// claims := JwtCustomClaims{
	// 	UserId: user.UserID,
	// 	Role:   user.Role,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString([]byte(PrivateKey))
	// if err != nil {
	// 	return "", err
	// }
	// return tokenString, nil

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}
	return t, nil
}
