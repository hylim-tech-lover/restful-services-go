package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hylim-tech-lover/restful-services-go/database"
	"github.com/hylim-tech-lover/restful-services-go/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Testing kk")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}

func ListFact(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}
