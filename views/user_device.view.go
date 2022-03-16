package views

import (
	"time"

	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type UserDeviceDevice struct {
	Name      string `json:"name"`
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
	Model     string `json:"model"`
	UUID      string `json:"uuid"`
}

type UserDevice struct {
	ID             string            `json:"id"`
	DIDAddress     string            `json:"did_address"`
	FirstName      string            `json:"first_name"`
	LastName       string            `json:"last_name"`
	Email          string            `json:"email"`
	Device         *UserDeviceDevice `json:"device"`
	RegisteredDate *time.Time        `json:"registered_date"`
}

func NewUserDevice(user *models.User, device *models.Device) *UserDevice {
	return &UserDevice{
		ID:         utils.GetUUID(),
		DIDAddress: utils.GetString(user.DIDAddress),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Device: &UserDeviceDevice{
			Name:      device.Name,
			OS:        device.OS,
			OSVersion: device.OSVersion,
			Model:     device.Model,
			UUID:      device.UUID,
		},
		RegisteredDate: user.CreatedAt,
	}
}
