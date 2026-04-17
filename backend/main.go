package main

import (
	"backend/config"
	"backend/handlers"
	"backend/middleware"
	"backend/models"
	"backend/seed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Item{},
		&models.Invoice{},
		&models.InvoiceDetail{},
	)

	seed.SeedItems()

	app := fiber.New()

	app.Post("/api/login", handlers.Login)
	app.Get("/api/items", handlers.GetItems)

	api := app.Group("/api", middleware.Protected())
	api.Post("/invoices", handlers.CreateInvoice)

	log.Fatal(app.Listen(":3000"))
}