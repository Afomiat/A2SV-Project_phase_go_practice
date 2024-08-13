package Infrastructure

import (
	"github.com/golang-jwt/jwt"
	"task1.go/task7/task_manager/Domain"
)

func GenerateToken(user Domain.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["username"] = user.Username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
