package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func WelcomePage(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Welcome to ThunderNight",
	}
	return c.Render(http.StatusOK, "welcome", data)
}
