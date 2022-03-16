package requests

import core "ssi-gitlab.teda.th/ssi/core"

type UserRegisterConfirmation struct {
	core.BaseValidator
	IDCardNo *string `json:"id_card_no"`
	OTPCode  *string `json:"otp_code"`
}

func (r UserRegisterConfirmation) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.IDCardNo, "id_card_no"))
	r.Must(r.IsStrRequired(r.OTPCode, "otp_code"))
	return r.Error()
}
