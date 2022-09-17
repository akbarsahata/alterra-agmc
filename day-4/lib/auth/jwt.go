package auth

import (
	"time"

	"github.com/akbarsahata/alterra-agmc/day-4/config"
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GetToken(userID string) (string, error) {
	claims := JwtClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.Env["JWT_SECRET"]))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
