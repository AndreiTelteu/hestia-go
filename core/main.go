package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"plugin"
	"strings"

	common "github.com/andreitelteu/hestia-go/common"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var staticContent embed.FS

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
			PluginRegister, err := lookUpSymbol[func() common.PluginInfo](p, "PluginRegister")
			if err != nil {
				return err
			}
			info := (*PluginRegister)()
			fmt.Println("Plugin name:", info.Name)
			fmt.Println("Plugin version:", info.Version)

			PluginInit, err := lookUpSymbol[func(common.PluginSdk)](p, "PluginInit")
			if err != nil {
				return err
			}
			sdk := common.PluginSdk{
				AddAction: func(name string, callback func(args ...interface{}) interface{}) {
					action := func() interface{} {
						return struct {
							pluginName string
							hook       func(args ...interface{}) interface{}
						}{
							pluginName: info.Name,
							hook:       callback,
						}
					}
					if _, ok := actions[name]; !ok {
						actions[name] = [](func() interface{}){action}
					} else {
						actions[name] = append(actions[name], action)
					}
				},
				RemoveAction: func(name string) {
					// i don't know how can i do this like in wordpress
					// in wordpress you use a string with the function name and it works
				},
			}
			(*PluginInit)(sdk)
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
				routeHook := hook().(struct {
					pluginName string
					hook       func(args ...interface{}) interface{}
				})
				pluginApp := fiber.New()
				routeHook.hook(pluginApp)
				api.Mount(
					routeHook.pluginName,
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
