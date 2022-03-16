package services

import (
	"errors"

	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"gorm.io/gorm"
)

type IConfigDIDService interface {
	GetDIDAddress() (*models.ConfigDID, core.IError)
}

type configDIDService struct {
	ctx core.IContext
}

func NewConfigDIDService(ctx core.IContext) IConfigDIDService {
	return &configDIDService{ctx: ctx}
}

func (s configDIDService) GetDIDAddress() (*models.ConfigDID, core.IError) {
	did := &models.ConfigDID{}
	err := s.ctx.DB().First(did).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, errmsgs.NotFound)
	}

	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return did, nil
}
