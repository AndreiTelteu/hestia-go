package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PluginRegister() map[string]string {
	return map[string]string{
		"name":    "example-plugin",
		"version": "0.0.1",
	}
}

func PluginInit(actions map[string][](func() interface{})) map[string][](func() interface{}) {
	fmt.Println("Hello, World! from plugin 4")
	// actions := map[string][](func() interface{}){}

	addAction := func(name string, action func() interface{}) {
		if _, ok := actions[name]; !ok {
			actions[name] = [](func() interface{}){}
		}
		actions[name] = append(actions[name], action)
	}

	addAction("routes", func() interface{} {
		return map[string]func(c *fiber.Ctx) error{
			"plugin": func(c *fiber.Ctx) error {
				return c.SendString("hello from plugin !")
			},
		}
	})
	return actions
}
