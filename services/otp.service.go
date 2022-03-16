package services

import (
	"errors"
	"fmt"
	"math/rand"

	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/helpers"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
	"gorm.io/gorm"
)

type OTPCreatePayload struct {
	UserID             string `json:"user_id"`
	IsAlreadyHasIDCard bool   `json:"is_already_has_idcard"`
}

type OTPVerifyPayload struct {
	UserID    string `json:"user_id"`
	OTPNumber string `json:"otp_number"`
}
type IOTPService interface {
	Create(payload *OTPCreatePayload) (*models.UserOTP, core.IError)
	Verify(payload *OTPVerifyPayload) core.IError
	Find(id string) (*models.UserOTP, core.IError)
	FindByUserID(userID string) (*models.UserOTP, core.IError)
	RevokeByUserID(id string) core.IError
}
type OTPService struct {
	ctx         core.IContext
	userService IUserService
}

func NewOTPService(ctx core.IContext) IOTPService {
	return &OTPService{
		ctx:         ctx,
		userService: NewUserService(ctx),
	}
}
func (s OTPService) Verify(payload *OTPVerifyPayload) core.IError {
	otp := &models.UserOTP{}
	err := s.ctx.DB().Where("user_id = ? AND otp_number = ? AND (revoked_at IS NULL AND verified_at IS NULL)", payload.UserID, payload.OTPNumber).First(otp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return s.ctx.NewError(err, emsgs.OTPNotValid)
	}
	if err != nil {
		return s.ctx.NewError(err, errmsgs.DBError)
	}
	err = s.ctx.DB().Model(models.UserOTP{}).
		Where("id = ?", otp.ID).
		Updates(models.UserOTP{
			VerifiedAt: utils.GetCurrentDateTime(),
			UpdatedAt:  utils.GetCurrentDateTime(),
		}).Error
	if err != nil {
		return s.ctx.NewError(err, errmsgs.DBError)
	}
	return nil
}
func (s OTPService) Create(payload *OTPCreatePayload) (*models.UserOTP, core.IError) {
	user, ierr := s.userService.FindByID(payload.UserID)
	if ierr != nil && !errmsgs.IsNotFoundError(ierr) {
		return nil, s.ctx.NewError(ierr, ierr)
	}
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}

	ierr = s.RevokeByUserID(payload.UserID)
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}

	uid := utils.GetUUID()
	otpNumber := fmt.Sprintf("%04d", rand.Intn(10000))
	err := s.ctx.DB().Create(&models.UserOTP{
		ID:        uid,
		OTPNumber: otpNumber,
		UserID:    payload.UserID,
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}).Error
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	if payload.IsAlreadyHasIDCard {
		err = helpers.SendRecoveryEmail(
			s.ctx.ENV().String(consts.ENVSMTPHost),
			s.ctx.ENV().String(consts.ENVSMTPPort),
			s.ctx.ENV().String(consts.ENVSenderEmail),
			user.Email,
			otpNumber,
		)
	} else {
		err = helpers.SendRegisterEmail(
			s.ctx.ENV().String(consts.ENVSMTPHost),
			s.ctx.ENV().String(consts.ENVSMTPPort),
			s.ctx.ENV().String(consts.ENVSenderEmail),
			user.Email,
			otpNumber,
		)
	}

	if err != nil {
		return nil, s.ctx.NewError(err, emsgs.EmailError)
	}

	return s.Find(uid)
}

func (s OTPService) Find(id string) (*models.UserOTP, core.IError) {
	otp := &models.UserOTP{}
	err := s.ctx.DB().Where("id = ?", id).First(otp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.UserNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return otp, nil
}
func (s OTPService) FindByUserID(userID string) (*models.UserOTP, core.IError) {
	otp := &models.UserOTP{}
	err := s.ctx.DB().Where("user_id = ? AND revoked_at IS NULL", userID).First(otp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.UserNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return otp, nil
}

func (s OTPService) RevokeByUserID(id string) core.IError {
	err := s.ctx.DB().Model(models.UserOTP{}).
		Where("user_id = ? AND revoked_at IS NULL", id).
		Updates(models.UserOTP{
			RevokedAt: utils.GetCurrentDateTime(),
			UpdatedAt: utils.GetCurrentDateTime(),
		}).Error
	if err != nil {
		return s.ctx.NewError(err, errmsgs.DBError)
	}
	return nil
}
