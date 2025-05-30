package routes

import (
	"test_junior/internal/config"
	"test_junior/internal/rest/handlers"

	"github.com/labstack/echo/v4"
)

func InitCoreRoutes(e *echo.Echo, cfg *config.Config) {
	h := handlers.NewHandler(cfg)

	e.Static("/docs", "static/swagger-ui")

	e.GET("/", h.Hello)
	e.GET("/docs/openapi.yaml", handlers.GetOpenAPISchema)
	e.GET("/info", h.GetServiceInfo)
	e.GET("/healthcheck", h.Healthcheck)
}
