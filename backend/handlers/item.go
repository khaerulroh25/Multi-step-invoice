package handlers

import (
	"backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) error {
	code := c.Query("code")

	var item models.Item

	if err := config.DB.Where("code = ?", code).First(&item).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	return c.JSON(item)
}