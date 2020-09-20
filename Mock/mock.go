package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type InformacionResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   string `json:"expiresIn"`
}

func main() {
	app := fiber.New()

	app.Post("/obtener/token", func(c *fiber.Ctx) error {

		fmt.Printf("username %v\n", c.Query("username"))
		fmt.Printf("password %v\n", c.Query("password"))

		data := InformacionResponse{
			AccessToken: "jrlkejrlfjeofjere9967HWFyaW8iOjEsInVzZXJfbmFtZSIJYJKHNKJHJKA878787827822W4iLCJziLCJzY29wZSI6WyJyJhbGciOiJSUzI1NiIsInR5c",
			TokenType:   "bearer",
			ExpiresIn:   "3600",
		}

		return c.JSON(data)

	})

	app.Listen(":8089")
}
