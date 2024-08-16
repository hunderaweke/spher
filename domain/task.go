package domain

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
type Task struct {
	gorm.Model
	Title       string    `json:"title"`
	Tags        []Tag     `gorm:"many2many:task_tags;" json:"tags"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"done"`
	StartTime   time.Time `json:"start_time"`
	Deadline    time.Time `json:"deadline"`
	Priority    int       `json:"priority"`
}
type (
	TaskRepository interface {
		Create(t Task) (*Task, error)
		Fetch() ([]Task, error)
		FetchByID(id uint) (*Task, error)
		FetchByTag(tag string) ([]Task, error)
		FetchByDeadline(deadline time.Time) ([]Task, error)
		FetchByPriority(priority int) ([]Task, error)
		FetchByStatus(status string) ([]Task, error)
	}
)
