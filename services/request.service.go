package services

import (
	"encoding/json"
	"errors"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
	"gorm.io/gorm"
)

type RequestPaginationOptions struct {
	Status         string
	CredentialType string
}

type RequestUpdatePayload struct {
	Status string
}

type RequestCreatePayload struct {
	RequestID      string           `json:"request_id"`
	RequestData    *json.RawMessage `json:"request_data"`
	SchemaType     string           `json:"schema_type"`
	CredentialType string           `json:"credential_type"`
	Signer         string           `json:"signer"`
	Requester      string           `json:"requester"`
}
type RequestCountOptions struct {
	RequestPaginationOptions
}

type IRequestService interface {
	Create(payload *RequestCreatePayload) (*models.Request, core.IError)
	Find(id string) (*models.Request, core.IError)
	Update(id string, payload *RequestUpdatePayload) (*models.Request, core.IError)
	PaginationByDID(did string, pageOptions *core.PageOptions, options *RequestPaginationOptions) ([]models.Request, *core.PageResponse, core.IError)
	CountByDID(did string, options *RequestCountOptions) (int64, core.IError)
}
type requestService struct {
	ctx                 core.IContext
	notificationService INotificationService
	deviceService       IDeviceService
}

func NewRequestService(ctx core.IContext) IRequestService {
	return &requestService{
		ctx:                 ctx,
		notificationService: NewNotificationService(ctx),
		deviceService:       NewDeviceService(ctx),
	}
}

func (s requestService) Create(payload *RequestCreatePayload) (*models.Request, core.IError) {
	id := models.NewRequestID(payload.RequestID)
	newRequest := &models.Request{
		ID:             id,
		RequestID:      payload.RequestID,
		RequestData:    payload.RequestData,
		SchemaType:     payload.SchemaType,
		CredentialType: payload.CredentialType,
		Signer:         payload.Signer,
		Requester:      payload.Requester,
		Status:         consts.RequestStatusUnsigned,
		CreatedAt:      utils.GetCurrentDateTime(),
		UpdatedAt:      utils.GetCurrentDateTime(),
	}
	err := s.ctx.DB().Create(newRequest).Error
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	signer, ierr := s.deviceService.FindByUserID(newRequest.Signer)
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}

	ierr = s.notificationService.SendByTokens([]SendTokenNotificationItem{{
		// TODO: create notification item
		NotificationItem: NotificationItem{
			Title:       "",
			Body:        "",
			ImageURL:    "",
			Category:    "",
			Icon:        "",
			ClickAction: "",
			Sound:       "",
			Priority:    "",
			Data:        nil,
		},
		Token: signer.Token,
	}})

	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}

	return s.Find(id)
}

func (s requestService) Find(id string) (*models.Request, core.IError) {
	req := &models.Request{}
	err := s.ctx.DB().First(req, "id = ? OR request_id = ?", id, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, errmsgs.NotFound)
	}

	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return req, nil
}

func (s requestService) Update(id string, payload *RequestUpdatePayload) (*models.Request, core.IError) {
	_, ierr := s.Find(id)
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}

	err := s.ctx.DB().Model(models.Request{}).Where("id = ?", id).Update("status", payload.Status).Error
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return s.Find(id)
}

func (s requestService) PaginationByDID(did string, pageOptions *core.PageOptions, options *RequestPaginationOptions) ([]models.Request, *core.PageResponse, core.IError) {
	items := make([]models.Request, 0)
	db := s.ctx.DB().Where("signer = ?", did)
	if options != nil {
		if options.Status != "" {
			db = s.ctx.DB().Where("status = ?", options.Status)
		}

		if options.CredentialType != "" {
			db = s.ctx.DB().Where("credential_type = ?", options.CredentialType)
		}
	}

	pageRes, err := core.Paginate(db, &items, pageOptions)
	if err != nil {
		return nil, nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return items, pageRes, nil
}

func (s requestService) CountByDID(did string, options *RequestCountOptions) (int64, core.IError) {
	var count int64
	db := s.ctx.DB().Where("signer = ?", did)
	if options != nil {
		if options.Status != "" {
			db = s.ctx.DB().Where("status = ?", options.Status)
		}

		if options.CredentialType != "" {
			db = s.ctx.DB().Where("credential_type = ?", options.CredentialType)
		}
	}

	err := db.Count(&count).Error
	if err != nil {
		return 0, s.ctx.NewError(err, errmsgs.DBError)
	}

	return count, nil
}
