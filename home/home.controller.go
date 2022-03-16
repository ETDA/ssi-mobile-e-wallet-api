package home

import (
	"net/http"

	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/requests"
	"gitlab.finema.co/finema/etda/mobile-app-api/services"
	"gitlab.finema.co/finema/etda/mobile-app-api/views"
	core "ssi-gitlab.teda.th/ssi/core"
	"ssi-gitlab.teda.th/ssi/core/errmsgs"
	"ssi-gitlab.teda.th/ssi/core/utils"
)

type HomeController struct{}

func (n *HomeController) Get(c core.IHTTPContext) error {
	return c.JSON(http.StatusOK, core.Map{
		"message": "Hello, I'm Mobile API",
	})
}
func (n *HomeController) Status(c core.IHTTPContext) error {
	return c.JSON(http.StatusOK, core.Map{
		"status": "OK",
	})
}
func (n *HomeController) GetDIDAddress(c core.IHTTPContext) error {
	configSvc := services.NewConfigDIDService(c)
	did, ierr := configSvc.GetDIDAddress()
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusCreated, core.Map{
		"did_address": did.DIDAddress,
	})
}

func (n *HomeController) Pagination(c core.IHTTPContext) error {
	userSvc := services.NewUserService(c)

	items, pageResponse, ierr := userSvc.Pagination(c.QueryParam("ids"), c.QueryParam("ignored_ids"), c.GetPageOptions())
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	return c.JSON(http.StatusOK, core.NewPagination(items, pageResponse))
}

func (n *HomeController) Register(c core.IHTTPContext) error {
	input := &requests.UserRegister{}
	if err := c.BindWithValidate(input); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}
	userSvc := services.NewUserService(c)
	// identityProofingSvc := services.NewIdentityProofingService(c)
	// ekycSvc := services.NewEKYCService(c, userSvc, identityProofingSvc)
	// payload := &services.EKYCVerifyPayload{}
	// _ = utils.Copy(payload, input)
	// _, ierr := ekycSvc.Create(payload)

	// if ierr != nil {
	// 	return c.JSON(ierr.GetStatus(), ierr.JSON())
	// }
	otpService := services.NewOTPService(c)
	user, ierr := userSvc.FindByIDCardNumber(utils.GetString(input.IDCardNo))
	if ierr == nil {
		_, ierr = otpService.Create(&services.OTPCreatePayload{
			UserID:             user.ID,
			IsAlreadyHasIDCard: true,
		})
		return c.JSON(http.StatusOK, views.NewUser(user))
	}

	user, ierr = userSvc.Create(&services.UserCreatePayload{
		IDCardNo:  utils.GetString(input.IDCardNo),
		FirstName: utils.GetString(input.FirstName),
		LastName:  utils.GetString(input.LastName),
		Email:     utils.GetString(input.Email),
	})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	deviceSvc := services.NewDeviceService(c)
	_, ierr = deviceSvc.Create(&services.DeviceCreatePayload{
		UserID:    user.ID,
		Name:      utils.GetString(input.Device.Name),
		OS:        utils.GetString(input.Device.OS),
		OSVersion: utils.GetString(input.Device.OSVersion),
		Model:     utils.GetString(input.Device.Model),
		UUID:      utils.GetString(input.Device.UUID),
	})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	_, ierr = otpService.Create(&services.OTPCreatePayload{
		UserID: user.ID,
	})
	return c.JSON(http.StatusCreated, views.NewUser(user))

}

// func (n *HomeController) RegisterConfirmation(c core.IHTTPContext) error {
// 	input := &requests.UserRegisterConfirmation{}
// 	if err := c.BindWithValidate(input); err != nil {
// 		return c.JSON(err.GetStatus(), err.JSON())
// 	}
//
// 	return c.JSON(http.StatusOK,)
// }
func (n *HomeController) UpdateUserDID(c core.IHTTPContext) error {
	input := &requests.UpdateUserDID{}
	if err := c.BindWithValidateMessage(input); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}
	userSvc := services.NewUserService(c)
	payload := &services.UserUpdatePayload{
		ID:         utils.GetString(input.ID),
		DIDAddress: utils.GetString(input.DIDAddress),
	}

	if input.Device != nil {

		payload.NewDevice = &services.DeviceCreatePayload{
			UserID:    utils.GetString(input.ID),
			Name:      utils.GetString(input.Device.Name),
			OS:        utils.GetString(input.Device.OS),
			OSVersion: utils.GetString(input.Device.OSVersion),
			Model:     utils.GetString(input.Device.Model),
			UUID:      utils.GetString(input.Device.UUID),
		}
	}
	user, ierr := userSvc.UpdateDID(payload)
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	return c.JSON(http.StatusOK, views.NewUser(user))
}

func (n *HomeController) FindUser(c core.IHTTPContext) error {
	deviceSvc := services.NewDeviceService(c)
	userSvc := services.NewUserService(c)

	user, ierr := userSvc.FindByDID(c.Param("did"))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	device, ierr := deviceSvc.FindByUserID(user.ID)
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, views.NewUserDevice(user, device))
}

func (n *HomeController) FindUserByDID(c core.IHTTPContext) error {
	userSvc := services.NewUserService(c)

	user, ierr := userSvc.FindByDID(c.Param("did"))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, views.UserList{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		DIDAddress: user.DIDAddress,
	})
}

func (n *HomeController) UpdateDeviceToken(c core.IHTTPContext) error {
	input := &requests.DeviceTokenUpdate{}
	if err := c.BindWithValidate(input); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	deviceSvc := services.NewDeviceService(c)
	payload := &services.DeviceTokenUpdatePayload{}
	_ = utils.Copy(payload, input)
	ierr := deviceSvc.UpdateToken(payload)
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, core.Map{
		"result": "success",
	})
}

