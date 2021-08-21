package common

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type BcryptMine interface {
	Hash(password string) (string, error)
	Compare(password string, hashedPassword string) bool
}

type bcyrptMine struct {
	cost   int
	secret string
}

func NewBcrypt() BcryptMine {
	return &bcyrptMine{
		cost:   bcrypt.MinCost,
		secret: getSecret(),
	}
}

func (b *bcyrptMine) Hash(password string) (string, error) {
	bytesHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	return string(bytesHashedPassword), err
}

func (b *bcyrptMine) Compare(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func getSecret() string {
	secret := os.Getenv("HASH_SECRET")
	if secret == "" {
		secret = "default_secret"
	}
	return secret
}
