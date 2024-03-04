package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (u *Handler) GetAllUsers(c echo.Context) error {
	service := UserService()

	res, err := service.ListUsers()
	if err != nil {
		return err
	}

	return responder.SuccessResponse(c, http.StatusOK, res)
}

func (u *Handler) CreateUser(c echo.Context) error {
	var body CreateUserInput

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := body.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	service := UserService()
	if err := service.CreateUser(body); err != nil {
		return err
	}

	return responder.SuccessResponse(c, http.StatusCreated, "")
}
