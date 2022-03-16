package seeds

import (
	"errors"
	"fmt"
	"os"

	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	"gitlab.finema.co/finema/etda/mobile-app-api/services"
	core "ssi-gitlab.teda.th/ssi/core"
	"gorm.io/gorm"
)

type tokenSeed struct {
	ctx      core.IContext
	tokenSvc services.ITokenService
}

func NewTokenSeed(ctx core.IContext) *tokenSeed {
	return &tokenSeed{
		ctx:      ctx,
		tokenSvc: services.NewTokenService(ctx),
	}
}
func (s tokenSeed) Run() error {
	token := &models.Token{}
	err := s.ctx.DB().First(token).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}
	token = models.NewToken()
	token.Role = consts.TokenAdminRole
	err = s.ctx.DB().Create(token).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}
	return nil
}
