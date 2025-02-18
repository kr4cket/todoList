package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"todoList/internal/common/models"
	"todoList/internal/service"
)

type MockTaskService struct {
	service.TaskService
	mock.Mock
}

type Mock struct {
	Name           string
	Body           string
	Mock           func()
	ExpectedStatus int
	ExpectedBody   string
	ID             string
}

func (m *MockTaskService) Get(c *fiber.Ctx) ([]models.Task, int, error) {
	args := m.Called(c)
	return args.Get(0).([]models.Task), args.Int(1), args.Error(2)
}

func (m *MockTaskService) Create(c *fiber.Ctx, task models.Task) (int, error) {
	args := m.Called(c, task)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskService) Update(c *fiber.Ctx, id int, task models.Task) (int, error) {
	args := m.Called(c, id, task)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskService) Delete(c *fiber.Ctx, id int) (int, error) {
	args := m.Called(c, id)
	return args.Int(0), args.Error(1)
}
