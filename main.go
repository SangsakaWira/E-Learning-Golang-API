package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	"kardiaq.id/m/v2/route"
	"kardiaq.id/m/v2/utility"
)

func main() {
	app := fiber.New()
	utility.ConnectDbORM()
	route.RouteInit(app)
	app.Listen(":3000")
}