func (n *HomeController) UpdateRequest(c core.IHTTPContext) error {
	input := &requests.RequestUpdate{}
	if err := c.BindWithValidateMessage(input); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	service := services.NewRequestService(c)
	_, ierr := service.Update(c.Param("request_id"), &services.RequestUpdatePayload{
		Status: utils.GetString(input.Status),
	})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, core.Map{
		"result": "success",
	})
}

func (n *HomeController) RequestWaitPagination(c core.IHTTPContext) error {
	service := services.NewRequestService(c)
	items, pageRes, ierr := service.PaginationByDID(c.Param("did"), c.GetPageOptions(),
		&services.RequestPaginationOptions{
			Status:         consts.RequestStatusUnsigned,
			CredentialType: consts.CredentialTypeVC,
		})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, core.NewPagination(views.NewRequests(items), pageRes))
}

func (n *HomeController) RequestWaitCount(c core.IHTTPContext) error {
	service := services.NewRequestService(c)
	count, ierr := service.CountByDID(c.Param("did"),
		&services.RequestCountOptions{
			RequestPaginationOptions: services.RequestPaginationOptions{
				Status:         consts.RequestStatusUnsigned,
				CredentialType: consts.CredentialTypeVC,
			},
		})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, core.Map{
		"count": count,
	})
}

func (n *HomeController) FindRequest(c core.IHTTPContext) error {
	service := services.NewRequestService(c)
	item, ierr := service.Find(c.Param("request_id"))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, item)
}

func (n *HomeController) EKYCVerify(c core.IHTTPContext) error {
	input := &requests.EKYCVerify{}
	if err := c.BindWithValidate(input); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	userService := services.NewUserService(c)
	identityProofingService := services.NewIdentityProofingService(c)

	service := services.NewEKYCService(c, userService, identityProofingService)
	payload := &services.EKYCVerifyPayload{}
	_ = utils.Copy(payload, input)
	item, ierr := service.Verify(payload)

	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, item)
}

func (n *HomeController) SendNotificationByDID(c core.IHTTPContext) error {
	input := &requests.NotificationDID{}
	if ierr := c.BindWithValidate(input); ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	userSvc := services.NewUserService(c)
	user, ierr := userSvc.FindByDID(utils.GetString(input.DIDAddress))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	deviceService := services.NewDeviceService(c)
	device, ierr := deviceService.FindByUserID(user.ID)
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	notifications := make([]services.NotificationTokenItem, 0)
	err := utils.Copy(&notifications, input.Notifications)
	if err != nil {
		c.NewError(err, errmsgs.InternalServerError)
		return c.JSON(errmsgs.InternalServerError.GetStatus(), errmsgs.InternalServerError.JSON())
	}

	notiSvc := services.NewNotificationService(c)
	sendItems := make([]services.SendTokenNotificationItem, 0)
	for _, item := range notifications {
		sendItems = append(sendItems, services.SendTokenNotificationItem{
			NotificationItem: item.NotificationItem,
			Token:            device.Token,
		})
	}

	ierr = notiSvc.SendByTokens(sendItems)
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.NoContent(http.StatusNoContent)
}
func (n *HomeController) RecoveryDID(c core.IHTTPContext) error {
	input := &requests.ResetDevice{}
	if ierr := c.BindWithValidate(input); ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	userID := c.Param("id")
	newKey := &services.NewKeyPayload{}
	newDevice := &services.DeviceCreatePayload{
		UserID:    userID,
		Name:      utils.GetString(input.Device.Name),
		OS:        utils.GetString(input.Device.OS),
		OSVersion: utils.GetString(input.Device.OSVersion),
		Model:     utils.GetString(input.Device.Model),
		UUID:      utils.GetString(input.Device.UUID),
	}
	utils.Copy(newKey, input.NewKey)
	deviceSvc := services.NewDeviceService(c)
	_, ierr := deviceSvc.ResetDevice(
		newKey,
		newDevice,
		userID,
	)
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	userSvc := services.NewUserService(c)
	user, ierr := userSvc.FindByID(c.Param("id"))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	return c.JSON(http.StatusOK, views.NewUser(user))
}

func (n *HomeController) ConfirmOTP(c core.IHTTPContext) error {
	input := &requests.ConfirmOTP{}
	if ierr := c.BindWithValidate(input); ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	otpService := services.NewOTPService(c)
	ierr := otpService.Verify(&services.OTPVerifyPayload{
		UserID:    c.Param("id"),
		OTPNumber: utils.GetString(input.OTPNumber),
	})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	return c.JSON(http.StatusOK, core.Map{"result": "success"})
}
func (n *HomeController) ResendOTP(c core.IHTTPContext) error {
	otpService := services.NewOTPService(c)
	_, ierr := otpService.Create(&services.OTPCreatePayload{
		UserID: c.Param("id"),
	})
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	return c.JSON(http.StatusOK, core.Map{"result": "success"})
}

func (n *HomeController) GetOTP(c core.IHTTPContext) error {
	otpService := services.NewOTPService(c)
	otps, ierr := otpService.FindByUserID(c.Param("id"))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	return c.JSON(http.StatusOK, otps)
}

func (n *HomeController) Verify(c core.IHTTPContext) error {
	input := &requests.Verify{}
	if ierr := c.BindWithValidateMessage(input); ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}
	svc := services.NewVerifyService(c)
	res, ierr := svc.Create(utils.GetString(input.JWT))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusCreated, res)
}

func (n *HomeController) GetVerify(c core.IHTTPContext) error {
	svc := services.NewVerifyService(c)
	jwt, ierr := svc.Get(c.Param("id"))
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, core.Map{
		"jwt": jwt,
	})
}
