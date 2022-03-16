package requests

import core "ssi-gitlab.teda.th/ssi/core"

type EKYCVerify struct {
	core.BaseValidator
	IDCardNo    *string `json:"id_card_no"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	DateOfBirth *string `json:"date_of_birth"`
	LaserID     *string `json:"laser_id"`
}

func (r EKYCVerify) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.IDCardNo, "id_card_no"))
	r.Must(r.IsStrRequired(r.FirstName, "first_name"))
	r.Must(r.IsStrRequired(r.LastName, "last_name"))
	r.Must(r.IsStrRequired(r.DateOfBirth, "date_of_birth"))
	r.Must(r.IsStrRequired(r.LaserID, "laser_id"))
	return r.Error()
}
