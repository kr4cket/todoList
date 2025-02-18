package task

import (
	"github.com/gofiber/fiber/v2"
	"todoList/internal/common/models"
	"todoList/internal/repository"
)

type Service struct {
	repo repository.TaskRepository
}

func NewService(repo repository.TaskRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(ctx *fiber.Ctx) ([]models.Task, int, error) {
	tasks, err := s.repo.Get(ctx)

	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return tasks, fiber.StatusOK, nil
}

func (s *Service) Create(ctx *fiber.Ctx, task models.Task) (int, error) {
	err := s.repo.Insert(ctx, task)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusCreated, nil
}
func (s *Service) Update(ctx *fiber.Ctx, id int, task models.Task) (int, error) {
	err := s.repo.Update(ctx, id, task)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusOK, nil
}
func (s *Service) Delete(ctx *fiber.Ctx, id int) (int, error) {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusOK, nil
}
