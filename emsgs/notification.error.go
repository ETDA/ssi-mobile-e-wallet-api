package emsgs

import (
	core "ssi-gitlab.teda.th/ssi/core"
	"net/http"
)

var NotificationError = func(msg string) core.IError {
	return core.Error{
		Status:  http.StatusBadRequest,
		Code:    "NOTIFICATION_ERROR",
		Message: msg}
}
