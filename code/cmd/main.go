package main

import (
	"go/context-todo/config"
	"go/context-todo/internal/auth"
	"go/context-todo/internal/calendar"
	"go/context-todo/internal/home"
	"go/context-todo/internal/list"
	"go/context-todo/internal/main_page"
	"go/context-todo/internal/more"
	"go/context-todo/internal/priority"
	"go/context-todo/internal/statuses"
	"go/context-todo/internal/sublists"
	"go/context-todo/internal/tasks"
	"go/context-todo/pkg/database"
	"go/context-todo/pkg/logger"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Expiration: 3 * time.Hour,
		Storage:    storage,
	})

	// Repositories
	authRepo := auth.NewAuthRepository(dbpool, customLogger)
	listRepo := list.NewListRepository(dbpool, customLogger)
	tasksRepo := tasks.NewTasksRepository(dbpool, customLogger)
	sublistsRepo := sublists.NewSublistsRepository(dbpool, customLogger)
	statusesRepo := statuses.NewStatusesRepository(dbpool, customLogger)
	priorityRepo := priority.NewPriorityRepository(dbpool, customLogger)

	// Handlers
	home.NewHomeHandler(app, customLogger)
	auth.NewAuthHandler(app, customLogger, authRepo, store)
	main_page.NewMainPageHandler(app, customLogger)
	tasks.NewTasksHandler(app, customLogger, store, tasksRepo, sublistsRepo, statusesRepo, priorityRepo)
	calendar.NewCalendarHandler(app, customLogger)
	more.NewMoreHandler(app, customLogger)
	list.NewListHandler(app, customLogger, listRepo, store)

	app.Listen(":3000")
}
