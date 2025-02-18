package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"todoList/internal/common/models"
)

func GetMock(mockService *MockTaskService) []Mock {
	return []Mock{
		{
			Name: "Success",
			Mock: func() {
				mockService.On("Get", mock.Anything).Return([]models.Task{}, fiber.StatusOK, nil)
			},
			ExpectedStatus: fiber.StatusOK,
			ExpectedBody:   `[]`,
		},
	}
}

func GetPostMock(mockService *MockTaskService) []Mock {
	return []Mock{
		{
			Name: "Success",
			Body: `{"title":"test","description":"test"}`,
			Mock: func() {
				mockService.On("Create", mock.Anything, mock.Anything).Return(fiber.StatusCreated, nil)
			},
			ExpectedStatus: fiber.StatusCreated,
			ExpectedBody:   `{"message":"Success","description":"Task Created Successfully"}`,
		},
		{
			Name:           "Validation Error",
			Body:           `{"title":"","description":"test"}`,
			Mock:           func() {},
			ExpectedStatus: fiber.StatusBadRequest,
			ExpectedBody:   `{"error":"Error while validating Task","message":"title is required"}`,
		},
	}
}

func GetPutMock(mockService *MockTaskService) []Mock {
	return []Mock{
		{
			Name: "Success",
			ID:   "1",
			Body: `{"title":"test","description":"test", "status":"new"}`,
			Mock: func() {
				mockService.On("Update", mock.Anything, 1, mock.Anything).Return(fiber.StatusOK, nil)
			},
			ExpectedStatus: fiber.StatusOK,
			ExpectedBody:   `{"message":"Success","description":"Task Updated Successfully"}`,
		},
		{
			Name:           "Invalid ID",
			ID:             "invalid",
			Body:           `{"title":"test","description":"test"}`,
			Mock:           func() {},
			ExpectedStatus: fiber.StatusBadRequest,
			ExpectedBody:   `{"error":"Error while parsing id","message":"failed to convert: strconv.Atoi: parsing \"invalid\": invalid syntax"}`,
		},
		{
			Name:           "Validation Error",
			ID:             "1",
			Body:           `{"title":"test","description":"test"}`,
			Mock:           func() {},
			ExpectedStatus: fiber.StatusBadRequest,
			ExpectedBody:   `{"error":"Error while validating Task","message":"invalid status"}`,
		},
	}
}

func GetDeleteMock(mockService *MockTaskService) []Mock {
	return []Mock{
		{
			Name: "Success",
			ID:   "1",
			Mock: func() {
				mockService.On("Delete", mock.Anything, 1).Return(fiber.StatusOK, nil)
			},
			ExpectedStatus: fiber.StatusOK,
			ExpectedBody:   `{"message":"Success","description":"Task Deleted Successfully"}`,
		},
		{
			Name:           "Invalid ID",
			ID:             "invalid",
			Mock:           func() {},
			ExpectedStatus: fiber.StatusBadRequest,
			ExpectedBody:   `{"error":"Error while parsing id","message":"failed to convert: strconv.Atoi: parsing \"invalid\": invalid syntax"}`,
		},
	}
}
