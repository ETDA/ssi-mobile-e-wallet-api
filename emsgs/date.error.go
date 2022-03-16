package emsgs

import (
	core "ssi-gitlab.teda.th/ssi/core"
	"net/http"
)

var InvalidDateFormat = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "INVALID_DATE_FORMAT",
	Message: "Please input date in format \"dd m.m. yyyy\"",
}
