package server

import (
	"net/http"

	"github.com/adenix/openapi-boilerplate/pkg/api"
	"github.com/labstack/echo/v4"
)

// GetSpec returns an OpenAPI v3 specification
func GetSpec(ctx echo.Context) error {
	spec, err := api.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, spec)
}
