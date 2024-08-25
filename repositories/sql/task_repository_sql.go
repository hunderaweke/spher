package repositories

import (
	"time"

	"github.com/hunderaweke/spher/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

/* func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	db.AutoMigrate(&domain.Task{})
	return &taskRepository{db: db}
} */

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

func (t *taskRepository) FetchByID(id string) (*domain.Task, error) {
	task := domain.Task{}
	result := t.db.Find(&task, "id = ?", id)
	return &task, result.Error
}

func (t *taskRepository) FetchByTags(tags []string) ([]domain.Task, error) {
	tasks := []domain.Task{}
	err := t.db.Preload("tags").Joins("JOIN task_tags ON task_task.task_id = tasks.id").
		Joins("JOIN tags ON tags.id = task_tags.tag_id").
		Where("tags.name IN ?", tags).
		Group("tasks.id").
		Having("COUNT(tags.id)=?", len(tags)).
		Find(&tasks)
	if err != nil {
		return []domain.Task{}, err.Error
	}
	return tasks, nil
}

func (t *taskRepository) FetchByDeadline(deadline time.Time) ([]domain.Task, error) {
	tasks := []domain.Task{}
	result := t.db.Where(&tasks, "DATE(deadline) = ?", deadline.Format("2006-01-02")).Find(&tasks)
	return tasks, result.Error
}

func (t *taskRepository) FetchByStartTime(deadline time.Time) ([]domain.Task, error) {
	tasks := []domain.Task{}
	result := t.db.Where(&tasks, "DATE(deadline) = ?", deadline.Format("2006-01-02")).Find(&tasks)
	return tasks, result.Error
}

func (t *taskRepository) FetchByPriority(priority int) ([]domain.Task, error) {
	tasks := []domain.Task{}
	result := t.db.Find(&tasks, "priority = ?", priority)
	return tasks, result.Error
}

func (t *taskRepository) FetchByStatus(status string) ([]domain.Task, error) {
	task := []domain.Task{}
	result := t.db.Find(&task, "status = ?", status)
	return task, result.Error
}
