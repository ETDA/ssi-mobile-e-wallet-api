package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/views"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type IVerifyService interface {
	Create(jwt string) (*views.Verify, core.IError)
	Get(id string) (string, core.IError)
}
type verifyServiceService struct {
	ctx core.IContext
}

func NewVerifyService(ctx core.IContext) IVerifyService {
	return &verifyServiceService{ctx: ctx}
}

func (s verifyServiceService) Create(jwt string) (*views.Verify, core.IError) {
	id := utils.GetUUID()
	err := s.ctx.Cache().Set(fmt.Sprintf("verify_jwt.%s", id), jwt, time.Minute*5)
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.CacheError)
	}

	return &views.Verify{
		Endpoint:  fmt.Sprintf("%s/mobile/verify/%s", s.ctx.ENV().String(consts.ENVMobileServiceBaseURL), id),
		Operation: consts.OperationVerifyVC,
	}, nil
}

func (s verifyServiceService) Get(id string) (string, core.IError) {
	var jwt string
	err := s.ctx.Cache().Get(&jwt, fmt.Sprintf("verify_jwt.%s", id))
	if errors.Is(err, redis.Nil) {
		return "", s.ctx.NewError(err, errmsgs.NotFound)
	}
	if err != nil {
		return "", s.ctx.NewError(err, errmsgs.CacheError)
	}

	err = s.ctx.Cache().Del(fmt.Sprintf("verify_jwt.%s", id))
	if err != nil {
		return "", s.ctx.NewError(err, errmsgs.CacheError)
	}

	return jwt, nil
}
