package routes

import (
	"github.com/labstack/echo/v4"

	"test_junior/internal/rest/handlers"
)

func InitApiRoutes(e *echo.Echo) {
	e.GET("/days", handlers.GetDaysLeft)
}
