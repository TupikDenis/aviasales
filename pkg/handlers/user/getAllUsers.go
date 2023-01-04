package user

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/user"
)

func GetAllUsers() []user.TransformedUser {
	var users []user.User
	var _users []user.TransformedUser

	db := handlers.Database()
	db.Order("id asc").Find(&users)

	for _, item := range users {
		_users = append(
			_users, user.TransformedUser{
				Id:       item.ID,
				Username: item.Username,
				Role:     item.Role,
			})
	}

	return _users
}
