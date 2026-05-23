package home

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router         fiber.Router
	customLoggerZS *zerolog.Logger
}

func NewHandler(router fiber.Router, loggerZS *zerolog.Logger) *HomeHandler {
	h := &HomeHandler{
		router:         router,
		customLoggerZS: loggerZS,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
	return h
}

func (h *HomeHandler) home(c fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}

func (h *HomeHandler) error(c fiber.Ctx) error {
	h.customLoggerZS.Info().Bool("IsAdmin", true).Str("email", "a@.ru").Int("id", 10).Msg("Инфо")
	return c.SendString("Error")
}
