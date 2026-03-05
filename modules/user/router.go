package user

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers user routes
func RegisterRoutes(router fiber.Router, handler UserHandler) {
	users := router.Group("/users")

	users.Post("/register", handler.Register)
	users.Post("/login", handler.Login)
	users.Get("/", handler.GetAll)
	users.Get("/:id", handler.GetByID)
	users.Put("/:id", handler.Update)
	users.Delete("/:id", handler.Delete)
}
