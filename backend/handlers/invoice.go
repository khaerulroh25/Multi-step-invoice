package handlers

import (
	"backend/config"
	"backend/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateInvoice(c *fiber.Ctx) error {
	type ItemReq struct {
		ItemID   uint `json:"item_id"`
		Quantity int  `json:"quantity"`
	}

	type Request struct {
		SenderName      string    `json:"sender_name"`
		SenderAddress   string    `json:"sender_address"`
		ReceiverName    string    `json:"receiver_name"`
		ReceiverAddress string    `json:"receiver_address"`
		Items           []ItemReq `json:"items"`
	}

	var body Request
	c.BodyParser(&body)

	userID := c.Locals("user_id").(float64)

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		invoice := models.Invoice{
			InvoiceNumber:   fmt.Sprintf("INV-%d", time.Now().Unix()),
			SenderName:      body.SenderName,
			SenderAddress:   body.SenderAddress,
			ReceiverName:    body.ReceiverName,
			ReceiverAddress: body.ReceiverAddress,
			CreatedBy:       uint(userID),
		}

		if err := tx.Create(&invoice).Error; err != nil {
			return err
		}

		total := 0

		for _, itemReq := range body.Items {
			var item models.Item
			if err := tx.First(&item, itemReq.ItemID).Error; err != nil {
				return err
			}

			subtotal := item.Price * itemReq.Quantity
			total += subtotal

			detail := models.InvoiceDetail{
				InvoiceID: invoice.ID,
				ItemID:    item.ID,
				Quantity:  itemReq.Quantity,
				Price:     item.Price,
				Subtotal:  subtotal,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&invoice).Update("total_amount", total).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Invoice created"})
}