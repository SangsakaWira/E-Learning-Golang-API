package route

import (
	"github.com/gofiber/fiber/v2"
	"kardiaq.id/m/v2/controller"
)

func RouteInit(route *fiber.App) {

	// Example
	route.Get("/example", controller.GetHelloWorld)
	route.Get("/example/product", controller.GetAllProducts)
	route.Get("/example/:value", controller.GetParams)
	route.Get("/example/json", controller.GetJSONExample)

	// Students
	route.Get("/pelanggan", controller.GetPelanggans)
	route.Post("/pelanggan", controller.CreatePelanggan)
	route.Get("/pelanggan/:id", controller.GetPelanggan)
	route.Put("/pelanggan", controller.UpdatePelanggan)
}
