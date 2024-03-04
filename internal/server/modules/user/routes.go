package user

import (
	"github.com/labstack/echo/v4"
)

func Routes(r *echo.Group) {
	h := Handler{}

	r.GET("", h.GetAllUsers)
	r.POST("", h.CreateUser)
}
