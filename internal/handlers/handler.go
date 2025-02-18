package handlers

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"todoList/internal/handlers/task"
	"todoList/internal/service"
)

type Handler struct {
	taskHandler *task.Handler
}

func New(services *service.Service) *Handler {
	return &Handler{
		taskHandler: task.NewHandler(services.TaskService),
	}
}

func (h *Handler) Init(app *fiber.App) *fiber.App {

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")
	api.Get("/tasks", h.taskHandler.Get)
	api.Post("/tasks", h.taskHandler.Post)
	api.Put("/tasks/:id", h.taskHandler.Put)
	api.Delete("/tasks/:id", h.taskHandler.Delete)

	return app
}
