package main

import (
	"log"

	db "bangkho.dev/tcl/test/backend/db"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("error while connecting to database: ", err)
	}

	app := fiber.New()
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "Demo server is working fine",
		})
	})

	err = app.Listen(":8000")
	if err != nil {
		log.Println("error while starting the server: ", err)
	}
}
