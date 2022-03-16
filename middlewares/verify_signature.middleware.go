package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/services"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type payload struct {
	core.BaseValidator
	Message *string `json:"message"`
}

func (r payload) Valid(ctx core.IContext) core.IError {
	if r.Must(r.IsStrRequired(r.Message, "message")) {
		r.Must(r.IsBase64(r.Message, "message"))
	}

	return r.Error()
}

type OperationPayload struct {
	DIDAddress string `json:"did_address"`
	Operation  string `json:"operation"`
}

func VerifySignatureMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(core.IHTTPContext)
		if cc.GetSignature() == "" {
			return c.JSON(http.StatusBadRequest, core.NewValidatorFields(core.RequiredM("x-signature")))
		}

		var isSigValid = false
		var didAddress string
		var operation string
		var message string
		if cc.Request().Method == http.MethodPost || cc.Request().Method == http.MethodPut {
			payloadData := &payload{}
			if err := cc.BindWithValidate(payloadData); err != nil {
				return c.JSON(err.GetStatus(), err.JSON())
			}
			message = utils.GetString(payloadData.Message)
			jsonString, err := utils.Base64Decode(message) // decode failed
			if err != nil {
				return c.JSON(errmsgs.BadRequest.GetStatus(), errmsgs.BadRequest.JSON())
			}
			messagePayload := &OperationPayload{}
			err = utils.JSONParse([]byte(jsonString), messagePayload) // unmarshall failed
			if err != nil {
				return c.JSON(errmsgs.BadRequest.GetStatus(), errmsgs.BadRequest.JSON())
			}
			didAddress = messagePayload.DIDAddress
			operation = messagePayload.Operation
			if (operation == consts.OperationRequestSign || operation == consts.OperationRequestUpdate || operation == consts.OperationRequestCreate) && cc.Param("did") != didAddress {
				return c.JSON(errmsgs.BadRequest.GetStatus(), errmsgs.BadRequest.JSON())
			}
		}

		if cc.Request().Method == http.MethodGet {
			didAddress = cc.Param("did")
			message = didAddress
		}

		c.Set("message", message)

		didService := services.NewDIDService(cc)
		didDocument, ierr := didService.Find(didAddress)
		if ierr != nil {
			return c.JSON(ierr.GetStatus(), ierr.JSON())
		}
		for _, verificationMethod := range didDocument.VerificationMethod {
			valid, _ := utils.VerifySignature(verificationMethod.PublicKeyPem, cc.GetSignature(), message)
			if valid {
				isSigValid = valid
				break
			}
		}

		if !isSigValid {
			return c.JSON(errmsgs.SignatureInValid.GetStatus(), errmsgs.SignatureInValid.JSON())
		}

		return next(c)
	}
}
