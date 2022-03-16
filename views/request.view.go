package views

import (
	"encoding/json"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	"ssi-gitlab.teda.th/ssi/core/utils"
	"time"
)

type Request struct {
	RequestID      string           `json:"request_id"`
	RequestData    *json.RawMessage `json:"request_data"`
	SchemaType     string           `json:"schema_type"`
	CredentialType string           `json:"credential_type"`
	Signer         string           `json:"signer"`
	Requester      string           `json:"requester"`
	Status         string           `json:"status"`
	RequestDate    *time.Time       `json:"request_date"`
}

func NewRequest(request *models.Request) *Request {
	view := &Request{}
	_ = utils.Copy(view, request)

	view.RequestDate = request.CreatedAt
	return view
}

func NewRequests(requests []models.Request) []Request {
	views := make([]Request, 0)

	for _, req := range requests {
		views = append(views, *NewRequest(&req))
	}

	return views
}
