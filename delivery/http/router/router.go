package router

import (
	"readmate/delivery/http/controller"
	"readmate/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {

	authController := controller.NewAuthController(domain)
	bookController := controller.NewBookController(domain)

	router := app.Group("/api/v1")

	authRouter := router.Group("auth")

	authRouter.Post("/register", authController.Register)
	authRouter.Post("/login", authController.Register)

	bookRouter := router.Group("book")

	bookRouter.Get("/", bookController.FindAll)
	bookRouter.Get("/:id", bookController.FindOne)
	bookRouter.Post("/", bookController.Create)
	// bookRouter.Put("/:id")
	// bookRouter.Delete("/:id")
}
