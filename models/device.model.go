package models

import "time"

type Device struct {
	ID        string     `json:"id" gorm:"column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	OS        string     `json:"os" gorm:"column:os"`
	OSVersion string     `json:"os_version" gorm:"column:os_version"`
	Model     string     `json:"model" gorm:"column:model"`
	UUID      string     `json:"uuid" gorm:"column:uuid"`
	Token     string     `json:"token" gorm:"column:token"`
	UserID    string     `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (r Device) TableName() string {
	return "devices"
}
