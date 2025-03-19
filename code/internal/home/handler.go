package home

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/", h.home)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
}
