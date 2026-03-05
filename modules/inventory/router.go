package inventory

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers inventory routes
func RegisterRoutes(router fiber.Router, handler InventoryHandler) {
	inventory := router.Group("/inventory")

	inventory.Post("/", handler.Create)
	inventory.Get("/", handler.GetAll)
	inventory.Get("/:id", handler.GetByID)
	inventory.Put("/:id", handler.Update)
	inventory.Delete("/:id", handler.Delete)
}
