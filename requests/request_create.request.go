package requests

import (
	"encoding/json"
	"fmt"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	core "ssi-gitlab.teda.th/ssi/core"
)

type RequestCreate struct {
	core.BaseValidator
	Operation      *string          `json:"operation"`
	DIDAddress     *string          `json:"did_address"`
	RequestID      *string          `json:"request_id"`
	RequestData    *json.RawMessage `json:"request_data"`
	SchemaType     *string          `json:"schema_type"`
	CredentialType *string          `json:"credential_type"`
	Signer         *string          `json:"signer"`
	Requester      *string          `json:"requester"`
}

func (r RequestCreate) Valid(ctx core.IContext) core.IError {
	c := ctx.(core.IHTTPContext)

	r.Must(r.IsStrIn(r.Operation, consts.OperationRequestCreate, "operation"))
	r.Must(r.IsStrRequired(r.DIDAddress, "did_address"))

	r.Must(r.IsJSONRequired(r.RequestData, "request_data"))
	r.Must(r.IsJSONObject(r.RequestData, "request_data"))
	r.Must(r.IsJSONObjectNotEmpty(r.RequestData, "request_data"))

	r.Must(r.IsStrRequired(r.RequestID, "request_id"))
	r.Must(r.IsStrIn(r.RequestID, c.Param("request_id"), "request_id"))

	r.Must(r.IsStrRequired(r.CredentialType, "credential_type"))
	r.Must(r.IsStrIn(r.CredentialType, fmt.Sprintf("%s|%s", consts.CredentialTypeVP, consts.CredentialTypeVC), "credential_type"))

	r.Must(r.IsStrRequired(r.SchemaType, "schema_type"))
	r.Must(r.IsStrRequired(r.Signer, "signer"))
	r.Must(r.IsStrRequired(r.Requester, "requester"))
	return r.Error()
}
