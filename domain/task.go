package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PriorityHigh   = 1
	PriorityMedium = 2
	PriorityLow    = 3
	TaskCollection = "tasks"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `json:"title"`
	Tags        []string           `json:"tags"`
	Description string             `json:"description,omitempty"`
	Status      string             `json:"status"`
	StartTime   time.Time          `bson:"start_time" json:"start_time"`
	Deadline    time.Time          `json:"deadline"`
	Priority    int                `json:"priority"`
}
type (
	TaskRepository interface {
		Create(t Task) (*Task, error)
		Fetch(filterOptions map[string]interface{}) ([]Task, error)
		FetchByID(id string) (*Task, error)
	}
	TaskUsecase interface {
		Create(t Task) (*Task, error)
		Fetch(filterOptions map[string]interface{}) ([]Task, error)
		FetchByID(id string) (*Task, error)
	}
)
