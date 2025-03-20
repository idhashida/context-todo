package main_page

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type MainPageHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewMainPageHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &MainPageHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/observe", h.getMainLayout)
}

func (h *MainPageHandler) getMainLayout(c *fiber.Ctx) error {
	component := views.MainLayout()
	return tadapter.Render(c, component, http.StatusOK)
}
