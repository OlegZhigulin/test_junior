package handlers

import (
	"net/http"
	"test_junior/internal/config"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (h *Handler) GetServiceInfo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.cfg.ServiceInformation)
}

func (h *Handler) Healthcheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}
