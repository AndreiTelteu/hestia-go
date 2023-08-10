package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var staticContent embed.FS

func PluginRegister() map[string]string {
	return map[string]string{
		"name":    "example-plugin",
		"version": "0.0.1",
	}
}

func PluginInit(actions map[string][](func() interface{})) map[string][](func() interface{}) {
	fmt.Println("Hello, World! from plugin 4")

	addAction := func(name string, action func() interface{}) {
		if _, ok := actions[name]; !ok {
			actions[name] = [](func() interface{}){}
		}
		actions[name] = append(actions[name], action)
	}

	addAction("routes", func() interface{} {
		return func(app *fiber.App) {
			app.Get("/plugin", func(c *fiber.Ctx) error {
				ip := c.IP()
				return c.SendString(fmt.Sprintf("hello from plugin2 ! IP=%s", ip))
			})
		}
	})

	addAction("routes", func() interface{} {
		return func(app *fiber.App) {
			app.Use("/", filesystem.New(filesystem.Config{
				Root:       http.FS(staticContent),
				PathPrefix: "/static",
				Browse:     true,
			}))
		}
	})
	return actions
}
