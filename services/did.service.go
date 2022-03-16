package services

import (
	"fmt"
	"net/http"

	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/helpers"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type IDIDService interface {
	Find(didAddress string) (*models.DIDDocument, core.IError)
	RetrieveDIDDocument(did string) (*DidDocumentResponse, core.IError)
	ResetterApproveReset(payload *RecoverResetPayload) (bool, core.IError)
	RegisterDID(publicKey string, privateKey string) (*DidDocumentResponse, core.IError)
	GetNonce(did string) (string, core.IError)
}

type didService struct {
	ctx core.IContext
}

func NewDIDService(ctx core.IContext) IDIDService {
	return &didService{ctx: ctx}
}

type DidDocumentResponse struct {
	Context            string                 `json:"@context"`
	ID                 string                 `json:"id"`
	VerificationMethod []DIDDocumentPublicKey `json:"verificationMethod"`
}

type DIDDocumentPublicKey struct {
	ID           string  `json:"id"`
	Type         string  `json:"type"`
	Controller   string  `json:"controller"`
	PublicKeyPem *string `json:"publicKeyPem"`
}

func (s didService) RetrieveDIDDocument(did string) (*DidDocumentResponse, core.IError) {
	res, err := s.ctx.Requester().Get("/did/"+did, &core.RequesterOptions{
		BaseURL: s.ctx.ENV().String(consts.ENVDIDServiceBaseURL),
	})

	if errmsgs.IsNotFoundErrorCode(res.ErrorCode) {
		return nil, s.ctx.NewError(err, emsgs.DIDNotFound)
	}

	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	resBody := &DidDocumentResponse{}
	_ = utils.JSONParse(res.RawData, resBody)
	return resBody, nil
}

type ResetterApproveRegisterPayload struct {
	DidAddress         string
	ResetterDIDAddress string
	CurrentKey         string
	PrivateKey         string
	NextKeyHash        string
}

func (s didService) RegisterDID(publicKey string, privateKey string) (*DidDocumentResponse, core.IError) {
	data := core.Map{
		"public_key": publicKey,
		"key_type":   "EcdsaSecp256r1VerificationKey2019",
		"operation":  "DID_REGISTER",
	}

	body, headers, ierr := s.NewMessage(data, privateKey)
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr, body)
	}

	res, err := s.ctx.Requester().Post("/did", body, &core.RequesterOptions{
		BaseURL: s.ctx.ENV().String(consts.ENVDIDServiceBaseURL),
		Headers: headers,
	})

	if errmsgs.IsNotFoundErrorCode(res.ErrorCode) {
		return nil, s.ctx.NewError(err, emsgs.DIDNotFound)
	}

	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.InternalServerError, body)
	}

	resBody := &DidDocumentResponse{}
	_ = utils.JSONParse(res.RawData, resBody)
	return resBody, nil
}

type NewKeyPayload struct {
	PublicKey  string  `json:"public_key"`
	Signature  string  `json:"signature"`
	Controller *string `json:"controller"`
}
type RecoverResetPayload struct {
	DIDAddress string
	RequestDID string
	PrivateKey string
	NewKey     *NewKeyPayload
}

func (s didService) GetNonce(did string) (string, core.IError) {
	res, err := s.ctx.Requester().Get(fmt.Sprintf("/did/%s/nonce", did), &core.RequesterOptions{
		BaseURL: s.ctx.ENV().String(consts.ENVDIDServiceBaseURL),
	})

	if err != nil {
		return "", s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	nonce := res.Data["nonce"].(string)
	return nonce, nil
}

func (s didService) NewMessage(data core.Map, privateKeyPEM string) (body core.Map, headers http.Header, error core.IError) {
	dataBase64 := utils.Base64Encode(utils.JSONToString(data))

	body = core.Map{
		"message": dataBase64,
	}

	pri, err := utils.LoadPrivateKey(privateKeyPEM)
	if err != nil {
		return nil, nil, s.ctx.NewError(err, errmsgs.InternalServerError, privateKeyPEM)
	}

	signature, err := utils.SignMessage(pri, dataBase64)
	if err != nil {
		return nil, nil, s.ctx.NewError(err, errmsgs.InternalServerError, dataBase64, privateKeyPEM)
	}

	headers = http.Header{}
	headers.Set("x-signature", signature)

	return body, headers, nil
}
func (s didService) ResetterApproveReset(payload *RecoverResetPayload) (bool, core.IError) {
	nonce, ierr := s.GetNonce(payload.DIDAddress)
	fmt.Println("nonce")
	if ierr != nil {
		return false, s.ctx.NewError(ierr, ierr)
	}

	data := core.Map{
		"operation":   "DID_KEY_RESET",
		"did_address": payload.DIDAddress,
		"request_did": payload.RequestDID,
		"nonce":       nonce,
		"new_key":     payload.NewKey,
	}

	body, headers, ierr := s.NewMessage(data, payload.PrivateKey)
	if ierr != nil {
		return false, s.ctx.NewError(ierr, ierr, body)
	}

	res, err := s.ctx.Requester().Post(fmt.Sprintf("/did/%s/keys/reset", payload.DIDAddress), body, &core.RequesterOptions{
		BaseURL: s.ctx.ENV().String(consts.ENVDIDServiceBaseURL),
		Headers: headers,
	})
	fmt.Println("Fck")
	if errmsgs.IsNotFoundErrorCode(res.ErrorCode) {
		return false, s.ctx.NewError(err, emsgs.DIDNotFound)
	}

	if err != nil {
		return false, s.ctx.NewError(err, errmsgs.InternalServerError, body)
	}

	return true, nil
}
func (s *didService) Find(didAddress string) (*models.DIDDocument, core.IError) {
	res, err := s.ctx.Requester().Get(fmt.Sprintf("/did/%s/document/latest", didAddress),
		&core.RequesterOptions{
			BaseURL: s.ctx.ENV().String(consts.ENVDIDServiceBaseURL),
		})
	if res == nil {
		return nil, s.ctx.NewError(errmsgs.InternalServerError, errmsgs.InternalServerError)
	}

	if err != nil {
		ierr := helpers.HTTPErrorToIError(res)
		return nil, s.ctx.NewError(ierr, ierr)
	}

	result := &models.DIDDocument{}
	err = utils.MapToStruct(res.Data, result)
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	return result, nil
}
