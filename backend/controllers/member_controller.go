package controllers

import (
	"strconv"
	"trello/database"
	"trello/models"

	"github.com/gofiber/fiber/v2"
)

func GetCardMembers(c *fiber.Ctx) error {
	cardIdStr := c.Params("card_id")
	cardId, err := strconv.Atoi(cardIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid card_id"})
	}

	var members []models.CardMember
	if err := database.DB.Where("card_id = ?", uint(cardId)).Preload("User").Find(&members).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to load members"})
	}

	return c.JSON(members)
}

// Thêm thành viên vào Card
func AddMemberToCard(c *fiber.Ctx) error {
	cardIdStr := c.Params("card_id")
	cardId, err := strconv.Atoi(cardIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid card_id"})
	}

	type Request struct {
		UserID uint   `json:"UserID"`
		Role   string `json:"role"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	member := models.CardMember{
		CardID: uint(cardId),
		UserID: body.UserID,
		Role:   body.Role,
	}

	if err := database.DB.Create(&member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add member"})
	}

	return c.JSON(member)
}

// Xoá thành viên ra khỏi Card
func RemoveMemberFromCard(c *fiber.Ctx) error {
	cardIdStr := c.Params("card_id")
	userIdStr := c.Params("UserID")

	cardId, err1 := strconv.Atoi(cardIdStr)
	userId, err2 := strconv.Atoi(userIdStr)

	if err1 != nil || err2 != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid parameters"})
	}

	if err := database.DB.Where("card_id = ? AND UserID = ?", cardId, userId).Delete(&models.CardMember{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to remove member"})
	}

	return c.JSON(fiber.Map{"message": "Member removed"})
}
