package task

import (
	"github.com/gofiber/fiber/v2"
	"todoList/internal/common/errors"
	"todoList/internal/common/models"
	"todoList/internal/common/responses"
	"todoList/internal/service"
)

type Handler struct {
	service service.TaskService
}

func NewHandler(service service.TaskService) *Handler {
	return &Handler{
		service: service,
	}
}

// Get @Summary Get tasks
// @Description Get all tasks from storage
// @ID get-tasks
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Task
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/tasks [get]
func (h *Handler) Get(c *fiber.Ctx) error {
	tasks, code, err := h.service.Get(c)

	if err != nil {
		c.Status(code)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while getting Tasks",
			Message: err.Error(),
		})
	}

	c.Status(code)
	return c.JSON(tasks)
}

// Post @Summary Create task
// @Description Creates task
// @ID create-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param input body models.Task true "Task"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/tasks/ [post]
func (h *Handler) Post(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while parsing Task",
			Message: err.Error(),
		})
	}

	err := models.ValidateCreateTask(*task)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while validating Task",
			Message: err.Error(),
		})
	}

	code, err := h.service.Create(c, *task)
	if err != nil {
		c.Status(code)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while creating Task",
			Message: err.Error(),
		})
	}

	c.Status(code)
	return c.JSON(responses.SuccessResponse{
		Message:     "Success",
		Description: "Task Created Successfully",
	})
}

// Put @Summary Update task
// @Description Update task by id
// @ID update-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task Id"
// @Param input body models.Task true "Task"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/tasks/{id} [put]
func (h *Handler) Put(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while parsing id",
			Message: err.Error(),
		})
	}

	task := new(models.Task)

	if err = c.BodyParser(task); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while parsing Task",
			Message: err.Error(),
		})
	}

	err = models.ValidateUpdateTask(*task)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while validating Task",
			Message: err.Error(),
		})
	}

	code, err := h.service.Update(c, id, *task)
	if err != nil {
		c.Status(code)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while updating Task",
			Message: err.Error(),
		})
	}
	return c.JSON(responses.SuccessResponse{
		Message:     "Success",
		Description: "Task Updated Successfully",
	})
}

// Delete @Summary Delete task
// @Description Deletes tasks by id
// @ID delete-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task Id"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/tasks/{id} [delete]
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while parsing id",
			Message: err.Error(),
		})
	}

	code, err := h.service.Delete(c, id)
	if err != nil {
		c.Status(code)
		return c.JSON(errors.ErrorResponse{
			Error:   "Error while deleting Task",
			Message: err.Error(),
		})
	}

	return c.JSON(responses.SuccessResponse{
		Message:     "Success",
		Description: "Task Deleted Successfully",
	})
}
