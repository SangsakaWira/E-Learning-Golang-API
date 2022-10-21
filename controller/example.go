package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetHelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetParams(c *fiber.Ctx) error {
	return c.SendString("value: " + c.Params("value"))
	// => Get request with value: hello world
}

func GetJSONExample(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"nama": "Sangsaka",
	})

}
