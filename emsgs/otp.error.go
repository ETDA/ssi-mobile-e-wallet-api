package emsgs

import (
	"net/http"

	core "ssi-gitlab.teda.th/ssi/core"
)

var OTPNotValid = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "OTP_NOT_VALID",
	Message: "OTP is not valid.",
}

var EmailError = core.Error{
	Status:  http.StatusInternalServerError,
	Code:    "EMAIL_ERROR",
	Message: "email server internal error"}
