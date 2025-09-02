package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/marsyaaurl/workhive-backend/config"
	"github.com/marsyaaurl/workhive-backend/controller"
	"github.com/marsyaaurl/workhive-backend/middleware"
	"github.com/marsyaaurl/workhive-backend/repository"
	"github.com/marsyaaurl/workhive-backend/service"
)

func main() {
	_ = godotenv.Load()

	db := config.ConnectDB()
	defer config.CloseDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("WorkHive API up 🚀")
	})

	employeeRepo := repository.NewEmployeeRepository(db)
	authService := service.NewAuthService(employeeRepo)
	authController := controller.NewAuthController(authService)

	// public routes
	app.Post("/signup", authController.Signup)
	app.Post("/login", authController.Login)

	// protected routes
	app.Get("/profile", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to your profile!"})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
