package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marsyaaurl/workhive-backend/model/request"
	"github.com/marsyaaurl/workhive-backend/service"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) Signup(ctx *fiber.Ctx) error {
	var req request.SignupRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	res, err := controller.AuthService.Signup(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	var req request.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	res, err := controller.AuthService.Login(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}
