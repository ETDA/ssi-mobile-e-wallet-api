package services

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type SendTokenNotificationItem struct {
	NotificationItem
	Token string `json:"token"`
}

type SendTopicNotificationItem struct {
	NotificationItem
	Topic string `json:"topic"`
}

type NotificationItem struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageURL string `json:"image_url"`

	Category    string            `json:"category"`
	Icon        string            `json:"icon"`
	ClickAction string            `json:"click_action"`
	Sound       string            `json:"sound"`
	Priority    string            `json:"priority"` // one of "normal" or "high"
	Data        map[string]string `json:"data"`
}

type NotificationTokenItem struct {
	NotificationItem
	Token string `json:"token"`
}

type INotificationService interface {
	SendByTokens(tokenNotificationItems []SendTokenNotificationItem) core.IError
	SendByDID(didAddress string, notificationItems []NotificationItem) core.IError
}

type notificationService struct {
	ctx        core.IContext
	didService IDeviceService
}

func NewNotificationService(ctx core.IContext) INotificationService {
	return &notificationService{
		ctx:        ctx,
		didService: NewDeviceService(ctx),
	}
}

func (s notificationService) SendByDID(didAddress string, notificationItems []NotificationItem) core.IError {
	// User cannot find device by DID address because Device have only relation with user not a did_address
	device, ierr := s.didService.FindByDID(didAddress)
	if ierr != nil {
		return s.ctx.NewError(ierr, ierr)
	}

	sendItems := make([]SendTokenNotificationItem, 0)
	for _, item := range notificationItems {
		sendItems = append(sendItems, SendTokenNotificationItem{
			NotificationItem: item,
			Token:            device.Token,
		})
	}

	return s.SendByTokens(sendItems)
}

func (s notificationService) SendByTokens(payload []SendTokenNotificationItem) core.IError {

	notifications := make([]NotificationTokenItem, 0)
	err := utils.Copy(&notifications, payload)

	if err != nil {
		return s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	app, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: s.ctx.ENV().String(consts.ENVFirebaseProjectID)})
	if err != nil {
		return s.ctx.NewError(err, emsgs.NotificationError(err.Error()))
	}

	msg, err := app.Messaging(context.Background())
	if err != nil {
		return s.ctx.NewError(err, emsgs.NotificationError(err.Error()))
	}
	messages := make([]*messaging.Message, 0)
	for _, notification := range payload {
		messages = append(messages, &messaging.Message{
			Data:  notification.Data,
			Token: notification.Token,
			Android: &messaging.AndroidConfig{
				Priority:     notification.Priority,
				Notification: nil,
			},
			APNS: &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						Alert: &messaging.ApsAlert{
							Title:       notification.Title,
							Body:        notification.Body,
							LaunchImage: notification.ImageURL,
						},
						Sound:            notification.Sound,
						Category:         notification.Category,
						ContentAvailable: true,
					},
				},
			},
		})
	}
	_, err = msg.SendAll(context.Background(), messages)
	if err != nil {
		return s.ctx.NewError(err, emsgs.NotificationError(err.Error()))
	}

	return nil
}
