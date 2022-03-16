// +build e2e

package services

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/views"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"testing"
)

type DeviceTestSuite struct {
	suite.Suite
	ctx            core.IContext
	ds             IDeviceService
	us             IUserService
	tsCreate       *DeviceServiceTestKit
	tsFindByUserID *DeviceServiceTestKit
	tsUpdateTokenS *DeviceServiceTestKit
	tsUpdateTokenE *DeviceServiceTestKit
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceTestSuite))
}

func (d *DeviceTestSuite) BeforeTest(_, _ string) {
	d.ctx = newMockContext()
	d.ds = NewDeviceService(d.ctx)
	d.us = NewUserService(d.ctx)

	d.tsCreate = NewDeviceServiceCreateTestKit(d.ds)
	err := d.tsCreate.beforeTest(d.ctx, d.tsCreate)
	d.NoError(err)

	d.tsFindByUserID = NewDeviceServiceFindByUserIDTestKit(d.ds)
	err = d.tsFindByUserID.beforeTest(d.ctx, d.tsFindByUserID)
	d.NoError(err)

	d.tsUpdateTokenS = NewDeviceServiceUpdateTokenSTestKit(d.ds)
	err = d.tsUpdateTokenS.beforeTest(d.ctx, d.tsUpdateTokenS)
	d.NoError(err)

	d.tsUpdateTokenE = NewDeviceServiceUpdateTokenETestKit(d.ds)
	err = d.tsUpdateTokenE.beforeTest(d.ctx, d.tsUpdateTokenE)
	d.NoError(err)
}

func (d *DeviceTestSuite) AfterTest(_, _ string) {
	err := d.tsCreate.afterTest(d.ctx, d.tsCreate)
	d.NoError(err)

	err = d.tsFindByUserID.afterTest(d.ctx, d.tsFindByUserID)
	d.NoError(err)

	err = d.tsUpdateTokenS.afterTest(d.ctx, d.tsUpdateTokenS)
	d.NoError(err)
}

func (d *DeviceTestSuite) TestDeviceService_Create_ExpectCorrectValue() {
	ierr := d.ds.Create(&DeviceCreatePayload{
		DIDAddress: d.tsCreate.dummyDIDAddress,
		ID:         d.tsCreate.dummyUserID,
		Device: &DeviceCreatePayloadDevice{
			Name:      d.tsCreate.dummyDevice.Name,
			OS:        d.tsCreate.dummyDevice.OS,
			OSVersion: d.tsCreate.dummyDevice.OSVersion,
			Model:     d.tsCreate.dummyDevice.Model,
			UUID:      d.tsCreate.dummyDevice.UUID,
		},
	})
	d.NoError(ierr)

	u, ierr := d.us.FindByID(d.tsCreate.dummyUserID)
	d.NoError(ierr)
	d.NotNil(u)

	de, ierr := d.ds.FindByUserID(d.tsCreate.dummyUserID)
	d.NoError(ierr)
	d.NotNil(d)

	ud := views.NewUserDevice(u, de)
	d.NotNil(ud)

	d.Equal(u.DIDAddress, ud.DIDAddress)
	d.Equal(u.IDCardNo, ud.IDCardNo)
	d.Equal(u.FirstName, ud.FirstName)
	d.Equal(u.LastName, ud.LastName)
	d.Equal(de.OS, ud.Device.OS)
	d.Equal(de.OSVersion, ud.Device.OSVersion)
	d.Equal(de.Name, ud.Device.Name)
	d.Equal(de.Model, ud.Device.Model)
	d.Equal(de.UUID, ud.Device.UUID)
}

func (d *DeviceTestSuite) TestDeviceService_UpdateToken_ExpectCollectValue() {
	device, ierr := d.ds.FindByUserID(d.tsUpdateTokenS.dummyUserID)
	d.NoError(ierr)
	d.NotNil(device)
	d.Empty(device.Token)

	ierr = d.ds.UpdateToken(&TokenUpdatePayload{
		DIDAddress: d.tsUpdateTokenS.dummyDIDAddress,
		UUID:       d.tsUpdateTokenS.dummyDevice.UUID,
		Token:      d.tsUpdateTokenS.dummyToken,
	})
	d.NoError(ierr)

	updatedDevice, ierr := d.ds.FindByUserID(d.tsUpdateTokenS.dummyUserID)
	d.NoError(ierr)
	d.NotNil(updatedDevice)
	d.NotEmpty(updatedDevice.Token)
	d.Equal(d.tsUpdateTokenS.dummyToken, updatedDevice.Token)
}

func (d *DeviceTestSuite) TestDeviceService_UpdateToken_ExpectUUIDNotFoundError() {
	device, ierr := d.ds.FindByUserID(d.tsUpdateTokenE.dummyUserID)
	d.NoError(ierr)
	d.NotNil(device)
	d.Empty(device.Token)

	ierr = d.ds.UpdateToken(&TokenUpdatePayload{
		DIDAddress: d.tsUpdateTokenE.dummyDIDAddress,
		UUID:       "<uuid not mismatch from dummy did_address>",
		Token:      d.tsUpdateTokenE.dummyToken,
	})
	d.Error(ierr)
	d.True(errors.Is(ierr, emsgs.DeviceUUIDNotFound))
	d.True(errmsgs.IsNotFoundError(ierr))
}
