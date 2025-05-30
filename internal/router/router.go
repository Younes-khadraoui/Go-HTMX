package router

import (
	"github.com/Younes-khadraoui/Go-HTMX/internal/api"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", api.WelcomePage)
	e.GET("/hello", api.HelloMessage)
}
