package services

import (
	"errors"

	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
	"gorm.io/gorm"
)

type DeviceCreatePayload struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
	Model     string `json:"model"`
	UUID      string `json:"uuid"`
}

// type DeviceCreatePayload struct {
// 	DIDAddress string                     `json:"did_address"`
// 	ID         string                     `json:"id"`
// 	Device     *DeviceCreatePayloadDevice `json:"device"`
// }

type DeviceTokenUpdatePayload struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}

type IDeviceService interface {
	Create(payload *DeviceCreatePayload) (*models.Device, core.IError)
	FindByUserID(id string) (*models.Device, core.IError)
	FindByID(uuid string) (*models.Device, core.IError)
	FindByUUID(id string) (*models.Device, core.IError)
	FindByDID(did string) (*models.Device, core.IError)
	UpdateToken(payload *DeviceTokenUpdatePayload) core.IError
	ResetDevice(newKey *NewKeyPayload, newDevice *DeviceCreatePayload, id string) (didAddress string, ierr core.IError)
}

type deviceService struct {
	ctx              core.IContext
	userService      IUserService
	configDIDService IConfigDIDService
	didService       IDIDService
}

func NewDeviceService(ctx core.IContext) IDeviceService {
	return &deviceService{
		ctx:              ctx,
		userService:      NewUserService(ctx),
		configDIDService: NewConfigDIDService(ctx),
		didService:       NewDIDService(ctx),
	}
}

func (s deviceService) Create(payload *DeviceCreatePayload) (*models.Device, core.IError) {
	device, ierr := s.FindByUserID(payload.UserID)
	if ierr != nil && !errmsgs.IsNotFoundError(ierr) {
		return nil, s.ctx.NewError(ierr, ierr)
	}
	if device != nil {
		return nil, s.ctx.NewError(emsgs.DuplicatedDIDAddress, emsgs.DuplicatedDIDAddress)
	}

	err := s.ctx.DB().Create(&models.Device{
		ID:        payload.UUID,
		Name:      payload.Name,
		OS:        payload.OS,
		OSVersion: payload.OSVersion,
		Model:     payload.Model,
		UUID:      payload.UUID,
		UserID:    payload.UserID,
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}).Error
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return s.FindByID(payload.UUID)
}

func (s deviceService) FindByID(id string) (*models.Device, core.IError) {
	device := &models.Device{}
	err := s.ctx.DB().Where("id = ?", id).First(device).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.DeviceNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return device, nil
}
func (s deviceService) FindByUUID(uuid string) (*models.Device, core.IError) {
	device := &models.Device{}
	err := s.ctx.DB().Where("uuid = ? AND deleted_at IS NULL", uuid).First(device).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.DeviceNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return device, nil
}
func (s deviceService) FindByUserID(id string) (*models.Device, core.IError) {
	device := &models.Device{}
	err := s.ctx.DB().Where("user_id = ? AND deleted_at IS NULL", id).First(device).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.DeviceNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return device, nil
}

func (s deviceService) FindByDID(did string) (*models.Device, core.IError) {
	device := &models.Device{}
	err := s.ctx.DB().Where("did_address = ?", did).First(device).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, s.ctx.NewError(err, emsgs.DeviceNotFound)
	}
	if err != nil {
		return nil, s.ctx.NewError(err, errmsgs.DBError)
	}

	return device, nil
}

func (s deviceService) UpdateToken(payload *DeviceTokenUpdatePayload) core.IError {
	device, ierr := s.FindByUUID(payload.UUID)
	if errmsgs.IsNotFoundError(ierr) {
		return s.ctx.NewError(emsgs.DeviceNotFound, emsgs.DeviceNotFound)
	}
	if ierr != nil {
		return s.ctx.NewError(ierr, ierr)
	}
	if device.UUID != payload.UUID {
		return s.ctx.NewError(emsgs.DeviceUUIDNotFound, emsgs.DeviceUUIDNotFound)
	}

	err := s.ctx.DB().Updates(&models.Device{
		ID:        payload.UUID,
		Token:     payload.Token,
		UpdatedAt: utils.GetCurrentDateTime(),
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return s.ctx.NewError(err, emsgs.DeviceNotFound)
	}
	if err != nil {
		return s.ctx.NewError(err, errmsgs.DBError)
	}

	return nil
}

func (s deviceService) ResetDevice(newKey *NewKeyPayload, newDevice *DeviceCreatePayload, id string) (didAddress string, ierr core.IError) {
	user, ierr := s.userService.FindByID(id)
	if ierr != nil {
		return "", s.ctx.NewError(ierr, ierr)
	}

	resetterDID, ierr := s.configDIDService.GetDIDAddress()
	if ierr != nil {
		return "", s.ctx.NewError(ierr, ierr)
	}

	_, ierr = s.didService.ResetterApproveReset(&RecoverResetPayload{
		DIDAddress: resetterDID.DIDAddress,
		NewKey:     newKey,
		RequestDID: utils.GetString(user.DIDAddress),
		PrivateKey: resetterDID.PrivateKeyPEM,
	})
	if ierr != nil {
		return "", s.ctx.NewError(ierr, ierr)
	}

	err := s.ctx.DB().Model(models.Device{}).
		Where("user_id = ?", user.ID).
		Updates(models.Device{DeletedAt: utils.GetCurrentDateTime()}).
		Error

	if err != nil {
		return "", s.ctx.NewError(err, errmsgs.DBError)
	}
	err = s.ctx.DB().Create(&models.Device{
		ID:        newDevice.UUID,
		Name:      newDevice.Name,
		OS:        newDevice.OS,
		OSVersion: newDevice.OSVersion,
		Model:     newDevice.Model,
		UUID:      newDevice.UUID,
		UserID:    newDevice.UserID,
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}).Error
	if err != nil {
		return "", s.ctx.NewError(err, errmsgs.DBError)
	}
	return utils.GetString(user.DIDAddress), nil
}
