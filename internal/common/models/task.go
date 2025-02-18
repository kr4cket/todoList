package models

import (
	"errors"
	"time"
)

type Status string

const (
	InProgress Status = "in_progress"
	Done       Status = "done"
	New        Status = "new"
)

type Task struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func ValidateUpdateTask(task Task) error {
	if !isStatusValid(task.Status) {
		return errors.New("invalid status")
	}

	return nil
}

func ValidateCreateTask(task Task) error {
	if isTitleEmpty(task.Title) {
		return errors.New("title is required")
	}
	return nil
}

func isTitleEmpty(name string) bool {
	return name == ""
}

func isStatusValid(status Status) bool {
	switch status {
	case InProgress:
		return true
	case Done:
		return true
	case New:
		return true
	default:
		return false
	}
}
