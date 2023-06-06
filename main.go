package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/hello-world", func(c *fiber.Ctx) error {
		return c.JSON("Hello World")
	})

	log.Fatal(app.Listen(":3000"))
}
