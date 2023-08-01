package handler

import (
	"go-playground/apps/api/app_config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupModelEndpoint(c echo.Context) error {
	model := c.Param("model")

	if model == "gpt-4" || model == "gpt-3.5-turbo" {
		app_config.ChatGPTModel = model
	}

	return c.String(http.StatusOK, "OK")
}
