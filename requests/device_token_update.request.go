package requests

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type DeviceTokenUpdate struct {
	core.BaseValidator
	UUID  *string `json:"uuid"`
	Token *string `json:"token"`
}

func (r DeviceTokenUpdate) Valid(ctx core.IContext) core.IError {

	r.Must(r.IsStrRequired(r.UUID, "uuid"))
	r.Must(r.IsExistsWithCondition(ctx, models.Device{}.TableName(), core.Map{
		"uuid": utils.GetString(r.UUID),
	}, "uuid"))

	r.Must(r.IsStrRequired(r.Token, "token"))
	return r.Error()
}
