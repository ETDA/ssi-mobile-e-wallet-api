package services

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/views"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
)

type IEKYCService interface {
	Verify(payload *EKYCVerifyPayload) (*views.EKYCVerify, core.IError)
}
type ekycService struct {
	ctx                     core.IContext
	userService             IUserService
	identityProofingService IIdentityProofingService
}

func NewEKYCService(ctx core.IContext, userService IUserService, identityProofingService IIdentityProofingService) IEKYCService {
	return &ekycService{
		ctx:                     ctx,
		userService:             userService,
		identityProofingService: identityProofingService,
	}
}

type EKYCVerifyPayload struct {
	IDCardNo    string
	FirstName   string
	LastName    string
	LaserID     string
	DateOfBirth string
}

func (s ekycService) Verify(payload *EKYCVerifyPayload) (*views.EKYCVerify, core.IError) {
	result, ierr := s.identityProofingService.IDCardVerify(&identityProofingIDCardVerifyPayload{
		CardID:    payload.IDCardNo,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		LaserID:   payload.LaserID,
		Birthdate: payload.DateOfBirth,
	})
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}
	if result == nil {
		return nil, s.ctx.NewError(errmsgs.InternalServerError, errmsgs.InternalServerError)
	}

	if !result.Status {
		return &views.EKYCVerify{
			CardStatus: false,
			Message:    result.Message,
		}, nil
	}

	user, _ := s.userService.FindByIDCardNumber(payload.IDCardNo)
	if user != nil {
		return views.NewEKYCViewSuccess(user), nil
	}

	newUser, ierr := s.userService.Create(&UserCreatePayload{
		IDCardNo:  payload.IDCardNo,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	})
	if ierr != nil {
		return nil, s.ctx.NewError(ierr, ierr)
	}

	return views.NewEKYCViewSuccess(newUser), nil
}
