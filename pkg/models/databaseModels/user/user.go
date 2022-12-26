package user

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type TransformedUser struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     uint   `json:"role"`
}

type Claims struct {
	jwt.StandardClaims
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     uint   `json:"role"`
}
