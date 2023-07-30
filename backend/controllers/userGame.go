package controllers

import "github.com/gofiber/fiber/v2"

func UserGame(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfully attemted the round",
	})
}
