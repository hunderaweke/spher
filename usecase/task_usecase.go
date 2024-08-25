package usecase

import (
	"time"

	"github.com/hunderaweke/spher/domain"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

func (usecase *taskUsecase) Create(t domain.Task) (*domain.Task, error) {
	return usecase.taskRepository.Create(t)
}

func (usecase *taskUsecase) Fetch() ([]domain.Task, error) {
	return usecase.taskRepository.Fetch()
}

func (usecase *taskUsecase) FetchByID(id uint) (*domain.Task, error) {
	return usecase.taskRepository.FetchByID(id)
}

func (usecase *taskUsecase) FetchByTags(tags []string) ([]domain.Task, error) {
	return usecase.taskRepository.FetchByTags(tags)
}

func (usecase *taskUsecase) FetchByDeadline(deadline time.Time) ([]domain.Task, error) {
	return usecase.taskRepository.FetchByDeadline(deadline)
}

func (usecase *taskUsecase) FetchByStatus(status string) ([]domain.Task, error) {
	return usecase.taskRepository.FetchByStatus(status)
}

func (usecase *taskUsecase) FetchByPriority(priority int) ([]domain.Task, error) {
	return usecase.taskRepository.FetchByPriority(priority)
}

func (usecase *taskUsecase) FetchByStartTime(startTime time.Time) ([]domain.Task, error) {
	return usecase.taskRepository.FetchByStartTime(startTime)
}
