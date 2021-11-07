package server

import (
	"fmt"
	"net/http"

	"github.com/adenix/openapi-boilerplate/pkg/api"
	"github.com/labstack/echo/v4"
)

// GetGreeting returns a friendly greeting to the world
func (s *server) GetGreeting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, api.Greeting{
		Message: "Hello, World!",
	})
}

// GetGreetingName returns a friendly greeting to the specified name
func (s *server) GetGreetingName(ctx echo.Context, name string) error {
	return ctx.JSON(http.StatusOK, api.Greeting{
		Message: fmt.Sprintf("Hello, %s!", name),
	})
}
