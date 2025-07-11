package controllers

import (
	"github.com/gofiber/fiber/v2"
	"trello/database"
	"trello/models"
)


func GetBoardByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var board models.Board

	if err := database.DB.First(&board, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Board không tồn tại",
		})
	}

	return c.JSON(board)
}
