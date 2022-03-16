package requests

import (
	"fmt"
	core "ssi-gitlab.teda.th/ssi/core"
)

type NotificationDID struct {
	core.BaseValidator
	DIDAddress    *string            `json:"did_address"`
	Notifications []NotificationItem `json:"notifications"`
}

func (r *NotificationDID) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.DIDAddress, "did_address"))
	r.Must(r.IsRequiredArray(r.Notifications, "notifications"))
	for i, item := range r.Notifications {
		item.SetPrefix(fmt.Sprintf("notifications[%v].", i))
		item.Valid(ctx)
		r.AddValidator(item.GetValidator())
	}
	return r.Error()
}
