package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoList/internal/handlers/task"
	"todoList/test/mocks/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T, app *fiber.App, path string, method string, mockService *handler.MockTaskService, function func(*handler.MockTaskService) []handler.Mock) {
	tests := function(mockService)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Mock()

			req := httptest.NewRequest(method, path, nil)
			if tt.Body != "" {
				req = httptest.NewRequest(method, path, strings.NewReader(tt.Body))
				req.Header.Set("Content-Type", "application/json")
			}
			if tt.ID != "" {
				req = httptest.NewRequest(method, path+"/"+tt.ID, strings.NewReader(tt.Body))
				req.Header.Set("Content-Type", "application/json")
			}
			if tt.ID != "" && tt.Body != "" {
				req = httptest.NewRequest(method, path+"/"+tt.ID, strings.NewReader(tt.Body))
				req.Header.Set("Content-Type", "application/json")
			}

			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.ExpectedStatus, resp.StatusCode)

			body := make([]byte, resp.ContentLength)
			resp.Body.Read(body)
			assert.Equal(t, tt.ExpectedBody, string(body))

			mockService.AssertExpectations(t)
		})
	}
}

func TestGet(t *testing.T) {
	mockService := new(handler.MockTaskService)
	handlers := task.NewHandler(mockService)
	method := http.MethodGet
	path := "/tasks"
	app := fiber.New()
	app.Get("/tasks", handlers.Get)

	setupTest(t, app, path, method, mockService, handler.GetMock)
}

func TestPost(t *testing.T) {
	mockService := new(handler.MockTaskService)
	handlers := task.NewHandler(mockService)
	method := http.MethodPost
	path := "/tasks"
	app := fiber.New()
	app.Post("/tasks", handlers.Post)

	setupTest(t, app, path, method, mockService, handler.GetPostMock)
}

func TestPut(t *testing.T) {
	mockService := new(handler.MockTaskService)
	handlers := task.NewHandler(mockService)
	method := http.MethodPut
	path := "/tasks"
	app := fiber.New()
	app.Put("/tasks/:id", handlers.Put)

	setupTest(t, app, path, method, mockService, handler.GetPutMock)
}

func TestDelete(t *testing.T) {
	mockService := new(handler.MockTaskService)
	handlers := task.NewHandler(mockService)
	method := http.MethodDelete
	path := "/tasks"
	app := fiber.New()
	app.Delete("/tasks/:id", handlers.Delete)

	setupTest(t, app, path, method, mockService, handler.GetDeleteMock)
}
