package views

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
)

type User struct {
	DIDAddress *string `json:"did_address"`
	UserID     string  `json:"user_id"`
}

func NewUser(user *models.User) *User {
	return &User{
		DIDAddress: user.DIDAddress,
		UserID:     user.ID,
	}
}

type UserList struct {
	ID         string  `json:"id" gorm:"column:id"`
	FirstName  string  `json:"first_name" gorm:"column:first_name"`
	LastName   string  `json:"last_name" gorm:"column:last_name"`
	DIDAddress *string `json:"did_address" gorm:"column:did_address"`
}

func NewUserListView(users []models.User) []UserList {
	userList := make([]UserList, 0)
	for _, user := range users {
		userView := &UserList{
			ID:         user.ID,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			DIDAddress: user.DIDAddress,
		}
		userList = append(userList, *userView)
	}
	return userList
}
