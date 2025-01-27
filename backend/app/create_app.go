package server

import (
	"context"
	"test_task/app/controller"
	"test_task/config"

	"github.com/gofiber/fiber"
)

func Serve(ctx context.Context, dbConfig config.DBConfig) *fiber.App {
	app := fiber.New()
	controllers := controller.New(ctx, dbConfig)

	app.Get("/list", controllers.ListNews)
	app.Post("/edit/:Id", controllers.PostNews)

	return app
}
