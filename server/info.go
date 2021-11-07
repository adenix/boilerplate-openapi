package server

import (
	"net/http"

	"github.com/adenix/openapi-boilerplate/pkg/api"
	"github.com/labstack/echo/v4"
)

var info interface{}

// init sets up the info object in memeory using the OpenAPI specification
func init() {
	spec, err := api.GetSwagger()
	if err != nil {
		return
	}

	info = struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Version     string `json:"version"`
	}{
		Name:        spec.Info.Title,
		Description: spec.Info.Description,
		Version:     spec.Info.Version,
	}
}

// GetInfo returns information about the service
func (s *server) GetInfo(ctx echo.Context) error {
	if info == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get API specification")
	}

	return ctx.JSON(http.StatusOK, info)
}
