package handlers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetOpenAPISchema(c echo.Context) error {
	schema, err := os.ReadFile("static/openapi.yaml")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to read OpenAPI schema")
	}

	return c.Blob(http.StatusOK, "application/yaml", schema)
}
