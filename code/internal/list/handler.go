package list

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/pkg/validator"
	"go/context-todo/views/components"
	"go/context-todo/views/layout"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type ListHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *ListRepository
	store        *session.Store
}

func NewListHandler(router fiber.Router, customLogger *zerolog.Logger, repository *ListRepository, store *session.Store) {
	h := &ListHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	h.router.Get("/list/new", h.createListForm)
	h.router.Get("/list/get", h.getCurrList)

	h.router.Post("/list/create", h.createNewList)
}

// возвращает форму для создания списка
func (h *ListHandler) createListForm(c *fiber.Ctx) error {
	component := layout.NewListForm()
	return tadapter.Render(c, component, http.StatusOK)
}

// возвращает новый <li> элемент списка + добавляет в БД
func (h *ListHandler) createNewList(c *fiber.Ctx) error {
	form := ListForm{
		Name:  c.FormValue("name"),
		Color: c.FormValue("color"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Name", Field: form.Name, Message: "name is not correct"},
		&validators.StringIsPresent{Name: "Color", Field: form.Color, Message: "color is not correct"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	var userId int = 0
	if ident, ok := sess.Get("user_id").(int); ok {
		userId = ident
	}
	err = h.repository.createList(form, userId)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification("initial server error", components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	h.customLogger.Info().Msg("list created")
	component = layout.ListLi(form.Name, form.Color)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *ListHandler) getCurrList(c *fiber.Ctx) error {
	component := components.Notification("list", "success")
	return tadapter.Render(c, component, http.StatusOK)
}
