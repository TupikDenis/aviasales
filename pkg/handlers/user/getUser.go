package user

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/user"
	"coursework/pkg/sevices/token"
	"crypto/sha256"
	"encoding/binary"
	"strconv"
)

func Login(username string, userPassword string) string {
	hash := sha256.New()
	hash.Write([]byte(userPassword))

	var persons []user.User
	var person user.TransformedUser

	db := handlers.Database()
	db.Where("username = ?", username).
		Where("password = ?", strconv.FormatInt(int64(binary.BigEndian.Uint64(hash.Sum(nil))), 16)).
		Find(&persons)

	person = user.TransformedUser{
		Id:       persons[0].ID,
		Username: persons[0].Username,
		Role:     persons[0].Role,
	}

	tokenStr, err := token.GenerateToken(person)

	if err != nil {
		panic(err)
	}

	return tokenStr
}
