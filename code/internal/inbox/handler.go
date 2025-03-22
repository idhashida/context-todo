package inbox

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type InboxHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewInboxPageHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &InboxHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/inbox", h.getInbox)
}

func (h *InboxHandler) getInbox(c *fiber.Ctx) error {
	component := views.MainLayout()
	return tadapter.Render(c, component, http.StatusOK)
}
