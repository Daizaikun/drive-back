package config

import (
	"github.com/Daizaikun/drive-back/app/config/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	/* "github.com/gofiber/swagger" */)

func Run() *fiber.App {

	Inits()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// app.Use(recover.New())

	/* app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})) */

	// swagger docs
	/* app.Get("/swagger/*", swagger.HandlerDefault) */

	// Crear routs de la aplicación

	routers.List(app)

	return app

}
