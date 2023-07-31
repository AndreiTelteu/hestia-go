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
	api := fiber.New()
	app.Mount("/api", api)

	//gofiber static files from embed
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticContent),
		PathPrefix: "/static",
		Browse:     true,
	}))

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":80")
}
