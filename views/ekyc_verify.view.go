package views

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type EKYCVerify struct {
	ID         string `json:"id,omitempty"`
	CardStatus bool   `json:"card_status"`
	Message    string `json:"message,omitempty"`
	DIDAddress string `json:"did_address,omitempty"`
}

func NewEKYCViewSuccess(user *models.User) *EKYCVerify {
	return &EKYCVerify{
		ID:         user.ID,
		CardStatus: true,
		DIDAddress: utils.GetString(user.DIDAddress),
	}
}
