// +build e2e

package services

import (
	"github.com/stretchr/testify/suite"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"testing"
)

type RequestServiceSuite struct {
	suite.Suite
	ctx                core.IContext
	rs                 IRequestService
	tkCreateS          *RequestServiceTestKit
	tkCreateE          *RequestServiceTestKit
	tkFindS            *RequestServiceTestKit
	tkFindE            *RequestServiceTestKit
	tkUpdateS          *RequestServiceTestKit
	tkUpdateE          *RequestServiceTestKit
	tkPaginationByDIDS *RequestServiceTestKit
}

func TestRequestServiceSuite(t *testing.T) {
	suite.Run(t, new(RequestServiceSuite))
}

func (r *RequestServiceSuite) BeforeTest(_, _ string) {
	r.ctx = newMockContext()
	r.rs = NewRequestService(r.ctx)

	r.tkCreateS = NewRequestServiceTestKitCreateS(r.rs)
	err := r.tkCreateS.beforeTest(r.ctx, r.tkCreateS)
	r.NoError(err)

	r.tkCreateE = NewRequestServiceTestKitCreateE(r.rs)
	err = r.tkCreateE.beforeTest(r.ctx, r.tkCreateE)
	r.NoError(err)

	r.tkFindS = NewRequestServiceTestKitFindS(r.rs)
	err = r.tkFindS.beforeTest(r.ctx, r.tkFindS)
	r.NoError(err)

	r.tkFindE = NewRequestServiceTestKitFindE(r.rs)
	err = r.tkFindE.beforeTest(r.ctx, r.tkFindE)
	r.NoError(err)

	r.tkUpdateS = NewRequestServiceTestKitUpdateS(r.rs)
	err = r.tkUpdateS.beforeTest(r.ctx, r.tkUpdateS)
	r.NoError(err)

	r.tkUpdateE = NewRequestServiceTestKitUpdateE(r.rs)
	err = r.tkUpdateE.beforeTest(r.ctx, r.tkUpdateE)
	r.NoError(err)

	r.tkPaginationByDIDS = NewRequestServiceTestKitPaginationByDID(r.rs)
	err = r.tkPaginationByDIDS.beforeTest(r.ctx, r.tkPaginationByDIDS)
	r.NoError(err)
}

func (r *RequestServiceSuite) AfterTest(_, _ string) {
	err := r.tkCreateS.afterTest(r.ctx, r.tkCreateS)
	r.NoError(err)

	err = r.tkCreateE.afterTest(r.ctx, r.tkCreateE)
	r.NoError(err)

	err = r.tkFindS.afterTest(r.ctx, r.tkFindS)
	r.NoError(err)

	err = r.tkFindE.afterTest(r.ctx, r.tkFindE)
	r.NoError(err)

	err = r.tkUpdateS.afterTest(r.ctx, r.tkUpdateS)
	r.NoError(err)

	err = r.tkUpdateE.afterTest(r.ctx, r.tkUpdateE)
	r.NoError(err)

	err = r.tkPaginationByDIDS.afterTest(r.ctx, r.tkPaginationByDIDS)
	r.NoError(err)
}

func (r *RequestServiceSuite) TestRequestService_Create_ExpectSuccess() {
	request, ierr := r.rs.Create(&RequestCreatePayload{
		RequestID:      r.tkCreateS.dummyRequestID,
		RequestData:    r.tkCreateS.dummyRequestData,
		SchemaType:     r.tkCreateS.dummySchemaType,
		CredentialType: r.tkCreateS.dummyCredentialType,
		Signer:         r.tkCreateS.dummySigner,
		Requester:      r.tkCreateS.dummyRequester,
	})
	r.NoError(ierr)
	r.NotNil(request)

	r.Equal(r.tkCreateS.dummyRequestID, request.RequestID)
	r.Equal(r.tkCreateS.dummyRequestData, request.RequestData)
	r.Equal(r.tkCreateS.dummySchemaType, request.SchemaType)
	r.Equal(r.tkCreateS.dummyCredentialType, request.CredentialType)
	r.Equal(r.tkCreateS.dummySigner, request.Signer)
	r.Equal(r.tkCreateS.dummyRequester, request.Requester)
}

func (r *RequestServiceSuite) TestRequestService_Create_ExpectDBError() {
	result, ierr := r.rs.Create(&RequestCreatePayload{
		RequestID:      r.tkCreateE.dummyRequestID,
		RequestData:    r.tkCreateE.dummyRequestData,
		SchemaType:     r.tkCreateE.dummySchemaType,
		CredentialType: r.tkCreateE.dummyCredentialType,
		Signer:         r.tkCreateE.dummySigner,
		Requester:      r.tkCreateE.dummyRequester,
	})
	r.Error(ierr)
	r.Nil(result)
}

func (r *RequestServiceSuite) TestRequestService_Find_ExpectSuccess() {
	request, ierr := r.rs.Find(r.tkFindS.dummyRequestID)
	r.NoError(ierr)
	r.NotNil(request)

	r.Equal(r.tkFindS.dummyRequestID, request.RequestID)
	r.Equal(r.tkFindS.dummyRequestData, request.RequestData)
	r.Equal(r.tkFindS.dummySchemaType, request.SchemaType)
	r.Equal(r.tkFindS.dummyCredentialType, request.CredentialType)
	r.Equal(r.tkFindS.dummySigner, request.Signer)
	r.Equal(r.tkFindS.dummyRequester, request.Requester)
}

func (r *RequestServiceSuite) TestRequestService_Update_ExpectSuccess() {
	request, ierr := r.rs.Find(r.tkUpdateS.dummyRequestID)
	r.NoError(ierr)
	r.NotNil(request)
	r.Equal(consts.RequestStatusUnsigned, request.Status)

	uRequest, ierr := r.rs.Update(request.ID, &RequestUpdatePayload{
		Status: consts.RequestStatusSigned,
	})
	r.NoError(ierr)
	r.NotNil(uRequest)

	r.NotEqual(request.Status, uRequest.Status)
	r.Equal(consts.RequestStatusSigned, uRequest.Status)
}

func (r *RequestServiceSuite) TestRequestService_Update_ExpectNotFoundError() {
	request, ierr := r.rs.Update(r.tkUpdateE.dummyID, &RequestUpdatePayload{
		Status: consts.RequestStatusSigned,
	})
	r.Error(ierr)
	r.True(errmsgs.IsNotFoundError(ierr))
	r.Nil(request)
}

func (r *RequestServiceSuite) TestRequestService_PaginationByDID_ExpectSuccess() {
	requests, page, ierr := r.rs.PaginationByDID(r.tkPaginationByDIDS.dummySigner, &core.PageOptions{}, &RequestPaginationOptions{})
	r.NoError(ierr)
	r.NotNil(page)
	r.NotNil(requests)

	r.Equal(int64(5), page.Count)
	r.Equal(len(requests), int(page.Count))

	actualSignedStatus := 0
	actualUnSignedStatus := 0

	for _, r := range requests {
		if r.Status == consts.RequestStatusSigned {
			actualSignedStatus += 1
		}
	}

	for _, r := range requests {
		if r.Status == consts.RequestStatusUnsigned {
			actualUnSignedStatus += 1
		}
	}

	r.Equal(actualSignedStatus, 3)
	r.Equal(actualUnSignedStatus, 2)
}
