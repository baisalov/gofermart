package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrAlreadyExist = errors.New("user already exist")
)

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

type Claims struct {
	jwt.RegisteredClaims
	User
}
