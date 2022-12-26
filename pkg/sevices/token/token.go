package token

import (
	"coursework/pkg/models/databaseModels/user"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var signingKey string = "grkjk#4#35FSFJLja#4353KSFjH"

func GenerateToken(person user.TransformedUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &user.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       person.Id,
		Username: person.Username,
		Role:     person.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (user.TransformedUser, error) {
	token, err := jwt.ParseWithClaims(accessToken, &user.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	var person user.TransformedUser

	if err != nil {
		return person, err
	}

	if claims, ok := token.Claims.(*user.Claims); ok && token.Valid {
		person = user.TransformedUser{
			Id:       claims.Id,
			Username: claims.Username,
			Role:     claims.Role,
		}
		return person, nil
	}

	fmt.Println(person)

	return person, errors.New("invalid auth token")
}
