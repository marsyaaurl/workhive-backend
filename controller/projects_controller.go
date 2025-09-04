package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marsyaaurl/workhive-backend/model/request"
	"github.com/marsyaaurl/workhive-backend/service"
)

type ProjectController struct {
	ProjectService service.ProjectService
}

func NewProjectController(projectService service.ProjectService) *ProjectController {
	return &ProjectController{ProjectService: projectService}
}

func (controller *ProjectController) FindAll(ctx *fiber.Ctx) error {
	employeeID := ctx.Query("project_owner")
	if employeeID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "project_owner is required"})
	}

	res, err := controller.ProjectService.FindAll(employeeID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (controller *ProjectController) Create(ctx *fiber.Ctx) error {
	var req request.CreateProjectRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	res, err := controller.ProjectService.Create(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (controller *ProjectController) Update(ctx *fiber.Ctx) error {
	var req request.UpdateProjectRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	res, err := controller.ProjectService.Update(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (controller *ProjectController) FindByID(ctx *fiber.Ctx) error {
	projectID := ctx.Params("project_id")
	employeeID := ctx.Query("project_owner")

	if projectID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "project_id is required"})
	}
	if employeeID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "project_owner is required"})
	}

	res, err := controller.ProjectService.FindByID(projectID, employeeID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (controller *ProjectController) Delete(ctx *fiber.Ctx) error {
	projectID := ctx.Params("project_id")
	if projectID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "project_id is required"})
	}

	err := controller.ProjectService.Delete(projectID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Project deleted successfully"})
}
