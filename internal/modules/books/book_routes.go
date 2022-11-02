package books

import (
	"github.com/gandarfh/api-httui/internal/middleware"
	"github.com/gandarfh/api-httui/internal/modules/books/controllers"
	"github.com/gandarfh/api-httui/pkg/permissions"

	"github.com/gofiber/fiber/v2"
)

// All public routes of book module
func PublicRoutes(route fiber.Router) {
	// get list of all books
	route.Get("/books", controllers.GetBooks)

	// get one book by ID
	route.Get("/books/:id", controllers.GetBook)
}

// All private routes of book module
func PrivateRoutes(route fiber.Router) {
	// create a new book
	route.Post("/books",
		middleware.Credentials(permissions.BookCreateCredential),
		controllers.CreateBook)

	// Update especific book
	route.Patch("/books/:id",
		middleware.Credentials(permissions.BookUpdateCredential),
		controllers.UpdateBook)
}
