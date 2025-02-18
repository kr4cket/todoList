package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"todoList/internal/common/models"
	"todoList/internal/repository/postgres"
)

type TaskRepository interface {
	Get(ctx *fiber.Ctx) ([]models.Task, error)
	Insert(ctx *fiber.Ctx, task models.Task) error
	Update(ctx *fiber.Ctx, id int, task models.Task) error
	Delete(ctx *fiber.Ctx, id int) error
}

type Repository struct {
	TaskRepository
}

func New(conn *pgx.Conn) *Repository {
	return &Repository{
		TaskRepository: postgres.NewTaskRepository(conn),
	}
}
