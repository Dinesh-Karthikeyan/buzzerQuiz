package controllers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

var question models.Question
var buzzerQueue = make(chan *fiber.Ctx, 10)
var counter int = 0

// to process and queue the buzzer request
func BuzzerPressHandler(c *fiber.Ctx) error {
	// Add the request to the buzzerQueue
	go processBuzzerPress(buzzerQueue, &counter)
	buzzerQueue <- c

	// Respond to the client, confirming the buzzer press was received
	return c.JSON(fiber.Map{"message": "Buzzer press received and queued", "pressed": "true"})
}

// func UserGame(c *fiber.Ctx) error {
// 	var data map[string]string
// 	if err := c.BodyParser(data); err != nil {
// 		c.Status(fiber.StatusInternalServerError)
// 		return c.JSON(fiber.Map{
// 			"message": "The data which was sent is corrupted",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"A": question.A,
// 		"B": question.B,
// 		"C": question.C,
// 		"D": question.D,
// 	})
// }
