package main

import (
	"net/http"
	"net/http/httptest"
	"test_junior/internal/config"
	"test_junior/internal/rest/routes"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestApiGetHelloWorld(t *testing.T) {
	e := echo.New()

	cfg := &config.Config{}
	routes.InitCoreRoutes(e, cfg)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
