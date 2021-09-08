package common

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId string) string
	VerifyToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_key"
	}
	return secret
}

func NewJWTServ() JWTService {
	return &jwtService{
		issuer:    "dorman",
		secretKey: getJwtSecret(),
	}
}

func (c *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour.Round(12)).Unix(),
			Issuer:    c.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(c.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

/*
	Using jwt.Parse to verify token and return error and jwt Token Struct
	To check Valid or Not
*/
func (c *jwtService) VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error happend header alg %q", t.Header["alg"])
		}
		return []byte(c.secretKey), nil
	})
}
