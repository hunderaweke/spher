package domain

import (
	"context"
	"time"
)

const CollectionTask = "tasks"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	UserID      string    `json:"user_id"`
	DueDate     time.Time `json:"due_date"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FindByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FindByUserID(c context.Context, userID string) ([]Task, error)
}
