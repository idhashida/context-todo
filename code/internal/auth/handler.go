package auth

import (
	"go/context-todo/pkg/tadapter"
	"go/context-todo/views"

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
	h.router.Get("/auth/login", h.login)
	h.router.Get("/auth/register", h.register)
}

func (h *AuthHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadapter.Render(c, component)
}

func (h *AuthHandler) register(c *fiber.Ctx) error {
	component := views.Register()
	return tadapter.Render(c, component)
}
