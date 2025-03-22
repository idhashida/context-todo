package tasks

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"
	"net/http"

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
	h.router.Get("/tasks", h.getTasks)
	h.router.Get("/tasks/new", h.createNewTask)
}

func (h TasksHandler) getTasks(c *fiber.Ctx) error {
	component := views.Tasks()
	return tadapter.Render(c, component, http.StatusOK)
}

func (h TasksHandler) createNewTask(c *fiber.Ctx) error {
	component := views.NewTask()
	return tadapter.Render(c, component, http.StatusOK)
}