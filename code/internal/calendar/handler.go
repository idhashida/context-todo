package calendar

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type CalendarHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewCalendarHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &CalendarHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/calendar", h.getCalendar)
}

func (h *CalendarHandler) getCalendar(c *fiber.Ctx) error {
	component := views.Calendar()
	return tadapter.Render(c, component)
}
