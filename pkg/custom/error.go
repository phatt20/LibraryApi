package custom

import "github.com/labstack/echo/v4"

type ErrorMessage struct {
	Message string `json:"message"`
}

func Error(pctx echo.Context, StatusCode int, message string) error {
	return pctx.JSON(StatusCode, &ErrorMessage{Message: message})
}
