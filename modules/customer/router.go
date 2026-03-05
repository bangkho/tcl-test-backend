package customer

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers customer routes
func RegisterRoutes(group fiber.Router, handler CustomerHandler) {
	customers := group.Group("/customers")

	customers.Post("/", handler.Create)
	customers.Get("/", handler.GetAll)
	customers.Get("/:id", handler.GetByID)
	customers.Put("/:id", handler.Update)
	customers.Delete("/:id", handler.Delete)
}
