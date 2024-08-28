package usecases

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

func (usecase *taskUsecase) Fetch(filterOptions map[string]interface{}, page, limit int) ([]domain.Task, error) {
	return usecase.taskRepository.Fetch(filterOptions, page, limit)
}

func (usecase *taskUsecase) FetchByID(id string) (*domain.Task, error) {
	return usecase.taskRepository.FetchByID(id)
}

func (usecase *taskUsecase) Update(taskID string, data domain.Task) (domain.Task, error) {
	return usecase.taskRepository.Update(taskID, data)
}

func (usecase *taskUsecase) Delete(taskID string) error {
	return usecase.taskRepository.Delete(taskID)
}
