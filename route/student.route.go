package route

import (
	"github.com/gofiber/fiber/v2"
	"kardiaq.id/m/v2/controller"
)

func RouteInit(route *fiber.App) {
	route.Get("/", controller.GetHelloWorld)
	route.Get("/product", controller.GetAllProducts)
	route.Get("/value/:value", controller.GetParams)
	route.Get("/json", controller.GetJSONExample)
}
