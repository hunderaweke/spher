package mongo

import (
	"context"
	"time"

	"github.com/hunderaweke/spher/domain"
	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	collection mongoifc.Collection
	ctx        context.Context
}

func NewTaskRepository(ctx context.Context, collection mongoifc.Collection) domain.TaskRepository {
	return &taskRepository{collection: collection, ctx: ctx}
}

func (repo *taskRepository) Create(t domain.Task) (*domain.Task, error) {
	t.ID = primitive.NewObjectID()
	if t.StartTime.IsZero() {
		t.StartTime = time.Now().Add(24 * time.Hour)
	}
	if t.Deadline.IsZero() {
		t.Deadline = time.Now().Add(24 * time.Hour)
	}
	_, err := repo.collection.InsertOne(repo.ctx, t)
	if err != nil {
		return &domain.Task{}, err
	}
	return &t, nil
}

func (repo *taskRepository) Fetch(filterOptions map[string]interface{}) ([]domain.Task, error) {
	filter := bson.M{}
	for key, val := range filterOptions {
		if key == "start_time" || key == "deadline" {
			date := val.(time.Time)
			date = truncateTime(date)
			filter[key] = bson.M{
				"$gte": primitive.NewDateTimeFromTime(date),
				"$lt":  primitive.NewDateTimeFromTime(date.Add(24 * time.Hour)),
			}
		} else if key == "tags" {
			filter["tags"] = bson.M{"$all": val}
		} else {
			filter[key] = val
		}
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

func (repo *taskRepository) FetchByID(id string) (*domain.Task, error) {
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.Task{}, err
	}
	var task domain.Task
	err = repo.collection.FindOne(repo.ctx, bson.M{"_id": taskID}).Decode(&task)
	if err != nil {
		return &domain.Task{}, err
	}
	if err != nil {
		return &domain.Task{}, err
	}
	return &task, nil
}

func truncateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
