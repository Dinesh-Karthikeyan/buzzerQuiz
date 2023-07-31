package controllers

import (
	"backend/database"
	"backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
	//checking if this is the first login
	if user.Jwt == "" {
		claims := &jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		ss, err := token.SignedString(secret)

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Could not sign token during registration with secret",
			})
		}

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    ss,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie) // prepared and sent a cookie to the newly registered user

		user.Jwt = ss
		database.DB.Save(&user) // updated the user's jwt field in the database
	}

	err = database.DB.Where("team_name=?", data["team_name"]).First(&user).Error
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	var jwtToken = c.Cookies("jwt") // checking for dual login
	if jwtToken != user.Jwt && jwtToken != "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "dual login not permitted",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successfull",
	})
}
