package routes

import (
	"github.com/ardatosun/go-rest-api/services"
	"github.com/gofiber/fiber"
)

const (
	RootUserUrl = "/api/v1/user"
)

func InitUserRoutes(app *fiber.App) {
	app.Get(RootUserUrl, services.FindAll)
	app.Get(RootUserUrl + "/:id", services.FindById)
	app.Post(RootUserUrl, services.Save)
}