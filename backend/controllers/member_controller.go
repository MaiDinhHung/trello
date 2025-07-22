package controllers

import (
	"strconv"
	"trello/database"
	"trello/models"

	"github.com/gofiber/fiber/v2"
)

func GetCardMembers(c *fiber.Ctx) error {
	cardId := c.Params("card_id")

	var members []models.CardMember
	if err := database.DB.Preload("User").Where("card_id = ?", cardId).Find(&members).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to load members"})
	}

	return c.JSON(members)
}


func AddMemberToCard(c *fiber.Ctx) error {
	cardId, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid card_id"})
	}

	var body struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var existing models.CardMember
	if err := database.DB.Where("card_id = ? AND user_id = ?", cardId, body.UserID).First(&existing).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Member already exists in this card"})
	}

	member := models.CardMember{
		CardID: uint(cardId),
		UserID: body.UserID,
		Role:   body.Role,
	}

	if err := database.DB.Create(&member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add member"})
	}

	database.DB.Preload("User").First(&member, member.ID)

	return c.JSON(member)
}


func RemoveMemberFromCard(c *fiber.Ctx) error {
	cardId, err1 := strconv.Atoi(c.Params("card_id"))
	userId, err2 := strconv.Atoi(c.Params("user_id"))

	if err1 != nil || err2 != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid parameters"})
	}

	if err := database.DB.Where("card_id = ? AND user_id = ?", cardId, userId).Delete(&models.CardMember{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to remove member"})
	}

	return c.JSON(fiber.Map{"message": "Member removed"})
}