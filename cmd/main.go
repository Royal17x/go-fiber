package main

import (
	"project/go-fiber/config"
	"project/go-fiber/internal/home"
	"project/go-fiber/pkg/logger"

	fiberzerolog "github.com/gofiber/contrib/v3/zerolog"
	"github.com/gofiber/fiber/v3"
	recover2 "github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	config.Init()
	_ = config.NewDatabaseConfig()
	loggerConfig := config.NewLoggerConfig()
	customLogger := logger.NewZeroSlogLogger(loggerConfig)

	app := fiber.New()
	app.Use(recover2.New())
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use()

	home.NewHandler(app, customLogger)

	app.Listen(":3000")
}
