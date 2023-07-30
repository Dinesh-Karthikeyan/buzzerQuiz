package controllers

import "github.com/gofiber/fiber/v2"

// to present the first three users the question
func processRequest(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"question": question.Question,
		"id":       question.ID,
		"A":        question.A,
		"B":        question.B,
		"C":        question.C,
		"D":        question.D,
	})
}

// to get the first three users
func processBuzzerPress(queue chan *fiber.Ctx, count *int) {
	if *count < 3 {
		c := <-queue
		processRequest(c)
		*count = *count + 1
	}
}
