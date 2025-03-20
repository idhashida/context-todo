package main

import (
	"go/context-todo/config"
	"go/context-todo/internal/auth"
	"go/context-todo/internal/calendar"
	"go/context-todo/internal/home"
	"go/context-todo/internal/main_page"
	"go/context-todo/internal/tasks"
	"go/context-todo/pkg/database"
	"go/context-todo/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConf := config.NewLogConfig()
	dbConf := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConf)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")

	dbpool := database.CreateDbPool(dbConf, customLogger)
	defer dbpool.Close()

	// Repositories
	authRepo := auth.NewAuthRepository(dbpool, customLogger)

	// Handlers
	home.NewHomeHandler(app, customLogger)
	auth.NewAuthHandler(app, customLogger, authRepo)
	main_page.NewMainPageHandler(app, customLogger)
	tasks.NewTasksHandler(app, customLogger)
	calendar.NewCalendarHandler(app, customLogger)

	app.Listen(":3000")
}
