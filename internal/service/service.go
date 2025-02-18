package service

import (
	"github.com/gofiber/fiber/v2"
	"todoList/internal/common/models"
	"todoList/internal/repository"
	"todoList/internal/service/task"
)

type TaskService interface {
	Get(ctx *fiber.Ctx) ([]models.Task, int, error)
	Create(ctx *fiber.Ctx, task models.Task) (int, error)
	Update(ctx *fiber.Ctx, id int, task models.Task) (int, error)
	Delete(ctx *fiber.Ctx, id int) (int, error)
}

type Service struct {
	TaskService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TaskService: task.NewService(repo.TaskRepository),
	}
}
