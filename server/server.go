package server

import (
	"github.com/adenix/openapi-boilerplate/pkg/api"
	"github.com/adenix/openapi-boilerplate/third_party"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// server is the implementation of the API server
type server struct{}

// Ensure that we implement the server interface
var _ api.ServerInterface = (*server)(nil)

// Run bootstraps and starts the API server at the specified address
func Run(addr string) error {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/v3/api-docs", GetSpec)
	e.GET("/swagger-ui/*", echo.WrapHandler(third_party.GetSwaggerUIHandler(e)))

	svr := &server{}
	api.RegisterHandlers(e, svr)

	return e.Start(addr)
}
