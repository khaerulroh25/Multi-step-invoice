package handlers

import (
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body Request
	c.BodyParser(&body)

	if body.Username == "admin" && body.Password == "admin123" {
		token, _ := utils.GenerateToken(1, "admin")
		return c.JSON(fiber.Map{"token": token})
	}

	if body.Username == "kerani" && body.Password == "kerani123" {
		token, _ := utils.GenerateToken(2, "kerani")
		return c.JSON(fiber.Map{"token": token})
	}

	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}