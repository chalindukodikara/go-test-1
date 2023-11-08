package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New(fiber.Config{
		AppName:                 "choreo-secret-manager",
		ReadTimeout:             time.Second * 5,
		Prefork:                 false,
		DisableKeepalive:        true,
		DisableStartupMessage:   true,
		EnableTrustedProxyCheck: true,
	})

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")
}
