package services

import (
	"github.com/stretchr/testify/mock"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type EKYCServiceTestKit struct {
	s                IEKYCService
	dummyIDCardNo    string
	dummyFirstName   string
	dummyLastName    string
	dummyLaserID     string
	dummyDateOfBirth string
	beforeTest       func(ctx core.IContext, t *EKYCServiceTestKit) error
	afterTest        func(ctx core.IContext, t *EKYCServiceTestKit) error
}

func NewEKYCServiceVerifySTestKit(s IEKYCService) *EKYCServiceTestKit {
	return &EKYCServiceTestKit{
		s:                s,
		dummyIDCardNo:    utils.GetUUID(),
		dummyFirstName:   "dummy-first-name",
		dummyLastName:    "dummy-last-name",
		dummyLaserID:     utils.GetUUID(),
		dummyDateOfBirth: utils.GetCurrentDateTime().String(),
		beforeTest: func(ctx core.IContext, t *EKYCServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *EKYCServiceTestKit) error {
			return nil
		},
	}
}

func NewEKYCServiceVerifyETestKit(s IEKYCService) *EKYCServiceTestKit {
	return &EKYCServiceTestKit{
		s:                s,
		dummyIDCardNo:    utils.GetUUID(),
		dummyFirstName:   "dummy-first-name",
		dummyLastName:    "dummy-last-name",
		dummyLaserID:     utils.GetUUID(),
		dummyDateOfBirth: utils.GetCurrentDateTime().String(),
		beforeTest: func(ctx core.IContext, t *EKYCServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *EKYCServiceTestKit) error {
			return nil
		},
	}
}

type MockIdentityProofing struct {
	mock.Mock
}

func NewMockIdentityProofing() *MockIdentityProofing {
	return &MockIdentityProofing{}
}

func (m *MockIdentityProofing) IDCardVerify(payload *identityProofingIDCardVerifyPayload) (bool, string, core.IError) {
	args := m.Called(payload)
	if _, ok := args.Get(2).(core.IError); ok {
		return args.Bool(0), args.String(1), args.Get(2).(core.IError)
	}
	return args.Bool(0), args.String(1), nil
}
