package controller

import (
	"github.com/gofiber/fiber/v2"
	model "kardiaq.id/m/v2/models"
	database "kardiaq.id/m/v2/utility"
)

func GetPelanggans(c *fiber.Ctx) error {
	pelanggan := []model.Pelanggan{}
	database.DB.Db.Find(&pelanggan)

	return c.Status(200).JSON(pelanggan)
}

func GetPelanggan(c *fiber.Ctx) error {
	pelanggans := []model.Pelanggan{}
	// if err := pelanggans; err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }
	database.DB.Db.Where("id = ?", c.Params("id")).Find(&pelanggans)
	return c.Status(200).JSON(pelanggans)
}

func CreatePelanggan(c *fiber.Ctx) error {
	return c.SendString("New Pelanggan Created")
}

func UpdatePelanggan(c *fiber.Ctx) error {
	return c.SendString("Update New Pelanggan")
}

func DeletePelanggan(c *fiber.Ctx) error {
	return c.SendString("Delete Pelanggan")
}
