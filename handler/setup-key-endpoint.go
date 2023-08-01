package handler

import (
	"go-playground/apps/api/app_config"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetupKeyEndpoint(c echo.Context) error {
	apiKey := c.Param("apiKey")

	if len(apiKey) == 51 && strings.HasPrefix(apiKey, "sk-") {
		app_config.ChatGPTApiKey = apiKey
	}

	return c.String(http.StatusOK, "OK")
}
