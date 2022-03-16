package requests

import core "ssi-gitlab.teda.th/ssi/core"

type DeviceDetail struct {
	Name      *string `json:"name"`
	OS        *string `json:"os"`
	OSVersion *string `json:"os_version"`
	Model     *string `json:"model"`
	UUID      *string `json:"uuid"`
}
type UserRegister struct {
	core.BaseValidator
	IDCardNo    *string       `json:"id_card_no"`
	FirstName   *string       `json:"first_name"`
	LastName    *string       `json:"last_name"`
	DateOfBirth *string       `json:"date_of_birth"`
	LaserID     *string       `json:"laser_id"`
	Email       *string       `json:"email"`
	Device      *DeviceDetail `json:"device"`
}

func (r UserRegister) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.IDCardNo, "id_card_no"))
	r.Must(r.IsStrRequired(r.FirstName, "first_name"))
	r.Must(r.IsStrRequired(r.LastName, "last_name"))
	r.Must(r.IsStrRequired(r.DateOfBirth, "date_of_birth"))
	r.Must(r.IsStrRequired(r.LaserID, "laser_id"))
	r.Must(r.IsStrRequired(r.Email, "email"))
	r.Must(r.IsEmail(r.Email, "email"))
	if r.Must(r.IsRequired(r.Device, "device")) {
		r.Must(r.IsStrRequired(r.Device.Name, "device.name"))
		r.Must(r.IsStrRequired(r.Device.OS, "device.os"))
		r.Must(r.IsStrRequired(r.Device.OSVersion, "device.os_version"))
		r.Must(r.IsStrRequired(r.Device.Model, "device.model"))
		r.Must(r.IsStrRequired(r.Device.UUID, "device.uuid"))
	}
	return r.Error()
}
