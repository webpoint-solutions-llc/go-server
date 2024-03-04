package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

func NewHandler() http.Handler {
	r := echo.New()

	r.HTTPErrorHandler = customErrHandler

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}

	r.Use(middleware.CORSWithConfig(corsConfig))
	r.Use(middleware.CSRF())
	r.Pre(middleware.RemoveTrailingSlash())

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server running")
	})

	AppRouter(r.Group("/api/v1"))

	return r
}

func customErrHandler(err error, c echo.Context) {
	var message any
	code := http.StatusInternalServerError
	c.Logger().Error(err)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	}
	if code >= 500 {
		message = "Unexpected Error Occured"
	}
	errRes := ErrorResponse{
		Success: false,
		Message: message,
	}
	if err := c.JSON(code, errRes); err != nil {
		c.Logger().Error(err)
	}
}
