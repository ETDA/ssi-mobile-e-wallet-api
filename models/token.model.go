package models

import (
	"time"

	"ssi-gitlab.teda.th/ssi/core/utils"
)

type Token struct {
	ID        string     `json:"id" gorm:"column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	Token     string     `json:"token" gorm:"column:token"`
	Role      string     `json:"role" gorm:"column:role"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (m Token) TableName() string {
	return "tokens"
}

func NewToken() *Token {
	currentTime := utils.GetCurrentDateTime()
	id := utils.GetUUID()
	return &Token{
		ID:        id,
		Token:     utils.NewSha256(id + currentTime.String()),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}
