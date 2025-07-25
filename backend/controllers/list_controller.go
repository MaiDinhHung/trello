package controllers

import (
	"trello/database"
	"trello/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

func CreateList(c *fiber.Ctx) error {
	var list models.List

	if err := c.BodyParser(&list); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Sai dữ liệu gửi lên", "detail": err.Error()})
	}

	if list.Title == "" || list.BoardID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Thiếu title hoặc board_id"})
	}

	if err := database.DB.Create(&list).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Tạo list thất bại", "detail": err.Error()})
	}

	return c.JSON(list)
}

func GetListsByBoardID(c *fiber.Ctx) error {
	boardID := c.Params("id")

	var lists []models.List

	if err := database.DB.
		Preload("Cards", func(db *gorm.DB) *gorm.DB {
			return db.
				Order("position ASC").
				Preload("Members.User").
				Preload("Comments").
				Preload("ChecklistItems")
		}).
		Where("board_id = ?", boardID).
		Order("position ASC").
		Find(&lists).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không lấy được danh sách cột"})
	}

		return c.JSON(lists)
	}


func UpdateList(c *fiber.Ctx) error {
	id := c.Params("id")
	var list models.List

	if err := database.DB.First(&list, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "List not found",
		})
	}

	var data struct {
		Title    string `json:"title"`
		Position int    `json:"position"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if data.Title != "" {
		list.Title = data.Title
	}

	if data.Position >= 0 {
		list.Position = data.Position
	}

	if err := database.DB.Save(&list).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update list",
		})
	}

	return c.JSON(list)
}

func DeleteList(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DB.Delete(&models.List{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete list",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
