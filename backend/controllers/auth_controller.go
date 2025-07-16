package controllers

import (
	"github.com/gofiber/fiber/v2"
	"trello/database"
	"trello/models"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dữ liệu không hợp lệ"})
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Lỗi tạo tài khoản"})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dữ liệu không hợp lệ"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy tài khoản"})
	}

	if input.Password != user.Password {
		return c.Status(401).JSON(fiber.Map{"error": "Sai mật khẩu"})
	}

    return c.JSON(fiber.Map{
		"token": user.ID,
		"user":  user,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to load users"})
	}
	return c.JSON(users)
}
