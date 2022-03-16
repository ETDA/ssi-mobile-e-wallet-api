package emsgs

import (
	"net/http"

	core "ssi-gitlab.teda.th/ssi/core"
)

var DuplicatedDIDAddress = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "DUPLICATED_DID_ADDRESS",
	Message: "did address already used",
}
var DIDNotFound = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "DID_NOT_FOUND",
	Message: "did not found"}
