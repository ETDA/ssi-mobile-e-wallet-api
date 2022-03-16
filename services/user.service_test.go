// +build e2e

package services

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	ctx          core.IContext
	s            IUserService
	tsCreateS    *UserServiceTestKit
	tsCreateE    *UserServiceTestKit
	tsFindByIDS  *UserServiceTestKit
	tsFindByIDE  *UserServiceTestKit
	tsFindByDIDS *UserServiceTestKit
	tsFindByDIDE *UserServiceTestKit
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (u *UserTestSuite) BeforeTest(_, _ string) {
	u.ctx = newMockContext()
	u.s = NewUserService(u.ctx)

	u.tsCreateS = NewUserServiceCreateSTestKit(u.s)
	err := u.tsCreateS.beforeTest(u.ctx, u.tsCreateS)
	u.NoError(err)

	u.tsCreateE = NewUserServiceCreateETestKit(u.s)
	err = u.tsCreateE.beforeTest(u.ctx, u.tsCreateE)
	u.NoError(err)

	u.tsFindByIDS = NewUserServiceFindByIDSTestKit(u.s)
	err = u.tsFindByIDS.beforeTest(u.ctx, u.tsFindByIDS)
	u.NoError(err)

	u.tsFindByIDE = NewUserServiceFindByIDETestKit(u.s)
	err = u.tsFindByIDE.beforeTest(u.ctx, u.tsFindByIDE)
	u.NoError(err)

	u.tsFindByDIDS = NewUserServiceFindByDIDSTestKit(u.s)
	err = u.tsFindByDIDS.beforeTest(u.ctx, u.tsFindByDIDS)
	u.NoError(err)

	u.tsFindByDIDE = NewUserServiceFindByDIDETestKit(u.s)
	err = u.tsFindByDIDE.beforeTest(u.ctx, u.tsFindByDIDE)
	u.NoError(err)
}

func (u *UserTestSuite) AfterTest(_, _ string) {
	err := u.tsCreateS.afterTest(u.ctx, u.tsFindByDIDE)
	u.NoError(err)

	err = u.tsCreateE.afterTest(u.ctx, u.tsCreateE)
	u.NoError(err)

	err = u.tsFindByIDS.afterTest(u.ctx, u.tsFindByIDS)
	u.NoError(err)

	err = u.tsFindByIDE.afterTest(u.ctx, u.tsFindByIDE)
	u.NoError(err)

	err = u.tsFindByDIDS.afterTest(u.ctx, u.tsFindByDIDS)
	u.NoError(err)

	err = u.tsFindByDIDE.afterTest(u.ctx, u.tsFindByDIDE)
	u.NoError(err)
}

func (u *UserTestSuite) TestUserService_Create_ExpectCollectValue() {
	user, ierr := u.s.Create(&UserCreatePayload{
		IDCardNo:  u.tsCreateS.dummyIDCardNo,
		FirstName: u.tsCreateS.dummyFirstName,
		LastName:  u.tsCreateS.dummyLastName,
	})
	u.NoError(ierr)
	u.NotNil(user)

	u.Equal(u.tsCreateS.dummyIDCardNo, user.IDCardNo)
	u.Equal(u.tsCreateS.dummyFirstName, user.FirstName)
	u.Equal(u.tsCreateS.dummyLastName, user.LastName)
}

func (u *UserTestSuite) TestUserService_Create_ExpectDuplicatedError() {
	user, ierr := u.s.Create(&UserCreatePayload{
		IDCardNo:  u.tsCreateE.dummyIDCardNo,
		FirstName: u.tsCreateE.dummyFirstName,
		LastName:  u.tsCreateE.dummyLastName,
	})
	u.Error(ierr)
	u.True(errors.Is(emsgs.DuplicatedUser, ierr))
	u.Nil(user)
}

func (u *UserTestSuite) TestUserService_FindByID_ExpectCollectValue() {
	user, ierr := u.s.FindByID(u.tsFindByIDS.dummyUserID)
	u.NoError(ierr)
	u.NotNil(user)

	u.Equal(u.tsFindByIDS.dummyUserID, user.ID)
	u.Equal(u.tsFindByIDS.dummyIDCardNo, user.IDCardNo)
	u.Equal(u.tsFindByIDS.dummyFirstName, user.FirstName)
	u.Equal(u.tsFindByIDS.dummyLastName, user.LastName)
	u.Equal(u.tsFindByIDS.dummyDIDAddress, user.DIDAddress)
}

func (u *UserTestSuite) TestUserService_FindByID_ExpectNotFoundError() {
	user, ierr := u.s.FindByID(u.tsFindByIDE.dummyUserID)
	u.Error(ierr)
	u.True(errmsgs.IsNotFoundError(ierr))
	u.Nil(user)
}

func (u *UserTestSuite) TestUserService_FindByDID_ExpectCollectValue() {
	user, ierr := u.s.FindByDID(u.tsFindByDIDS.dummyDIDAddress)
	u.NoError(ierr)
	u.NotNil(user)

	u.Equal(u.tsFindByDIDS.dummyUserID, user.ID)
	u.Equal(u.tsFindByDIDS.dummyIDCardNo, user.IDCardNo)
	u.Equal(u.tsFindByDIDS.dummyFirstName, user.FirstName)
	u.Equal(u.tsFindByDIDS.dummyLastName, user.LastName)
	u.Equal(u.tsFindByDIDS.dummyDIDAddress, user.DIDAddress)
}

func (u *UserTestSuite) TestUserService_FindByDID_ExpectNotFoundError() {
	user, ierr := u.s.FindByDID(u.tsFindByDIDE.dummyDIDAddress)
	u.Error(ierr)
	u.True(errmsgs.IsNotFoundError(ierr))
	u.Nil(user)
}
