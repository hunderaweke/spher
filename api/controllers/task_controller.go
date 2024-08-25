package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hunderaweke/spher/api/utils"
	"github.com/hunderaweke/spher/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(u domain.TaskUsecase) TaskController {
	return TaskController{taskUsecase: u}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	newTask, err := utils.Decode[domain.Task](r.Body)
	if err != nil {
		utils.PostJSON(w, map[string]string{"message": err.Error()}, http.StatusNotAcceptable)
		return
	}
	createdTask, err := c.taskUsecase.Create(newTask)
	if err != nil {
		utils.PostJSON(w, map[string]string{"message": err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.PostJSON(w, createdTask, http.StatusOK)
}

func (c *TaskController) FetchTasks(w http.ResponseWriter, r *http.Request) {
	filterOptions := make(map[string]interface{})
	if tagsQuery := r.URL.Query().Get("tags"); tagsQuery != "" {
		tags := strings.Split(tagsQuery, ",")
		filterOptions["tags"] = tags
	}

	if date := r.URL.Query().Get("deadline"); date != "" {
		deadline, err := time.Parse("2006-01-02", date)
		if err != nil {
			utils.PostJSON(w, map[string]string{"error": "invalid date format use 'YYYY-MM-DD'"}, http.StatusNotAcceptable)
			return
		}
		filterOptions["deadline"] = deadline
	}
	if date := r.URL.Query().Get("startTime"); date != "" {
		startTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			utils.PostJSON(w, map[string]string{"error": "invalid date format use 'YYYY-MM-DD'"}, http.StatusNotAcceptable)
			return
		}
		filterOptions["start_time"] = startTime
	}
	if status := r.URL.Query().Get("status"); status != "" {
		filterOptions["status"] = status
	}
	tasks, err := c.taskUsecase.Fetch(filterOptions)
	if err != nil {
		utils.PostJSON(w, map[string]string{"message": err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.PostJSON(w, tasks, http.StatusOK)
}

func (c *TaskController) FetchTaskByID(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	if !primitive.IsValidObjectID(taskID) {
		utils.PostJSON(w, map[string]string{"message": "invalid object ID"}, http.StatusNotAcceptable)
		return
	}
	task, err := c.taskUsecase.FetchByID(taskID)
	if err != nil {
		utils.PostJSON(w, map[string]string{"message": err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.PostJSON(w, task, http.StatusOK)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	if !primitive.IsValidObjectID(taskID) {
		utils.PostJSON(w, map[string]string{"message": "invalid object ID"}, http.StatusNotAcceptable)
		return
	}
	updateData, err := utils.Decode[domain.Task](r.Body)
	if err != nil {
		utils.PostJSON(w, map[string]string{"error": "unacceptable data"}, http.StatusNotAcceptable)
		return
	}
	task, err := c.taskUsecase.Update(taskID, updateData)
	if err != nil {
		utils.PostJSON(w, map[string]string{"message": err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.PostJSON(w, task, http.StatusOK)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	if !primitive.IsValidObjectID(taskID) {
		utils.PostJSON(w, map[string]string{"message": "invalid object ID"}, http.StatusNotAcceptable)
		return
	}
	err := c.taskUsecase.Delete(taskID)
	if err != nil {
		utils.PostJSON(w, map[string]string{"error": err.Error()}, http.StatusNotAcceptable)
		return
	}
	utils.PostJSON(w, "", http.StatusNoContent)
}
