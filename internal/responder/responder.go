package responder

import "github.com/labstack/echo/v4"

type response struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

func SuccessResponse(c echo.Context, code int, i interface{}) error {
	res := response{
		Success: true,
		Payload: i,
	}
	return c.JSON(code, res)
}
