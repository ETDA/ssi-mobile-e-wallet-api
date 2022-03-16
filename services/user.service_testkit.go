package services

import (
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type UserServiceTestKit struct {
	s               IUserService
	dummyUserID     string
	dummyDIDAddress string
	dummyIDCardNo   string
	dummyFirstName  string
	dummyLastName   string
	dummyLaserNo    string
	beforeTest      func(ctx core.IContext, t *UserServiceTestKit) error
	afterTest       func(ctx core.IContext, t *UserServiceTestKit) error
}

func NewUserServiceCreateSTestKit(s IUserService) *UserServiceTestKit {
	return &UserServiceTestKit{
		dummyUserID:     utils.GetUUID(),
		dummyDIDAddress: utils.GetUUID(),
		dummyIDCardNo:   utils.GetUUID(),
		dummyFirstName:  "dummy-first-name",
		dummyLastName:   "dummy-first-name",
		dummyLaserNo:    utils.GetUUID(),
		beforeTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("id_card_no = ?", t.dummyIDCardNo).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewUserServiceCreateETestKit(s IUserService) *UserServiceTestKit {
	return &UserServiceTestKit{
		dummyUserID:     utils.GetUUID(),
		dummyDIDAddress: utils.GetUUID(),
		dummyIDCardNo:   utils.GetUUID(),
		dummyFirstName:  "dummy-first-name",
		dummyLastName:   "dummy-first-name",
		dummyLaserNo:    utils.GetUUID(),
		beforeTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("id_card_no = ?", t.dummyIDCardNo).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(models.User{
				ID:         utils.GetUUID(),
				IDCardNo:   t.dummyIDCardNo,
				FirstName:  t.dummyFirstName,
				LastName:   t.dummyLastName,
				DIDAddress: &t.dummyDIDAddress,
				CreatedAt:  utils.GetCurrentDateTime(),
				UpdatedAt:  utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("id_card_no = ?", t.dummyIDCardNo).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewUserServiceFindByIDSTestKit(s IUserService) *UserServiceTestKit {
	return &UserServiceTestKit{
		dummyUserID:     utils.GetUUID(),
		dummyDIDAddress: utils.GetUUID(),
		dummyIDCardNo:   utils.GetUUID(),
		dummyFirstName:  "dummy-first-name",
		dummyLastName:   "dummy-first-name",
		dummyLaserNo:    utils.GetUUID(),
		beforeTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("id = ?", t.dummyUserID).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(models.User{
				ID:         t.dummyUserID,
				IDCardNo:   t.dummyIDCardNo,
				FirstName:  t.dummyFirstName,
				LastName:   t.dummyLastName,
				DIDAddress: &t.dummyDIDAddress,
				CreatedAt:  utils.GetCurrentDateTime(),
				UpdatedAt:  utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("id = ?", t.dummyUserID).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewUserServiceFindByIDETestKit(s IUserService) *UserServiceTestKit {
	return &UserServiceTestKit{
		dummyUserID:     utils.GetUUID(),
		dummyDIDAddress: utils.GetUUID(),
		dummyIDCardNo:   utils.GetUUID(),
		dummyFirstName:  "dummy-first-name",
		dummyLastName:   "dummy-first-name",
		dummyLaserNo:    utils.GetUUID(),
		beforeTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			return nil
		},
	}
}

func NewUserServiceFindByDIDSTestKit(s IUserService) *UserServiceTestKit {
	return &UserServiceTestKit{
		dummyUserID:     utils.GetUUID(),
		dummyDIDAddress: utils.GetUUID(),
		dummyIDCardNo:   utils.GetUUID(),
		dummyFirstName:  "dummy-first-name",
		dummyLastName:   "dummy-first-name",
		dummyLaserNo:    utils.GetUUID(),
		beforeTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(models.User{
				ID:         t.dummyUserID,
				IDCardNo:   t.dummyIDCardNo,
				FirstName:  t.dummyFirstName,
				LastName:   t.dummyLastName,
				DIDAddress: &t.dummyDIDAddress,
				CreatedAt:  utils.GetCurrentDateTime(),
				UpdatedAt:  utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewUserServiceFindByDIDETestKit(s IUserService) *UserServiceTestKit {
	return &UserServiceTestKit{
		dummyUserID:     utils.GetUUID(),
		dummyDIDAddress: utils.GetUUID(),
		dummyIDCardNo:   utils.GetUUID(),
		dummyFirstName:  "dummy-first-name",
		dummyLastName:   "dummy-first-name",
		dummyLaserNo:    utils.GetUUID(),
		beforeTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *UserServiceTestKit) error {
			return nil
		},
	}
}
