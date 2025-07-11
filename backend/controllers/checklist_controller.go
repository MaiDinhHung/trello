package controllers

import (
	"trello/database"
	"trello/models"
	"github.com/gofiber/fiber/v2"
)

func GetChecklistItems(c *fiber.Ctx) error {
	cardID := c.Params("card_id")
	var items []models.ChecklistItem
	if err := database.DB.Where("card_id = ?", cardID).Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể lấy checklist"})
	}
	return c.JSON(items)
}

func CreateChecklistItem(c *fiber.Ctx) error {
	cardID := c.Params("card_id")
	var item models.ChecklistItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dữ liệu không hợp lệ"})
	}
	item.CardID = parseUint(cardID)
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể tạo checklist"})
	}
	return c.JSON(item)
}

func UpdateChecklistItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.ChecklistItem
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dữ liệu không hợp lệ"})
	}
	if err := database.DB.Save(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể cập nhật"})
	}
	return c.JSON(item)
}

func DeleteChecklistItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.ChecklistItem{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể xoá"})
	}
	return c.SendStatus(204)
}