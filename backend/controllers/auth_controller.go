package controllers


import (
	"trello/database"
	"trello/models"
	"trello/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func Register(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dữ liệu không hợp lệ"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Lỗi mã hoá mật khẩu"})
	}
	input.Password = string(hashedPassword)

	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Tạo tài khoản thất bại"})
	}

	return c.JSON(fiber.Map{
		"message": "Đăng ký thành công",
	})
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

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Sai mật khẩu"})
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Không thể tạo token"})
    }

    return c.JSON(fiber.Map{
        "token": token,
        "user": fiber.Map{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
            "avatar": user.Avatar,
        },
    })
}





func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to load users"})
	}
	return c.JSON(users)
}
