package services

import (
	"fmt"
	"os"

	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

func newMockContext() core.IContext {
	env := core.NewENVPath("./../")
	envCfg := &core.ENVConfig{
		DBHost:     "localhost",
		DBName:     "my_database",
		DBUser:     "my_user",
		DBPassword: "my_password",
		DBPort:     "3306",
	}

	mysql, err := core.NewDatabase(envCfg).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "mysql: %v", err)
		os.Exit(1)
	}

	return core.NewContext(&core.ContextOptions{
		ENV: env,
		DB:  mysql,
	})
}

type DeviceServiceTestKit struct {
	s               IDeviceService
	dummyDIDAddress string
	dummyID         string
	dummyUserID     string
	dummyDevice     *DeviceCreatePayload
	dummyIDCardNo   string
	dummyFirstName  string
	dummyLastName   string
	dummyToken      string
	beforeTest      func(ctx core.IContext, t *DeviceServiceTestKit) error
	afterTest       func(ctx core.IContext, t *DeviceServiceTestKit) error
}

func NewDeviceServiceCreateTestKit(s IDeviceService) *DeviceServiceTestKit {
	return &DeviceServiceTestKit{
		s:               s,
		dummyDIDAddress: utils.GetUUID(),
		dummyID:         utils.GetUUID(),
		dummyUserID:     utils.GetUUID(),
		dummyDevice: &DeviceCreatePayload{
			Name:      "dummy-device-name",
			OS:        "dummy-device-os",
			OSVersion: "dummy-device-os-version",
			Model:     "dummy-device-model",
			UUID:      utils.GetUUID(),
		},
		dummyIDCardNo:  utils.GetUUID(),
		dummyFirstName: "dummy-first-name",
		dummyLastName:  "dummy-last-name",
		dummyToken:     utils.GetMD5Hash(utils.GetUUID()),
		beforeTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
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
		afterTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewDeviceServiceFindByUserIDTestKit(s IDeviceService) *DeviceServiceTestKit {
	return &DeviceServiceTestKit{
		s:               s,
		dummyDIDAddress: utils.GetUUID(),
		dummyID:         utils.GetUUID(),
		dummyUserID:     utils.GetUUID(),
		dummyDevice: &DeviceCreatePayload{
			Name:      "dummy-device-name",
			OS:        "dummy-device-os",
			OSVersion: "dummy-device-os-version",
			Model:     "dummy-device-model",
			UUID:      utils.GetUUID(),
		},
		dummyIDCardNo:  utils.GetUUID(),
		dummyFirstName: "dummy-first-name",
		dummyLastName:  "dummy-last-name",
		dummyToken:     utils.GetMD5Hash(utils.GetUUID()),
		beforeTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(models.Device{
				ID:        t.dummyID,
				Name:      t.dummyDevice.Name,
				OS:        t.dummyDevice.OS,
				OSVersion: t.dummyDevice.OSVersion,
				Model:     t.dummyDevice.Model,
				UUID:      t.dummyDevice.UUID,
				UserID:    t.dummyUserID,
				CreatedAt: utils.GetCurrentDateTime(),
				UpdatedAt: utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewDeviceServiceUpdateTokenSTestKit(s IDeviceService) *DeviceServiceTestKit {
	return &DeviceServiceTestKit{
		s:               s,
		dummyDIDAddress: utils.GetUUID(),
		dummyID:         utils.GetUUID(),
		dummyUserID:     utils.GetUUID(),
		dummyDevice: &DeviceCreatePayload{
			Name:      "dummy-device-name",
			OS:        "dummy-device-os",
			OSVersion: "dummy-device-os-version",
			Model:     "dummy-device-model",
			UUID:      utils.GetUUID(),
		},
		dummyIDCardNo:  utils.GetUUID(),
		dummyFirstName: "dummy-first-name",
		dummyLastName:  "dummy-last-name",
		dummyToken:     utils.GetMD5Hash(utils.GetUUID()),
		beforeTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
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

			err = ctx.DB().Create(models.Device{
				ID:        t.dummyDevice.UUID,
				Name:      t.dummyDevice.Name,
				OS:        t.dummyDevice.OS,
				OSVersion: t.dummyDevice.OSVersion,
				Model:     t.dummyDevice.Model,
				UUID:      t.dummyDevice.UUID,
				UserID:    t.dummyUserID,
				CreatedAt: utils.GetCurrentDateTime(),
				UpdatedAt: utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewDeviceServiceUpdateTokenETestKit(s IDeviceService) *DeviceServiceTestKit {
	return &DeviceServiceTestKit{
		s:               s,
		dummyDIDAddress: utils.GetUUID(),
		dummyID:         utils.GetUUID(),
		dummyUserID:     utils.GetUUID(),
		dummyDevice: &DeviceCreatePayload{
			Name:      "dummy-device-name",
			OS:        "dummy-device-os",
			OSVersion: "dummy-device-os-version",
			Model:     "dummy-device-model",
			UUID:      utils.GetUUID(),
		},
		dummyIDCardNo:  utils.GetUUID(),
		dummyFirstName: "dummy-first-name",
		dummyLastName:  "dummy-last-name",
		dummyToken:     utils.GetMD5Hash(utils.GetUUID()),
		beforeTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
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

			err = ctx.DB().Create(models.Device{
				ID:        t.dummyDevice.UUID,
				Name:      t.dummyDevice.Name,
				OS:        t.dummyDevice.OS,
				OSVersion: t.dummyDevice.OSVersion,
				Model:     t.dummyDevice.Model,
				UUID:      t.dummyDevice.UUID,
				UserID:    t.dummyUserID,
				CreatedAt: utils.GetCurrentDateTime(),
				UpdatedAt: utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *DeviceServiceTestKit) error {
			err := ctx.DB().Where("did_address = ?", t.dummyDIDAddress).Delete(&models.User{}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Where("user_id = ?", t.dummyUserID).Delete(&models.Device{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}
