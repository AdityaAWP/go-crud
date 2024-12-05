package routes

import (
	"github.com/AdityaAWP/go-crud/controller"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", controller.UserControllerGetAll)
	r.Post("/user", controller.UserControllerCreate)

}
