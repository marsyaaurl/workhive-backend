package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marsyaaurl/workhive-backend/model/request"
	"github.com/marsyaaurl/workhive-backend/service"
)

type SprintController struct {
	SprintService service.SprintService
}

func NewSprintController(sprintService service.SprintService) *SprintController {
	return &SprintController{SprintService: sprintService}
}

func (controller *SprintController) FindAll(ctx *fiber.Ctx) error {
	projectID := ctx.Params("project_id")
	if projectID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "project_id is required"})
	}

	res, err := controller.SprintService.FindAll(projectID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()}) // ✅ Added return
	}
	return ctx.JSON(res)
}

func (controller *SprintController) Create(ctx *fiber.Ctx) error {
	// ✅ Ambil project_id dari URL parameter jika menggunakan route /projects/:project_id/sprints
	projectID := ctx.Params("project_id")

	var req request.CreateSprintRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// ✅ Set project_id dari URL parameter jika ada
	if projectID != "" {
		req.ProjectID = projectID
	}

	res, err := controller.SprintService.Create(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (controller *SprintController) Update(ctx *fiber.Ctx) error {
	// ✅ Ambil sprint_id dari URL parameter
	sprintID := ctx.Params("sprint_id")
	if sprintID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "sprint_id is required"})
	}

	var req request.UpdateSprintRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// ✅ Set sprint_id dari URL parameter
	req.SprintID = sprintID

	res, err := controller.SprintService.Update(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (controller *SprintController) FindByID(ctx *fiber.Ctx) error {
	sprintID := ctx.Params("sprint_id")
	if sprintID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "sprint_id is required"})
	}

	res, err := controller.SprintService.FindByID(sprintID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()}) // ✅ Already has return
	}
	return ctx.JSON(res)
}

func (controller *SprintController) Delete(ctx *fiber.Ctx) error {
	sprintID := ctx.Params("sprint_id")
	if sprintID == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "sprint_id is required"})
	}

	err := controller.SprintService.Delete(sprintID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()}) // ✅ Added return
	}
	return ctx.JSON(fiber.Map{"message": "Sprint deleted successfully"}) // ✅ Fixed message
}
