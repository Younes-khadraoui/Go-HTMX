package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloMessage(c echo.Context) error {
	return c.HTML(http.StatusOK, `<p class="text-purple-700">Hello from HTMX!</p>`)
}
