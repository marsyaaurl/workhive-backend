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

	// repositories
	employeeRepo := repository.NewEmployeeRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	sprintRepo := repository.NewSprintRepository(db)

	// services
	authService := service.NewAuthService(employeeRepo)
	projectService := service.NewProjectService(projectRepo)
	sprintService := service.NewSprintService(sprintRepo)

	// controllers
	authController := controller.NewAuthController(authService)
	projectController := controller.NewProjectController(projectService)
	sprintController := controller.NewSprintController(sprintService)

	// public routes
	app.Post("/signup", authController.Signup)
	app.Post("/login", authController.Login)

	// protected routes (pakai JWT)
	protected := app.Group("/", middleware.JWTProtected())

	protected.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to your profile!"})
	})

	// projects routes
	protected.Get("/projects", projectController.FindAll)
	protected.Get("/projects/:project_id", projectController.FindByID)
	protected.Post("/add-project", projectController.Create)
	protected.Put("/projects/:project_id", projectController.Update)
	protected.Delete("/projects/:project_id", projectController.Delete)

	// sprints routes
	protected.Get("/projects/:project_id/sprints", sprintController.FindAll) // list sprints
	protected.Get("/sprints/:sprint_id", sprintController.FindByID)          // detail sprint (singular)
	protected.Post("/projects/:project_id/sprints", sprintController.Create)
	protected.Put("/sprints/:sprint_id", sprintController.Update) // singular
	protected.Delete("/sprints/:sprint_id", sprintController.Delete)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
