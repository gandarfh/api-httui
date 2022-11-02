package routes

import (
	"github.com/gandarfh/api-httui/internal/modules/books"
	"github.com/gandarfh/api-httui/internal/modules/healthcheck"
	"github.com/gandarfh/api-httui/internal/modules/tokens"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Routes to health check
	healthcheck.PublicRoutes(api)

	// All public routes of book module
	books.PublicRoutes(api)

	// All public routes of tokens module
	tokens.PublicRoutes(api)
}
