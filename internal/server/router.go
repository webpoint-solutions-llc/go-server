package server

import (
	"github.com/labstack/echo/v4"
)

func AppRouter(r *echo.Group) {
	user.Routes(r.Group("/users"))
}
