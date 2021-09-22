package common

import (
	"fmt"
	"os"
	"time"

	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(user entity.User) string
	VerifyToken(token string) (jwt.Claims, error)
	GeneratePairToken(user entity.User) (tokenAccess string, refreshToken string)
}

type jwtService struct {
	secretKey string
	issuer    string
}

type jwtCustomClaim struct {
	UserID   uint64 `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_key"
	}
	return secret
}

func getRefreshSecret() string {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		secret = "refresh_default_secret"
	}
	return secret
}

func NewJWTServ() JWTService {
	return &jwtService{
		issuer:    "dorman",
		secretKey: getJwtSecret(),
	}
}

func (c *jwtService) GenerateRefresh(userId uint64) string {
	rtSecret := getRefreshSecret()
	token := jwt.New(jwt.SigningMethodHS256)
	rtClaim := token.Claims.(jwt.MapClaims)
	rtClaim["user_id"] = userId
	rtClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := token.SignedString([]byte(rtSecret))
	if err != nil {
		panic(err)
	}
	return rt
}

func (c *jwtService) GenerateToken(user entity.User) (tokenAccess string) {
	ttl := time.Now().Add(time.Microsecond.Round(12)).Unix()
	claims := &jwtCustomClaim{
		UserID:   user.ID,
		Name:     user.Name,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ttl,
			Issuer:    c.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenAccess, err := token.SignedString([]byte(c.secretKey))
	if err != nil {
		panic(err)
	}
	return
}

func (c *jwtService) GeneratePairToken(user entity.User) (tokenAccess string, refreshToken string) {
	tokenAccess = c.GenerateToken(user)
	refreshToken = c.GenerateRefresh(user.ID)
	return
}

/*
	Using jwt.Parse to verify token and return error and jwt Token Struct
	To check Valid or Not
*/
func (c *jwtService) VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(c.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwtCustomClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("not valid token")
}
