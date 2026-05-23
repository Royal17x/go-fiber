package main

import (
	"log"
	"project/go-fiber/config"
	"project/go-fiber/internal/home"

	"github.com/gofiber/fiber/v3"
	recover2 "github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	app.Use(recover2.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
