package main

import (
	"log"

	db "bangkho.dev/tcl/test/backend/db"
	"bangkho.dev/tcl/test/backend/db/migration"
	"bangkho.dev/tcl/test/backend/modules/customer"
	"bangkho.dev/tcl/test/backend/modules/inventory"
	"bangkho.dev/tcl/test/backend/modules/transaction"
	"bangkho.dev/tcl/test/backend/modules/user"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("error while connecting to database: ", err)
	}

	// Run migrations
	if err := migration.MigrationInventory(db.GetDB()); err != nil {
		log.Fatal("error running inventory migration: ", err)
	}
	if err := migration.MigrationUser(db.GetDB()); err != nil {
		log.Fatal("error running user migration: ", err)
	}
	if err := migration.MigrationCustomer(db.GetDB()); err != nil {
		log.Fatal("error running customer migration: ", err)
	}
	if err := migration.MigrationTransaction(db.GetDB()); err != nil {
		log.Fatal("error running transaction migration: ", err)
	}

	log.Println("Migrations completed successfully")

	app := fiber.New()

	// Health check
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Demo server is working fine",
		})
	})

	// Initialize repositories
	customerRepository := customer.NewCustomerRepository(db.GetDB())
	transactionRepository := transaction.NewTransactionRepository(db.GetDB())
	inventoryRepository := inventory.NewInventoryRepository(db.GetDB())
	userRepository := user.NewUserRepository(db.GetDB())

	// Initialize services
	customerService := customer.NewCustomerService(customerRepository)
	transactionService := transaction.NewTransactionService(transactionRepository, customerRepository, inventoryRepository)
	inventoryService := inventory.NewInventoryService(inventoryRepository)
	userService := user.NewUserService(userRepository)

	// Initialize handlers
	customerHandler := customer.NewCustomerHandler(customerService)
	transactionHandler := transaction.NewTransactionHandler(transactionService)
	inventoryHandler := inventory.NewInventoryHandler(inventoryService)
	userHandler := user.NewUserHandler(userService)

	// Register routes
	api := app.Group("/api")

	customer.RegisterRoutes(api, customerHandler)
	transaction.RegisterRoutes(api, transactionHandler)
	inventory.RegisterRoutes(api, inventoryHandler)
	user.RegisterRoutes(api, userHandler)

	err = app.Listen(":8000")
	if err != nil {
		log.Println("error while starting the server: ", err)
	}
}
