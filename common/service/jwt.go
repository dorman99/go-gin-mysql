package common

import (
	"fmt"
	"os"
	"time"

	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/dorman99/go_gin_mysql/helper"
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(user entity.User) string
	VerifyToken(token string) (jwt.Claims, error)
	GeneratePairToken(user entity.User) (tokenAccess string, refreshToken string)
}

type jwtService struct {
	secretKey    string
	issuer       string
	redisService RedisService
}

type jwtCustomClaim struct {
	UserID   uint64 `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	AccessId string `json:"acccesId"`
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

func NewJWTServ(redisClientService RedisService) JWTService {
	return &jwtService{
		issuer:       "dorman",
		secretKey:    getJwtSecret(),
		redisService: redisClientService,
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

func (c *jwtService) setToRedis(userId int, uuid string) (err error) {
	key := generateKey(userId)
	err = c.redisService.Set(key, uuid, 0)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func generateKey(userId int) (key string) {
	key = fmt.Sprintf("acces_user:%d", userId)
	return
}

func (c *jwtService) GenerateToken(user entity.User) (tokenAccess string) {
	ttl := time.Now().Add(time.Hour.Round(12)).Unix()
	uuid := helper.GenerateUUID()
	claims := &jwtCustomClaim{
		UserID:   user.ID,
		Name:     user.Name,
		Username: user.Username,
		AccessId: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ttl,
			Issuer:    c.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	go c.setToRedis(int(user.ID), uuid)

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
		key := generateKey(int(claims.UserID))
		accessId, err := c.redisService.Get(key)
		if err != nil {
			return nil, err
		} else if accessId != claims.AccessId {
			return nil, fmt.Errorf("unauthorized")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("not valid token")
}
