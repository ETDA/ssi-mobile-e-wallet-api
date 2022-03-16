package home

import (
	"github.com/labstack/echo/v4"
	"gitlab.finema.co/finema/etda/mobile-app-api/middlewares"
	core "ssi-gitlab.teda.th/ssi/core"
)

func NewHomeHTTPHandler(r *echo.Echo) {
	home := &HomeController{}

	r.GET("/", core.WithHTTPContext(home.Get))
	r.GET("/mobile/did_address", core.WithHTTPContext(home.GetDIDAddress))
	r.GET("/mobile/status", core.WithHTTPContext(home.Status), middlewares.IsAuth)
	r.GET("/mobile/users", core.WithHTTPContext(home.Pagination), middlewares.IsAuth)
	r.POST("/mobile/users", core.WithHTTPContext(home.Register))
	r.GET("/mobile/users/did_address/:did", core.WithHTTPContext(home.FindUserByDID), middlewares.IsAuth)
	r.GET("/mobile/users/:did", core.WithHTTPContext(home.FindUser), middlewares.VerifySignatureMiddleware)
	r.PUT("/mobile/users/:id", core.WithHTTPContext(home.UpdateUserDID), middlewares.VerifySignatureMiddleware)
	r.POST("/mobile/users/:id/otp/confirm", core.WithHTTPContext(home.ConfirmOTP))
	r.POST("/mobile/users/:id/otp/resend", core.WithHTTPContext(home.ResendOTP))
	r.POST("/mobile/users/:id/recovery", core.WithHTTPContext(home.RecoveryDID))
	r.PUT("/mobile/devices/:id/token", core.WithHTTPContext(home.UpdateDeviceToken))
	r.POST("/mobile/notification", core.WithHTTPContext(home.SendNotificationByDID), middlewares.IsAuth)
	r.POST("/mobile/verify", core.WithHTTPContext(home.Verify), middlewares.VerifySignatureMiddleware)
	r.GET("/mobile/verify/:id", core.WithHTTPContext(home.GetVerify))

	// r.PUT("mobile/confirm", core.WithHTTPContext(home.))
	// r.POST("/mobile", core.WithHTTPContext(home.RegisterDevice), middlewares.VerifySignatureMiddleware)
	// r.GET("/mobile/:did", core.WithHTTPContext(home.FindDevice))
	//
	// r.POST("/mobile/:did/sign", core.WithHTTPContext(home.CreateRequest), middlewares.VerifySignatureMiddleware)
	// r.GET("/mobile/:did/sign", core.WithHTTPContext(home.RequestWaitPagination), middlewares.VerifySignatureMiddleware)
	// r.GET("/mobile/:did/sign/count", core.WithHTTPContext(home.RequestWaitPagination))
	//
	// r.GET("/mobile/:did/sign/:request_id", core.WithHTTPContext(home.FindRequest), middlewares.VerifySignatureMiddleware)
	// r.PUT("/mobile/:did/sign/:request_id", core.WithHTTPContext(home.UpdateRequest), middlewares.VerifySignatureMiddleware)
	//
	// r.POST("/mobile/ekyc", core.WithHTTPContext(home.EKYCVerify))

}
