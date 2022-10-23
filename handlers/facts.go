package handlers

import (
	"github.com/Paul-Boersma/go-div-rhino/database"
	"github.com/Paul-Boersma/go-div-rhino/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.DB.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.DB.Create(&fact)
	return c.Status(200).JSON(fact)
}
