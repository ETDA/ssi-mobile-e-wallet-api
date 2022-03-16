package services

import (
	"encoding/json"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/models"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type RequestServiceTestKit struct {
	s                   IRequestService
	dummyID             string
	dummyRequestID      string
	dummyRequestData    *json.RawMessage
	dummySchemaType     string
	dummyCredentialType string
	dummySigner         string
	dummyRequester      string
	dummyStatus         string
	beforeTest          func(ctx core.IContext, t *RequestServiceTestKit) error
	afterTest           func(ctx core.IContext, t *RequestServiceTestKit) error
}

func NewRequestServiceTestKitCreateS(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Where("request_id = ?", t.dummyRequestID).Delete(models.Request{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewRequestServiceTestKitCreateE(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Create(&models.Request{
				ID:             t.dummyID,
				RequestID:      t.dummyRequestID,
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusUnsigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Where("request_id = ?", t.dummyRequestID).Delete(models.Request{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewRequestServiceTestKitFindS(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Create(&models.Request{
				ID:             t.dummyID,
				RequestID:      t.dummyRequestID,
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusUnsigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Where("request_id = ?", t.dummyRequestID).Delete(models.Request{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewRequestServiceTestKitFindE(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			return nil
		},
	}
}

func NewRequestServiceTestKitUpdateS(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Create(&models.Request{
				ID:             t.dummyID,
				RequestID:      t.dummyRequestID,
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusUnsigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Where("request_id = ?", t.dummyRequestID).Delete(models.Request{}).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewRequestServiceTestKitUpdateE(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			return nil
		},
	}
}

func NewRequestServiceTestKitPaginationByDID(s IRequestService) *RequestServiceTestKit {
	return &RequestServiceTestKit{
		s:              s,
		dummyID:        utils.GetUUID(),
		dummyRequestID: utils.GetUUID(),
		dummyRequestData: func() *json.RawMessage {
			data := &core.Map{
				"@context": []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				"type": []string{
					"VerifiableCredential",
					"UniversityDegreeCredential",
				},
				"credentialSubject": core.Map{
					"degree": core.Map{
						"type": "BachelorDegree",
						"name": "Bachelor of Science and Arts",
					},
				},
			}
			b := &json.RawMessage{}
			_ = utils.MapToStruct(data, b)
			return b
		}(),
		dummySchemaType:     "dummy-schema-type",
		dummyCredentialType: "dummy-credential-type",
		dummySigner:         utils.GenerateDID(utils.GetUUID(), "example"),
		dummyRequester:      utils.GenerateDID(utils.GetUUID(), "example"),
		dummyStatus:         "dummy-status",
		beforeTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Create(&models.Request{
				ID:             utils.GetUUID(),
				RequestID:      utils.GetUUID(),
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusUnsigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(&models.Request{
				ID:             utils.GetUUID(),
				RequestID:      utils.GetUUID(),
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusUnsigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(&models.Request{
				ID:             utils.GetUUID(),
				RequestID:      utils.GetUUID(),
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusSigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(&models.Request{
				ID:             utils.GetUUID(),
				RequestID:      utils.GetUUID(),
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusSigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			err = ctx.DB().Create(&models.Request{
				ID:             utils.GetUUID(),
				RequestID:      utils.GetUUID(),
				RequestData:    t.dummyRequestData,
				SchemaType:     t.dummySchemaType,
				CredentialType: t.dummyCredentialType,
				Signer:         t.dummySigner,
				Requester:      t.dummyRequester,
				Status:         consts.RequestStatusSigned,
				CreatedAt:      utils.GetCurrentDateTime(),
				UpdatedAt:      utils.GetCurrentDateTime(),
			}).Error
			if err != nil {
				return err
			}

			return nil
		},
		afterTest: func(ctx core.IContext, t *RequestServiceTestKit) error {
			err := ctx.DB().Where("signer = ?", t.dummySigner).Error
			if err != nil {
				return nil
			}

			return nil
		},
	}
}
