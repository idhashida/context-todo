package inbox

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type InboxHandler struct {
	router          fiber.Router
	customLogger    *zerolog.Logger
	repositoryInbox *InboxRepository
	store           *session.Store
}

func NewInboxHandler(router fiber.Router, customLogger *zerolog.Logger, repositoryInbox *InboxRepository, store *session.Store) {
	h := &InboxHandler{
		router:          router,
		customLogger:    customLogger,
		repositoryInbox: repositoryInbox,
		store:           store,
	}
	h.router.Get("/inbox", h.getInbox)
}

func (h *InboxHandler) getInbox(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	var userId int = 0
	if ident, ok := sess.Get("user_id").(int); ok {
		userId = ident
	}
	if userId == 0 {
		return c.SendStatus(500)
	}
	inboxTasks, err := h.repositoryInbox.getAllInboxTasks(userId)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	var tasks = map[int]string {}
	for _, t := range inboxTasks {
		tasks[t.Id] = t.Title
	}
	component := views.Inbox(tasks)
	return tadapter.Render(c, component, http.StatusOK)
}
