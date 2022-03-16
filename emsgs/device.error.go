package emsgs

import (
	core "ssi-gitlab.teda.th/ssi/core"
	"net/http"
)

var DeviceNotFound = core.Error{
	Status:  http.StatusNotFound,
	Code:    "DEVICE_NOT_FOUND",
	Message: "device is not found",
}

var DeviceUUIDNotFound = core.Error{
	Status:  http.StatusNotFound,
	Code:    "UUID_NOT_FOUND",
	Message: "your did_address is not own this device",
}
