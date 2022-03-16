package services

import (
	"errors"
	"strings"

	"gitlab.finema.co/finema/etda/mobile-app-api/views"

	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
	"gorm.io/gorm"
)

type UserCreatePayload struct {
	IDCardNo  string `json:"id_card_no"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type UserUpdatePayload struct {
	ID         string               `json:"id"`
	DIDAddress string               `json:"did_address"`
	NewDevice  *DeviceCreatePayload `json:"new_device"`
}

type IUserService interface {
	Create(payload *UserCreatePayload) (*models.User, core.IError)
	UpdateDID(payload *UserUpdatePayload) (*models.User, core.IError)
	FindByID(id string) (*models.User, core.IError)
	FindByDID(did string) (*models.User, core.IError)
	FindByIDCardNumber(idCardNumber string) (*models.User, core.IError)
	Pagination(ids string, ignoreIds string, pageOptions *core.PageOptions) ([]views.UserList, *core.PageResponse, core.IError)
}

type userService struct {
	ctx core.IContext
}

func NewUserService(ctx core.IContext) IUserService {
	return &userService{
		ctx: ctx,
	}

}

func (s userService) Pagination(ids string, ignoredIds string, pageOptions *core.PageOptions) ([]views.UserList, *core.PageResponse, core.IError) {
	items := make([]models.User, 0)
	if len(pageOptions.OrderBy) <= 0 {
		pageOptions.OrderBy = []string{"created_at desc"}

	}
	db := s.ctx.DB()
	if ignoredIds != "" {
		db = s.ctx.DB().Where("id NOT IN ?", strings.Split(ignoredIds, ","))
	}
	if ids != "" {
		db = s.ctx.DB().Where("id IN ?", strings.Split(ids, ","))
	}
	db = core.SetSearchSimple(db, pageOptions.Q, []string{"email", "first_name", "last_name"})
	pageRes, err := core.Paginate(db, &items, pageOptions)
	if err != nil {
		return nil, nil, s.ctx.NewError(err, errmsgs.DBError)
	}
	viewItems := views.NewUserListView(items)
	return viewItems, pageRes, nil
}
func (s userService) Create(payload *UserCreatePayload) (*models.User, core.IError) {
	user, ierr := s.FindByIDCardNumber(payload.IDCardNo)
	if ierr != nil && !errmsgs.IsNotFoundError(ierr) {
		return nil, s.ctx.NewError(ierr, ierr)
	}
	if user != nil {
		return nil, s.ctx.NewError(emsgs.DuplicatedUser, emsgs.DuplicatedUser)
	}

	uid := utils.GetUUID()
	err := s.ctx.DB().Create(models.User{
		ID:        uid,
		IDCardNo:  payload.IDCardNo,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}).Error

	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return s.FindByID(uid)
}
func (s userService) UpdateDID(payload *UserUpdatePayload) (*models.User, core.IError) {

	err := s.ctx.DB().Updates(&models.User{
		ID:         payload.ID,
		DIDAddress: &payload.DIDAddress,
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.DeviceNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}
	if payload.NewDevice != nil {
		err = s.ctx.DB().Model(models.Device{}).
			Where("user_id = ?", payload.ID).
			Updates(models.Device{DeletedAt: utils.GetCurrentDateTime()}).
			Error

		if err != nil {
			return nil, s.ctx.NewError(err, errmsgs.DBError)
		}
		err = s.ctx.DB().Create(&models.Device{
			ID:        payload.NewDevice.UUID,
			Name:      payload.NewDevice.Name,
			OS:        payload.NewDevice.OS,
			OSVersion: payload.NewDevice.OSVersion,
			Model:     payload.NewDevice.Model,
			UUID:      payload.NewDevice.UUID,
			UserID:    payload.ID,
			CreatedAt: utils.GetCurrentDateTime(),
			UpdatedAt: utils.GetCurrentDateTime(),
		}).Error
		if err != nil {
			return nil, s.ctx.NewError(err, errmsgs.DBError)
		}
	}

	return s.FindByID(payload.ID)
}
func (s userService) FindByID(id string) (*models.User, core.IError) {
	user := &models.User{}
	err := s.ctx.DB().Where("id = ?", id).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.UserNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return user, nil
}

func (s userService) FindByDID(did string) (*models.User, core.IError) {
	user := &models.User{}
	err := s.ctx.DB().Where("did_address = ?", did).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.UserNotFound)
	}

	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return user, nil
}

func (s userService) FindByIDCardNumber(idCardNumber string) (*models.User, core.IError) {
	user := &models.User{}
	err := s.ctx.DB().Where("id_card_no = ?", idCardNumber).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.UserNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return user, nil
}
