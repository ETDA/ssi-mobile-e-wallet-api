package seeds

import (
	"fmt"
	"os"

	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	"gitlab.finema.co/finema/etda/mobile-app-api/services"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type ConfigDIDSeed struct {
	ctx          core.IContext
	configDIDSvc services.IConfigDIDService
	didSvc       services.IDIDService
}

func NewConfigDIDSeed(ctx core.IContext) *ConfigDIDSeed {
	return &ConfigDIDSeed{
		ctx:          ctx,
		configDIDSvc: services.NewConfigDIDService(ctx),
		didSvc:       services.NewDIDService(ctx),
	}
}
func (s ConfigDIDSeed) Run() error {
	_, ierr := s.configDIDSvc.GetDIDAddress()
	if !errmsgs.IsNotFoundError(ierr) {
		fmt.Fprintf(os.Stderr, "%v", ierr)
		return ierr
	}
	keygen, err := utils.GenerateKeyPair()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return ierr
	}
	did, ierr := s.didSvc.RegisterDID(keygen.PublicKeyPem, keygen.PrivateKeyPem)
	if ierr != nil {
		fmt.Fprintf(os.Stderr, "%v", ierr)
		return ierr
	}
	err = s.ctx.DB().Create(&models.ConfigDID{
		DIDAddress:    did.ID,
		PublicKeyPEM:  keygen.PublicKeyPem,
		PrivateKeyPEM: keygen.PrivateKeyPem,
	}).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return ierr
	}
	return nil
}
