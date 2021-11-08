package third_party

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed swagger_ui
var swaggerUI embed.FS

// SwaggerUIHandler returns an HTTP handler for the Swagger UI
func SwaggerUIHandler(e *echo.Echo) http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	subFS, err := fs.Sub(swaggerUI, "swagger_ui")
	if err != nil {
		e.Logger.Fatalf("failed to create sub filesystem for Swagger UI: %w", err)
	}
	return http.StripPrefix("/swagger-ui", http.FileServer(http.FS(subFS)))
}
