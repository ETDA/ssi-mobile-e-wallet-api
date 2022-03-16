package emsgs

import (
	core "ssi-gitlab.teda.th/ssi/core"
	"net/http"
)

var UserNotFound = core.Error{
	Status:  http.StatusNotFound,
	Code:    "USER_NOT_FOUND",
	Message: "user is not found",
}

var DuplicatedUser = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "DUPLICATED_USER",
	Message: "user is already exists",
}
