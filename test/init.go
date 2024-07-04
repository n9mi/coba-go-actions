package test

import (
	"coba-go-actions/config"
	"coba-go-actions/route"

	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

func init() {
	app = fiber.New()
	viperCfg, err := config.NewViperCfg()
	if err != nil {
		panic(err)
	}

	route.AssingRoute(app, viperCfg)
}
