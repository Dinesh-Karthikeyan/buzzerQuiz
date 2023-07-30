package controllers

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(data); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "The data which was sent is corrupted",
		})
	}
	var user models.User
	err := database.DB.Where("team_name=?", data["team_name"]).First(&user).Error
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successfull",
	})
}
