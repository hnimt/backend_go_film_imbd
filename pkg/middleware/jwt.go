package middleware

import (
	"micro_backend_film/pkg/security"

	jwtware "github.com/gofiber/jwt/v3"
)

func JwtMiddleWare() {
	jwtware.New(jwtware.Config{
		SigningKey: []byte(security.JWT_KEY),
	})
}
