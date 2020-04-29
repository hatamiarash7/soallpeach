package main

import (
	"github.com/gofiber/fiber"
	"os"
	"strconv"
)

var index int64

func main() {
	index = 0

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		newValue := c.Body()
		i, _ := strconv.ParseInt(newValue, 10, 64)
		index += i
	})

	app.Get("/count", func(c *fiber.Ctx) {
		c.Send(index)
	})

	app.Listen(getPort("PORT"))
}

func getPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		return "3000"
	}
	return port
}
