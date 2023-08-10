package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var staticContent embed.FS

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
func main() {
	app := fiber.New(fiber.Config{
		ProxyHeader: fiber.HeaderXForwardedFor,
	})
	api := fiber.New()
	app.Mount("/api", api)

	// gofiber static files from embed
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticContent),
		PathPrefix: "/static",
		Browse:     true,
	}))

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	actions := map[string][](func() interface{}){}
	// scan plugins directory for .so files
	err := filepath.Walk("../build", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".so") {
			// load plugin
			fmt.Println("Loading plugin:", path)
			p, err := plugin.Open(path)
			if err != nil {
				return err
			}
			PluginRegister, err := lookUpSymbol[func() map[string]string](p, "PluginRegister")
			if err != nil {
				return err
			}
			info := (*PluginRegister)()
			fmt.Println("Plugin name:", info["name"])
			fmt.Println("Plugin version:", info["version"])

			PluginInit, err := lookUpSymbol[func(map[string][](func() interface{})) map[string][](func() interface{})](p, "PluginInit")
			if err != nil {
				return err
			}
			actions = (*PluginInit)(actions)
			for action, hooks := range actions {
				for hookIndex, hook := range hooks {
					if action == "routes" {
						routeDefinition := hook().(map[string]interface{})
						if _, ok := routeDefinition["handler"]; !ok {
							actions[action] = RemoveIndex(actions[action], hookIndex)
							continue
						}
						if _, ok := routeDefinition["path"]; !ok {
							routeDefinition["path"] = "/"
						}
						if _, ok := routeDefinition["method"]; !ok {
							routeDefinition["method"] = "GET"
						}
						routeDefinition["plugin-prefix"] = info["name"]
						actions[action][hookIndex] = func() interface{} {
							return routeDefinition
						}
					}
				}
			}
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}
	for action, hooks := range actions {
		for _, hook := range hooks {
			if action == "routes" {
				routeDefinition := hook().(map[string]interface{})
				pluginApp := fiber.New()
				pluginApp.Add(
					routeDefinition["method"].(string),
					routeDefinition["path"].(string),
					routeDefinition["handler"].(func(c *fiber.Ctx) error),
				)
				api.Mount(
					routeDefinition["plugin-prefix"].(string),
					pluginApp,
				)
			}
		}
	}

	app.Listen(":80")
}

func lookUpSymbol[M any](plugin *plugin.Plugin, symbolName string) (*M, error) {
	symbol, err := plugin.Lookup(symbolName)
	if err != nil {
		return nil, err
	}
	result := symbol.(M)
	return &result, nil
	// switch symbol.(type) {
	// case *M:
	// 	return symbol.(*M), nil
	// case M:
	// 	result := symbol.(M)
	// 	return &result, nil
	// default:
	// 	return nil, fmt.Errorf("unexpected type from module symbol: %T", symbol)
	// }
}
