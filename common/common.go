package common

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type PluginInfo struct {
	Name    string
	Version string
}

type PluginSdk struct {
	AddAction    func(name string, callback func(args ...interface{}) interface{})
	RemoveAction func(name string)
}

func PluginSdkRouteParams(args []interface{}) (*fiber.App, error) {
	if app, ok := args[0].(*fiber.App); ok {
		return app, nil
	}
	return nil, fmt.Errorf("PluginSdkRouteParams args could not be converted")
}
