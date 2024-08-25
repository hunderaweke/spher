package mongo

import (
	"context"
	"time"

	"github.com/hunderaweke/spher/domain"
	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
)

type taskRepository struct {
	collection mongoifc.Collection
	ctx        context.Context
}

func (repo *taskRepository) Create(t domain.Task) (*domain.Task, error) {
	_, err := repo.collection.InsertOne(repo.ctx, t)
	if err != nil {
		return &domain.Task{}, err
	}
	return &t, nil
}

func (repo *taskRepository) Fetch() ([]domain.Task, error) {
	resp, err := repo.collection.Find(repo.ctx, bson.M{})
	if err != nil {
		return []domain.Task{}, err
	}
	tasks := []domain.Task{}
	for resp.Next(repo.ctx) {
		var t domain.Task
		if err := resp.Decode(&t); err != nil {
			return []domain.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (repo *taskRepository) FetchByID(id uint) (*domain.Task, error) {
	resp, err := repo.collection.Find(repo.ctx, bson.M{"id": id})
	if err != nil {
		return &domain.Task{}, err
	}
	var task domain.Task
	err = resp.Decode(&task)
	if err != nil {
		return &domain.Task{}, err
	}
	return &task, nil
}

func (repo *taskRepository) FetchByTags(tags []string) ([]domain.Task, error) {
	filter := bson.M{
		"tags": bson.M{"$all": tags},
	}
	resp, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return []domain.Task{}, err
	}
	tasks := []domain.Task{}
	for resp.Next(repo.ctx) {
		var t domain.Task
		if err := resp.Decode(&t); err != nil {
			return []domain.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func truncateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (repo *taskRepository) FetchByDeadline(deadline time.Time) ([]domain.Task, error) {
	deadline = truncateTime(deadline)
	filter := bson.M{
		"deadline": bson.M{
			"$gte": deadline,
			"$lt":  deadline.Add(24 * time.Hour),
		},
	}
	resp, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return []domain.Task{}, err
	}
	tasks := []domain.Task{}
	for resp.Next(repo.ctx) {
		var t domain.Task
		if err := resp.Decode(&t); err != nil {
			return []domain.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (repo *taskRepository) FetchByStartTime(startTime time.Time) ([]domain.Task, error) {
	startTime = truncateTime(startTime)
	filter := bson.M{
		"start_time": bson.M{
			"$gte": startTime,
			"$lt":  startTime.Add(24 * time.Hour),
		},
	}
	resp, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return []domain.Task{}, err
	}
	tasks := []domain.Task{}
	for resp.Next(repo.ctx) {
		var t domain.Task
		if err := resp.Decode(&t); err != nil {
			return []domain.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (repo *taskRepository) FetchByPriority(priority int) ([]domain.Task, error) {
	filter := bson.M{"priority": priority}
	resp, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return []domain.Task{}, err
	}
	tasks := []domain.Task{}
	for resp.Next(repo.ctx) {
		var t domain.Task
		if err := resp.Decode(&t); err != nil {
			return []domain.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (repo *taskRepository) FetchByStatus(status string) ([]domain.Task, error) {
	filter := bson.M{"status": status}
	resp, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return []domain.Task{}, err
	}
	tasks := []domain.Task{}
	for resp.Next(repo.ctx) {
		var t domain.Task
		if err := resp.Decode(&t); err != nil {
			return []domain.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
