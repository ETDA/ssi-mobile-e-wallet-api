package requests

import (
	core "ssi-gitlab.teda.th/ssi/core"
)

type NotificationItem struct {
	core.BaseValidator
	Title    *string `json:"title"`
	Body     *string `json:"body"`
	ImageURL *string `json:"image_url"`

	Category    *string           `json:"category"`
	Icon        *string           `json:"icon"`
	ClickAction *string           `json:"click_action"`
	Sound       *string           `json:"sound"`
	Priority    *string           `json:"priority"` // one of "normal" or "high"
	Data        map[string]string `json:"data"`
}

func (r *NotificationItem) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.Title, "title"))
	r.Must(r.IsStrRequired(r.Body, "body"))
	r.Must(r.IsURL(r.ImageURL, "image_url"))
	r.Must(r.IsStrIn(r.Priority, "normal|high", "priority"))
	return r.Error()
}
