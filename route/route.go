package route

import (
	"coba-go-actions/model"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func AssingRoute(app *fiber.App, viperCfg *viper.Viper) {
	app.Get("/", func(c *fiber.Ctx) error {
		data := &model.DataResponse{
			Data: "OK",
		}

		return c.Status(fiber.StatusOK).JSON(data)
	})

	app.Get("/foo", func(c *fiber.Ctx) error {
		data := &model.DataResponse{
			Data: viperCfg.GetString("FOO"),
		}

		return c.Status(fiber.StatusOK).JSON(data)
	})
}
