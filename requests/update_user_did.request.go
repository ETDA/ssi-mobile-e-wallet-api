package requests

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
)

type UpdateUserDID struct {
	core.BaseValidator
	ID         *string       `json:"id"`
	DIDAddress *string       `json:"did_address"`
	Device     *DeviceDetail `json:"device"`
}

func (r UpdateUserDID) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.ID, "id"))
	r.Must(r.IsExists(ctx, r.ID, models.User{}.TableName(), "id", "id"))
	r.Must(r.IsStrRequired(r.DIDAddress, "did_address"))
	if r.Device != nil {
		r.Must(r.IsStrRequired(r.Device.Name, "device.name"))
		r.Must(r.IsStrRequired(r.Device.OS, "device.os"))
		r.Must(r.IsStrRequired(r.Device.OSVersion, "device.os_version"))
		r.Must(r.IsStrRequired(r.Device.Model, "device.model"))
		r.Must(r.IsStrRequired(r.Device.UUID, "device.uuid"))
	}

	return r.Error()
}
