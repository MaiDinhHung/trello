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
func GetBoardMembers(c *fiber.Ctx) error {
	boardId := c.Params("id")
	var members []models.User

	err := database.DB.
		Raw(`
			SELECT DISTINCT users.* FROM users
			JOIN card_members ON card_members.user_id = users.id
			JOIN cards ON cards.id = card_members.card_id
			JOIN lists ON lists.id = cards.list_id
			WHERE lists.board_id = ?
		`, boardId).Scan(&members).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể lấy danh sách thành viên"})
	}

	return c.JSON(members)
}
