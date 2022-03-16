package models

import "time"

type User struct {
	ID         string     `json:"id" gorm:"column:id"`
	IDCardNo   string     `json:"id_card_no" gorm:"column:id_card_no"`
	FirstName  string     `json:"first_name" gorm:"column:first_name"`
	LastName   string     `json:"last_name" gorm:"column:last_name"`
	DIDAddress *string    `json:"did_address" gorm:"column:did_address"`
	Email      string     `json:"email" gorm:"column:email"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (r User) TableName() string {
	return "users"
}
