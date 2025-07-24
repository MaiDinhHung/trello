package controllers

import (
	"fmt"
	"trello/database"
	"trello/models"

	"github.com/gofiber/fiber/v2"
)

// Tạo bình luận mới
func CreateComment(c *fiber.Ctx) error {
	cardID := c.Params("card_id")

	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Không thể parse dữ liệu",
		})
	}

	comment.CardID = parseUint(cardID)

	if err := database.DB.Create(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể tạo bình luận",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

// Lấy danh sách bình luận theo card
func GetComments(c *fiber.Ctx) error {
	cardID := c.Params("card_id")
	var comments []models.Comment
	// database.DB.Preload("User").Where("card_id = ?", cardID).Find(&comments)
	if err := database.DB.
		Preload("User").
		Where("card_id = ?", cardID).
		Order("created_at asc").
		Find(&comments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể lấy bình luận",
		})
	}

	return c.JSON(comments)
}

func UpdateComment(c *fiber.Ctx) error {
	commentID := c.Params("ID")
	cardID := parseUint(c.Params("card_id"))

	var updated models.Comment
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Không thể parse dữ liệu",
		})
	}

	var comment models.Comment
	if err := database.DB.
		Where("id = ? AND card_id = ?", commentID, cardID).
		First(&comment).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Không tìm thấy bình luận",
		})
	}

	comment.Content = updated.Content
	if err := database.DB.Save(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể cập nhật bình luận",
		})
	}

	return c.JSON(comment)
}

func DeleteComment(c *fiber.Ctx) error {
	commentID := c.Params("ID")
	cardID := parseUint(c.Params("card_id"))

	var comment models.Comment
	if err := database.DB.
		Where("id = ? AND card_id = ?", commentID, cardID).
		First(&comment).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Không tìm thấy bình luận",
		})
	}

	if err := database.DB.Delete(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể xoá bình luận",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Helper: parse uint
func parseUint(s string) uint {
	var id uint
	fmt.Sscanf(s, "%d", &id)
	return id
}
