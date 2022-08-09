package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (bool, string, error)
}

type JwtService struct {
	Username string
	jwt.RegisteredClaims
}

func NewService() *JwtService {
	return &JwtService{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func (s *JwtService) GenerateToken(username string) (string, error) {
	claim := JwtService{}
	claim.Username = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (bool, string, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &JwtService{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false, "0", err
	}

	if claim, ok := token.Claims.(*JwtService); ok && token.Valid {
		return true, claim.Username, nil
	} else {
		return false, claim.Username, err
	}

}
