package main

import "fmt"

func PluginRegister() map[string]string {
	return map[string]string{
		"name":    "example-plugin",
		"version": "0.0.1",
	}
}

func PluginInit() {
	fmt.Println("Hello, World! from plugin 4")
}
