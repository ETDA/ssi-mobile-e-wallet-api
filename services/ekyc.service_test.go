// +build e2e

package services

import (
	"github.com/stretchr/testify/suite"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"testing"
)

type EKYCTestSuite struct {
	suite.Suite
	ctx       core.IContext
	es        IEKYCService
	us        IUserService
	mips      *MockIdentityProofing
	tsVerifyS *EKYCServiceTestKit
	tsVerifyE *EKYCServiceTestKit
}

func TestEKYCTestSuite(t *testing.T) {
	suite.Run(t, new(EKYCTestSuite))
}

func (e *EKYCTestSuite) BeforeTest(_, _ string) {
	e.mips = NewMockIdentityProofing()

	e.ctx = newMockContext()
	e.us = NewUserService(e.ctx)
	e.es = NewEKYCService(e.ctx, e.us, e.mips)

	e.tsVerifyS = NewEKYCServiceVerifySTestKit(e.es)
	err := e.tsVerifyS.beforeTest(e.ctx, e.tsVerifyS)
	e.NoError(err)

	e.tsVerifyE = NewEKYCServiceVerifyETestKit(e.es)
	err = e.tsVerifyE.beforeTest(e.ctx, e.tsVerifyE)
	e.NoError(err)
}

func (e *EKYCTestSuite) AfterTest(_, _ string) {
	err := e.tsVerifyS.afterTest(e.ctx, e.tsVerifyS)
	e.NoError(err)

	err = e.tsVerifyE.afterTest(e.ctx, e.tsVerifyE)
	e.NoError(err)
}

func (e *EKYCTestSuite) TestEKYCService_Verify_ExpectCorrectValue() {
	e.mips.On("IDCardVerify", &identityProofingIDCardVerifyPayload{
		CardID:    e.tsVerifyS.dummyIDCardNo,
		FirstName: e.tsVerifyS.dummyFirstName,
		LastName:  e.tsVerifyS.dummyLastName,
		LaserID:   e.tsVerifyS.dummyLaserID,
		Birthdate: e.tsVerifyS.dummyDateOfBirth,
	}).Return(true, "", nil)

	ev, ierr := e.es.Verify(&EKYCVerifyPayload{
		IDCardNo:    e.tsVerifyS.dummyIDCardNo,
		FirstName:   e.tsVerifyS.dummyFirstName,
		LastName:    e.tsVerifyS.dummyLastName,
		LaserID:     e.tsVerifyS.dummyLaserID,
		DateOfBirth: e.tsVerifyS.dummyDateOfBirth,
	})
	e.NoError(ierr)
	e.NotNil(ev)
	e.True(ev.CardStatus)

	user, ierr := e.us.FindByID(ev.ID)
	e.NoError(ierr)
	e.NotNil(user)

	e.Equal(e.tsVerifyS.dummyIDCardNo, user.IDCardNo)
	e.Equal(e.tsVerifyS.dummyFirstName, user.FirstName)
	e.Equal(e.tsVerifyS.dummyLastName, user.LastName)
}

func (e *EKYCTestSuite) TestEKYCService_Verify_ExpectError() {
	e.mips.On("IDCardVerify", &identityProofingIDCardVerifyPayload{
		CardID:    e.tsVerifyE.dummyIDCardNo,
		FirstName: e.tsVerifyE.dummyFirstName,
		LastName:  e.tsVerifyE.dummyLastName,
		LaserID:   e.tsVerifyE.dummyLaserID,
		Birthdate: e.tsVerifyE.dummyDateOfBirth,
	}).Return(false, "<error-message>", errmsgs.InternalServerError)

	ev, ierr := e.es.Verify(&EKYCVerifyPayload{
		IDCardNo:    e.tsVerifyE.dummyIDCardNo,
		FirstName:   e.tsVerifyE.dummyFirstName,
		LastName:    e.tsVerifyE.dummyLastName,
		LaserID:     e.tsVerifyE.dummyLaserID,
		DateOfBirth: e.tsVerifyE.dummyDateOfBirth,
	})
	e.NoError(ierr)
	e.NotNil(ev)
	e.False(ev.CardStatus)

	user, ierr := e.us.FindByID(ev.ID)
	e.Error(ierr)
	e.True(errmsgs.IsNotFoundError(ierr))
	e.Nil(user)
}

//c0fbe8ff-036e-4317-8c0d-880485a7081f
