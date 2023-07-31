package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var staticContent embed.FS

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//gofiber static files from embed
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticContent),
		PathPrefix: "/static",
		Browse:     true,
	}))

	app.Listen(":3000")
}
