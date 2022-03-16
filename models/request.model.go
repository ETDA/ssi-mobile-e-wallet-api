package models

import (
	"encoding/json"
	"ssi-gitlab.teda.th/ssi/core/utils"
	"time"
)

type Request struct {
	ID             string           `json:"id" gorm:"column:id"`
	RequestID      string           `json:"request_id" gorm:"column:request_id"`
	RequestData    *json.RawMessage `json:"request_data" gorm:"column:request_data"`
	SchemaType     string           `json:"schema_type" gorm:"column:schema_type"`
	CredentialType string           `json:"credential_type" gorm:"column:credential_type"`
	Signer         string           `json:"signer" gorm:"column:signer"`
	Requester      string           `json:"requester" gorm:"column:requester"`
	Status         string           `json:"status" gorm:"column:status"`
	CreatedAt      *time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time       `json:"updated_at" gorm:"column:updated_at"`
}

func (r Request) TableName() string {
	return "requests"
}

func NewRequestID(requestID string) string {
	return utils.NewSha256(requestID + utils.GetCurrentDateTime().String())
}
