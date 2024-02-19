package controller

import (
	"readmate/delivery/http/dto/request"
	"readmate/domain"
	"readmate/shared/util"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	domain domain.Domain
}

func NewAuthController(domain domain.Domain) AuthController {
	return AuthController{
		domain: domain,
	}
}

func (input *AuthController) Register(ctx *fiber.Ctx) error {
	var user request.RequestRegisterDTO
	if err := ctx.BodyParser(&user); err != nil {
		resp, statusCode := util.ResponseError(fiber.ErrBadRequest.Code, "Invalid Request Body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := input.domain.User.Create(user); err != nil {
		resp, statuCode := util.ResponseError(fiber.ErrBadRequest.Code, "Failed To Register")
		return ctx.Status(statuCode).JSON(resp)
	}

	resp, statusCode := util.ResponseSuccess(200, "User Created Successfully", ctx.Body())
	return ctx.Status(statusCode).JSON(resp)
}
