package auth

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/pkg/validator"
	"go/context-todo/views"
	"go/context-todo/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type AuthHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewAuthHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &AuthHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/login", h.login)
	h.router.Get("/register", h.register)
	h.router.Post("/auth/login", h.loginPost)
	h.router.Post("/auth/register", h.registerPost)
}

func (h *AuthHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadapter.Render(c, component)
}

func (h *AuthHandler) register(c *fiber.Ctx) error {
	component := views.Register()
	return tadapter.Render(c, component)
}

func (h *AuthHandler) loginPost(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "email is not correct"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password, Message: "password missing"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	// тут должен быть редирект в защищенную часть сайта
	return c.SendString("/login redirect")
}

func (h *AuthHandler) registerPost(c *fiber.Ctx) error {
	form := RegisterForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Username: c.FormValue("username"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Username", Field: form.Username, Message: "username is not correct"},
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "email is not correct"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password, Message: "password is not correct"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	// тут должен быть редирект в защищенную часть сайта
	return c.SendString("/register redirect")
}
