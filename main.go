// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)
	app.Get("/book", book)
	app.Get("/super", super)
	app.Get("/coffee", coffee)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("I made a ☕ for you!")
}

func book(c *fiber.Ctx) error {
	return c.SendString("I read this book for you 123!")
}

func super(c *fiber.Ctx) error {
	return c.SendString("I read this super!")
}

func coffee(c *fiber.Ctx) error {
	return c.SendString("I made delicious coffee!")
}
