package repositories

import (
	"time"

	"github.com/hunderaweke/spher/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	db.AutoMigrate(&domain.Task{})
	return &taskRepository{db: db}
}

func (t *taskRepository) Create(task domain.Task) (*domain.Task, error) {
	result := t.db.Create(&task)
	if result.Error != nil {
		return &task, result.Error
	}
	return &task, nil
}

func (t *taskRepository) Fetch() ([]domain.Task, error) {
	tasks := []domain.Task{}
	result := t.db.Find(&tasks)
	return tasks, result.Error
}

func (t *taskRepository) FetchByID(id uint) (*domain.Task, error) {
	task := domain.Task{}
	result := t.db.Find(&task, "id = ?", id)
	return &task, result.Error
}

func (t *taskRepository) FetchByTag(tag string) ([]domain.Task, error) {
	task := []domain.Task{}
	result := t.db.Find(&task, "tag= ?", tag)
	return task, result.Error
}

func (t *taskRepository) FetchByDeadline(deadline time.Time) ([]domain.Task, error) {
	task := []domain.Task{}
	result := t.db.Find(&task, "deadline = ?", deadline)
	return task, result.Error
}

func (t *taskRepository) FetchByPriority(priority int) ([]domain.Task, error) {
	task := []domain.Task{}
	result := t.db.Find(&task, "priority = ?", priority)
	return task, result.Error
}

func (t *taskRepository) FetchByStatus(status string) ([]domain.Task, error) {
	task := []domain.Task{}
	result := t.db.Find(&task, "status = ?", status)
	return task, result.Error
}
