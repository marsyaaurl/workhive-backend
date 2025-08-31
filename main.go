package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/marsyaaurl/workhive-backend/config"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	defer config.CloseDB()

	app := fiber.New()

	// health check server
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("WorkHive API up 🚀")
	})

	// health check DB (query ringan)
	app.Get("/health/db", func(c *fiber.Ctx) error {
		var one int
		err := config.DB.QueryRow(context.Background(), "SELECT 1").Scan(&one)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"ok": false, "error": err.Error()})
		}
		return c.JSON(fiber.Map{"ok": true, "db": one})
	})

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
