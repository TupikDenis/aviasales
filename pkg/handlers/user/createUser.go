package user

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/user"
	"crypto/sha256"
	"encoding/binary"
	"gorm.io/gorm"
	"strconv"
)

func CreateUser(username string, userPassword string) bool {
	hash := sha256.New()
	hash.Write([]byte(userPassword))

	person := user.User{
		Model:    gorm.Model{},
		Username: username,
		Password: strconv.FormatInt(int64(binary.BigEndian.Uint64(hash.Sum(nil))), 16),
		Role:     0,
	}

	db := handlers.Database()

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		return false
	}

	db.Save(&person)

	return true
}
