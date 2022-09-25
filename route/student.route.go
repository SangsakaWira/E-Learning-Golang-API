package route

import "github.com/gofiber/fiber/v2"

func RouteInit(route *fiber.App) {

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// GET http://localhost:8080/hello%20world
	route.Get("/value/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	// GET http://localhost:3000/api/user/john
	route.Get("/json", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"nama": "Sangsaka",
		})

	})
}
