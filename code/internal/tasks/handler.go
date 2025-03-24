package tasks

import (
	"go/context-todo/internal/priority"
	"go/context-todo/internal/statuses"
	"go/context-todo/internal/sublists"
	"go/context-todo/pkg/tadapter"
	"go/context-todo/pkg/validator"
	"go/context-todo/views"
	"go/context-todo/views/components"
	"go/context-todo/views/widgets"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type TasksHandler struct {
	router             fiber.Router
	customLogger       *zerolog.Logger
	store              *session.Store
	repositoryTasks    *TasksRepository
	repositorySublists *sublists.SublistsRepository
	repositoryStatuses *statuses.StatusesRepository
	repositoryPriority *priority.PriorityRepository
}

func NewTasksHandler(
	router fiber.Router,
	customLogger *zerolog.Logger,
	store *session.Store,
	repositoryTasks *TasksRepository,
	repositorySublists *sublists.SublistsRepository,
	repositoryStatuses *statuses.StatusesRepository,
	repositoryPriority *priority.PriorityRepository,
) {
	h := &TasksHandler{
		router:             router,
		customLogger:       customLogger,
		store:              store,
		repositoryTasks:    repositoryTasks,
		repositorySublists: repositorySublists,
		repositoryStatuses: repositoryStatuses,
		repositoryPriority: repositoryPriority,
	}
	h.router.Get("/tasks", h.getTasks)
	h.router.Get("/tasks/new", h.getTaskForm)
	h.router.Get("/tasks/update/form/:id", h.getUpdateTaskForm)

	h.router.Post("/tasks/create", h.createNewTask)

	h.router.Patch("/tasks/update/:id", h.updateCurrentTask)
}

func (h TasksHandler) getTasks(c *fiber.Ctx) error {
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
	lists, err := h.repositoryTasks.GetAllListsForUser(userId)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Tasks(lists)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h TasksHandler) getTaskForm(c *fiber.Ctx) error {
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
	lists, err := h.repositoryTasks.GetAllListsForUser(userId)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	sublists, err := h.repositorySublists.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	statuses, err := h.repositoryStatuses.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	priorities, err := h.repositoryPriority.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	component := views.NewTask(lists, sublists, statuses, priorities)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h TasksHandler) getUpdateTaskForm(c *fiber.Ctx) error {
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
	r, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	taskInfo, err := h.repositoryTasks.getInfoTask(r)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	liId := strconv.Itoa(taskInfo[0].ListId)
	stId := strconv.Itoa(taskInfo[0].StatusId)
	dd := taskInfo[0].Deadline.Format("2006-01-02")
	sbId := strconv.Itoa(taskInfo[0].SublistId)
	prId := strconv.Itoa(taskInfo[0].PriorityId)
	var taskInfoMap = map[string]string{
		"task_id":     c.Params("id"),
		"list_id":     liId,
		"title":       taskInfo[0].Title,
		"desc":        taskInfo[0].Desc,
		"context":     taskInfo[0].Context,
		"scenario":    taskInfo[0].Scenario,
		"criterion":   taskInfo[0].Criterion,
		"status_id":   stId,
		"deadline":    dd,
		"sublist_id":  sbId,
		"priority_id": prId,
	}
	lists, err := h.repositoryTasks.GetAllListsForUser(userId)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	sublists, err := h.repositorySublists.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	statuses, err := h.repositoryStatuses.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	priorities, err := h.repositoryPriority.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	component := widgets.EditTaskForm(taskInfoMap, lists, sublists, statuses, priorities)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h TasksHandler) createNewTask(c *fiber.Ctx) error {
	list_id, err := strconv.Atoi(c.FormValue("list_id"))
	if err != nil {
		panic(err)
	}
	status_id, err := strconv.Atoi(c.FormValue("status_id"))
	if err != nil {
		panic(err)
	}
	sublist_id, err := strconv.Atoi(c.FormValue("sublist_id"))
	if err != nil {
		panic(err)
	}
	priority_id, err := strconv.Atoi(c.FormValue("priority_id"))
	if err != nil {
		panic(err)
	}
	deadline, err := time.Parse("2006-01-02", c.FormValue("deadline"))
	if err != nil {
		panic(err)
	}
	form := TaskForm{
		ListId:     list_id,
		Title:      c.FormValue("title"),
		Desc:       c.FormValue("desc"),
		Context:    c.FormValue("context"),
		Scenario:   c.FormValue("scenario"),
		Criterion:  c.FormValue("criterion"),
		StatusId:   status_id,
		Deadline:   deadline,
		SublistId:  sublist_id,
		PriorityId: priority_id,
	}
	errors := validate.Validate(
		// &validators.IntIsPresent{Name: "ListId", Field: form.ListId, Message: "list is missing"},
		&validators.StringIsPresent{Name: "Title", Field: form.Title, Message: "title is missing"},
		// &validators.StringIsPresent{Name: "Desc", Field: form.Desc, Message: "description is missing"},
		// &validators.StringIsPresent{Name: "Context", Field: form.Context, Message: "context is missing"},
		// &validators.StringIsPresent{Name: "Scenario", Field: form.Scenario, Message: "scenario is missing"},
		// &validators.StringIsPresent{Name: "Criterion", Field: form.Criterion, Message: "criterion is missing"},
		&validators.IntIsPresent{Name: "StatusId", Field: form.StatusId, Message: "status is missing"},
		// &validators.StringIsPresent{Name: "Deadline", Field: form.Deadline, Message: "deadline is missing"},
		// &validators.StringIsPresent{Name: "SublistId", Field: form.SublistId, Message: "sublist is missing"},
		&validators.IntIsPresent{Name: "PriorityId", Field: form.PriorityId, Message: "priority is missing"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	// тут в БД вставляем новую таску
	// sess, err := h.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
	// var userId int = 0
	// if ident, ok := sess.Get("user_id").(int); ok {
	// 	userId = ident
	// }
	err = h.repositoryTasks.createTask(form)
	if err != nil {
		h.customLogger.Err(err).Msg("DB ERROR")
	}
	// c.Response().Header.Add("Hx-Redirect", "/observe")
	// return c.Redirect("/observe", http.StatusOK)
	component = components.Notification("success!", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h TasksHandler) updateCurrentTask(c *fiber.Ctx) error {
	list_id, err := strconv.Atoi(c.FormValue("list_id"))
	if err != nil {
		panic(err)
	}
	status_id, err := strconv.Atoi(c.FormValue("status_id"))
	if err != nil {
		panic(err)
	}
	sublist_id, err := strconv.Atoi(c.FormValue("sublist_id"))
	if err != nil {
		panic(err)
	}
	priority_id, err := strconv.Atoi(c.FormValue("priority_id"))
	if err != nil {
		panic(err)
	}
	deadline, err := time.Parse("2006-01-02", c.FormValue("deadline"))
	if err != nil {
		panic(err)
	}
	form := TaskForm{
		ListId:     list_id,
		Title:      c.FormValue("title"),
		Desc:       c.FormValue("desc"),
		Context:    c.FormValue("context"),
		Scenario:   c.FormValue("scenario"),
		Criterion:  c.FormValue("criterion"),
		StatusId:   status_id,
		Deadline:   deadline,
		SublistId:  sublist_id,
		PriorityId: priority_id,
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Title", Field: form.Title, Message: "title is missing"},
		&validators.IntIsPresent{Name: "StatusId", Field: form.StatusId, Message: "status is missing"},
		&validators.IntIsPresent{Name: "PriorityId", Field: form.PriorityId, Message: "priority is missing"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	taskId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.customLogger.Err(err)
	}
	err = h.repositoryTasks.patchTask(form, taskId)
	if err != nil {
		panic(err)
	}
	component = components.Notification("success!", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}
