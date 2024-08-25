package usecase

import (
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

func (usecase *taskUsecase) Fetch(filterOptions map[string]interface{}) ([]domain.Task, error) {
	return usecase.taskRepository.Fetch(filterOptions)
}

func (usecase *taskUsecase) FetchByID(id string) (*domain.Task, error) {
	return usecase.taskRepository.FetchByID(id)
}
