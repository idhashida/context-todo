package more

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type MoreHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewMoreHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &MoreHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/more", h.getMore)
}

func (h *MoreHandler) getMore(c *fiber.Ctx) error {
	component := views.More()
	return tadapter.Render(c, component, http.StatusOK)
}
