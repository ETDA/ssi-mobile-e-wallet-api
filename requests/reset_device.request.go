package requests

import (
	core "ssi-gitlab.teda.th/ssi/core"
)

type RequestNewKey struct {
	PublicKey  *string `json:"public_key"`
	Signature  *string `json:"signature"`
	Controller *string `json:"controller"`
}
type ResetDevice struct {
	core.BaseValidator
	NewKey *RequestNewKey `json:"new_key"`
	Device *DeviceDetail  `json:"device"`
}

func (r ResetDevice) Valid(ctx core.IContext) core.IError {
	if r.Must(r.IsRequired(r.NewKey, "new_key")) {
		r.Must(r.IsStrRequired(r.NewKey.PublicKey, "new_key.public_key"))
		r.Must(r.IsStrRequired(r.NewKey.Signature, "new_key.signature"))
	}
	if r.Must(r.IsRequired(r.Device, "device")) {
		r.Must(r.IsStrRequired(r.Device.Name, "device.name"))
		r.Must(r.IsStrRequired(r.Device.OS, "device.os"))
		r.Must(r.IsStrRequired(r.Device.OSVersion, "device.os_version"))
		r.Must(r.IsStrRequired(r.Device.Model, "device.model"))
		r.Must(r.IsStrRequired(r.Device.UUID, "device.uuid"))
	}
	// cliSvc := services.NewClientService(ctx)
	// session, ierr := cliSvc.FindUIDByReferenceID(utils.GetString(r.ReferenceID))
	// if errmsgs.IsNotFoundError(ierr) {
	// 	r.Must(false, core.ExistsM("reference_id"))
	// } else {
	// 	if ierr != nil {
	// 		return ctx.NewError(ierr, ierr)
	// 	}
	// }
	//
	// c, _ := ctx.(core.IHTTPContext)
	// c.Set("uid", session.UID)
	// ok, _ := r.IsExists(ctx,
	// 	&session.UID,
	// 	models.User{}.TableName(),
	// 	"uid",
	// 	"uid")
	// if !ok {
	// 	return ctx.NewError(emsgs.UIDDuplicate, emsgs.UIDDuplicate)
	// }
	// }

	return r.Error()
}
