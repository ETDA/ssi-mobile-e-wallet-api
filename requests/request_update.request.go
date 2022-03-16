package requests

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
)

type RequestUpdate struct {
	core.BaseValidator
	Operation  *string `json:"operation"`
	DIDAddress *string `json:"did_address"`
	Status     *string `json:"status"`
	RequestID  *string `json:"request_id"`
}

func (r RequestUpdate) Valid(ctx core.IContext) core.IError {
	c := ctx.(core.IHTTPContext)

	r.Must(r.IsStrIn(r.Operation, consts.OperationRequestUpdate, "operation"))
	r.Must(r.IsStrRequired(r.DIDAddress, "did_address"))

	r.Must(r.IsStrRequired(r.Status, "status"))
	r.Must(r.IsStrIn(r.Status, consts.RequestStatusSigned, "status"))

	r.Must(r.IsStrRequired(r.RequestID, "request_id"))
	r.Must(r.IsStrIn(r.RequestID, c.Param("request_id"), "request_id"))
	r.Must(r.IsExists(ctx, r.RequestID, models.Request{}.TableName(), "request_id", "request_id"))

	return r.Error()
}
