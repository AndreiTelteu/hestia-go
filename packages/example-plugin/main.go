package main

import (
	"embed"
	"fmt"
	"net/http"

	common "github.com/andreitelteu/hestia-go/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var staticContent embed.FS

func PluginRegister() common.PluginInfo {
	return common.PluginInfo{
		Name:    "example-plugin",
		Version: "0.0.1",
	}
}

func PluginInit(sdk common.PluginSdk) {
	fmt.Println("Hello, World! from plugin 4")

	sdk.AddAction("routes", func(args ...interface{}) interface{} {
		app, err := common.PluginSdkRouteParams(args)
		if err != nil {
			return nil
		}
		app.Get("/plugin", func(c *fiber.Ctx) error {
			ip := c.IP()
			return c.SendString(fmt.Sprintf("hello from plugin2 ! IP=%s", ip))
		})
		return nil
	})

	sdk.AddAction("routes", func(args ...interface{}) interface{} {
		app, err := common.PluginSdkRouteParams(args)
		if err != nil {
			return nil
		}
		app.Use("/", filesystem.New(filesystem.Config{
			Root:       http.FS(staticContent),
			PathPrefix: "/static",
			Browse:     true,
		}))
		return nil
	})
}
