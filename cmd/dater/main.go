package main

import (
	"log"
	"test_junior/internal/config"

	"github.com/labstack/echo/v4"
	"test_junior/internal/rest/middlewares"
	"test_junior/internal/rest/routes"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	e := echo.New()
	middlewares.InitMiddlewares(e, cfg)
	routes.InitCoreRoutes(e, cfg)
	routes.InitApiRoutes(e)
	e.Logger.Fatal(e.Start(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
}
