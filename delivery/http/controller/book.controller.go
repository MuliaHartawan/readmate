package controller

import (
	"readmate/delivery/http/dto/request"
	"readmate/domain"
	"readmate/domain/entity"
	"readmate/shared/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	domain domain.Domain
}

func NewBookController(domain domain.Domain) BookController {
	return BookController{
		domain: domain,
	}
}

func (input *BookController) FindAll(ctx *fiber.Ctx) error {

	var books []*entity.Book
	books, err := input.domain.Book.GetAll()
	if err != nil {
		resp, statuCode := util.ResponseError(fiber.ErrBadRequest.Code, "Failed To Get all Books")
		return ctx.Status(statuCode).JSON(resp)
	}

	return ctx.Render("resource/views/home", fiber.Map{
		"Books": books,
	})
}

func (input *BookController) FindOne(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		resp, statuCode := util.ResponseError(fiber.ErrBadRequest.Code, "Failed To Get One Book")
		return ctx.Status(statuCode).JSON(resp)
	}
	var books *entity.Book

	books, err = input.domain.Book.GetByID(uint(id))
	if err != nil {
		resp, statuCode := util.ResponseError(fiber.ErrBadRequest.Code, "Failed To Get all Books")
		return ctx.Status(statuCode).JSON(resp)
	}

	return ctx.Render("resource/views/home", fiber.Map{
		"Books": books,
	})
}
func (input *BookController) Create(ctx *fiber.Ctx) error {
	var book *request.RequestBookCreateDTO

	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := util.ResponseError(fiber.ErrBadRequest.Code, "Invalid Request Body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := input.domain.Book.Create(book); err != nil {
		resp, statuCode := util.ResponseError(fiber.ErrBadRequest.Code, "Failed To Create Book")
		return ctx.Status(statuCode).JSON(resp)
	}

	resp, statusCode := util.ResponseSuccess(200, "Book Created Successfully", ctx.Body())
	return ctx.Status(statusCode).JSON(resp)
}
func (input *BookController) Update(ctx *fiber.Ctx) error {
	return nil
}
func (input *BookController) Remove(ctx *fiber.Ctx) error {
	return nil
}
