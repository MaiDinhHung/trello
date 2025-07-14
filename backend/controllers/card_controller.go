package controllers

import (
	"trello/database"
	"trello/models"

	"github.com/gofiber/fiber/v2"
)

func CreateCard(c *fiber.Ctx) error {
	card := new(models.Card)

	if err := c.BodyParser(card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if card.Title == "" || card.ListID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing title or list_id",
		})
	}

	if err := database.DB.Create(&card).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create card",
		})
	}

	return c.JSON(card)
}

func UpdateCard(c *fiber.Ctx) error {
	id := c.Params("id")
	var card models.Card

	if err := database.DB.First(&card, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Card not found",
		})
	}

	var data struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		ListID      *uint   `json:"list_id"`
		Position    *int    `json:"position"`
		StartDate   *string `json:"start_date"`
		EndDate     *string `json:"end_date"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if data.Title != nil {
		card.Title = *data.Title
	}
	if data.Description != nil {
		card.Description = *data.Description
	}
	if data.ListID != nil {
		card.ListID = *data.ListID
	}
	if data.Position != nil {
		card.Position = *data.Position
	}
	if data.StartDate != nil {
		card.StartDate = *data.StartDate
	}
	if data.EndDate != nil {
		card.EndDate = *data.EndDate
	}

	if err := database.DB.Save(&card).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update card",
		})
	}

	return c.JSON(card)
}




func GetCardsByListID(c *fiber.Ctx) error {
	listID := c.Params("id")
	var cards []models.Card

	if err := database.DB.
		Where("list_id = ?", listID).
		Order("position ASC").
		Find(&cards).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch cards",
		})
	}

	return c.JSON(cards)
}

func DeleteCard(c *fiber.Ctx) error {
	id := c.Params("id")

	var card models.Card
	if err := database.DB.First(&card, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Card not found",
		})
	}

	if err := database.DB.Delete(&card).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete card",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Card deleted successfully",
	})
}

