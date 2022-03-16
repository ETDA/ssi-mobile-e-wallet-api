package services

import (
	"encoding/json"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/helpers"
	"gitlab.finema.co/finema/etda/mobile-app-api/views"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"net/http"
)

type identityProofingIDCardVerifyPayload struct {
	CardID    string `json:"card_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	LaserID   string `json:"laser_id"`
	Birthdate string `json:"birthdate"`
}

type IIdentityProofingService interface {
	IDCardVerify(payload *identityProofingIDCardVerifyPayload) (*views.IdentityProofingIDCardVerify, core.IError)
}

type identityProofingService struct {
	ctx core.IContext
}

func NewIdentityProofingService(ctx core.IContext) IIdentityProofingService {
	return &identityProofingService{ctx: ctx}
}

func (s identityProofingService) IDCardVerify(payload *identityProofingIDCardVerifyPayload) (*views.IdentityProofingIDCardVerify, core.IError) {
	newBirthdate, err := helpers.BirthdateTransform(payload.Birthdate)
	if err != nil {
		return &views.IdentityProofingIDCardVerify{
			Status: false,
		}, s.ctx.NewError(err, emsgs.InvalidDateFormat)
	}

	payload.Birthdate = newBirthdate
	res, err := s.ctx.Requester().Post(
		"/id-card/verify",
		payload,
		&core.RequesterOptions{
			BaseURL: s.ctx.ENV().String(consts.ENVIdentityProofingServerBaseURL),
			Headers: http.Header{
				"Authorization": []string{s.ctx.ENV().String(consts.ENVIdentityProofingServerAPIToken)},
			},
		})
	if err != nil {
		return &views.IdentityProofingIDCardVerify{
			Status: false,
		}, s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	result := &views.IdentityProofingIDCardVerify{}
	err = json.Unmarshal(res.RawData, result)
	if err != nil {
		return &views.IdentityProofingIDCardVerify{
			Status: false,
		}, s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	return result, nil
}
