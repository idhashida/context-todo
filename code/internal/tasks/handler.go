package tasks

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type TasksHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewTasksHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &TasksHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/tasks", h.getAllTasks)
}

func (h *TasksHandler) getAllTasks(c *fiber.Ctx) error {
	component := views.Tasks()
	return tadapter.Render(c, component)
}
