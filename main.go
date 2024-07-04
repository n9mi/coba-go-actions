package main

import (
	"coba-go-actions/config"
	"coba-go-actions/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	viperCfg, err := config.NewViperCfg()
	if err != nil {
		panic(err)
	}

	route.AssingRoute(app, viperCfg)

	app.Listen(":3000")
}
