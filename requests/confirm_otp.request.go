package requests

import (
	core "ssi-gitlab.teda.th/ssi/core"
)

type ConfirmOTP struct {
	core.BaseValidator
	OTPNumber *string `json:"otp_number"`
}

func (r ConfirmOTP) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.OTPNumber, "otp_number"))
	return r.Error()
}
