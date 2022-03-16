package middlewares

import (
	"github.com/labstack/echo/v4"
	"gitlab.finema.co/finema/etda/mobile-app-api/consts"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"gitlab.finema.co/finema/etda/mobile-app-api/services"
	core "ssi-gitlab.teda.th/ssi/core"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(core.IHTTPContext)

		tokenID := cc.Get(consts.ContextKeyTokenID).(string)

		service := services.NewTokenService(cc)
		t, ierr := service.FindByID(tokenID)
		if ierr != nil {
			return c.JSON(ierr.GetStatus(), ierr.JSON())
		}

		if t.Role != consts.TokenAdminRole {
			return c.JSON(emsgs.PermissionDenied.GetStatus(), emsgs.PermissionDenied.JSON())
		}

		return next(c)
	}
}
