package transaction

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers transaction routes
func RegisterRoutes(router fiber.Router, handler TransactionHandler) {
	transactions := router.Group("/transactions")

	transactions.Post("/", handler.Create)
	transactions.Get("/", handler.GetAll)
	transactions.Get("/:id", handler.GetByID)
	transactions.Put("/:id", handler.Update)
	transactions.Delete("/:id", handler.Delete)
}
