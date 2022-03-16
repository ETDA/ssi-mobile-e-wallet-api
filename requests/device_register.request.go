package requests

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
)

type DeviceRegister struct {
	core.BaseValidator
	Operation  *string `json:"operation"`
	DIDAddress *string `json:"did_address"`
	ID         *string `json:"id"`
	Device     *struct {
		Name      *string `json:"name"`
		OS        *string `json:"os"`
		OSVersion *string `json:"os_version"`
		Model     *string `json:"model"`
		UUID      *string `json:"uuid"`
	} `json:"device"`
}

func (r DeviceRegister) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrIn(r.Operation, consts.OperationDeviceRegister, "operation"))
	r.Must(r.IsStrRequired(r.DIDAddress, "did_address"))

	r.Must(r.IsStrRequired(r.ID, "id"))
	r.Must(r.IsExists(ctx, r.ID, models.User{}.TableName(), "id", "id"))

	r.Must(r.IsStrRequired(r.Device.Name, "device.name"))
	r.Must(r.IsStrRequired(r.Device.OS, "device.os"))
	r.Must(r.IsStrRequired(r.Device.OSVersion, "device.os_version"))
	r.Must(r.IsStrRequired(r.Device.Model, "device.model"))
	r.Must(r.IsStrRequired(r.Device.UUID, "device.uuid"))
	r.Must(r.IsStrUnique(ctx, r.ID, models.Device{}.TableName(), "uuid", "", "device.uuid"))

	return r.Error()
}
