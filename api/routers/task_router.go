package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/hunderaweke/spher/api/controllers"
	"github.com/hunderaweke/spher/domain"
	"github.com/hunderaweke/spher/repositories/mongo"
	"github.com/hunderaweke/spher/usecase"
	"github.com/sv-tools/mongoifc"
)

func SetupTaskRoutes(r chi.Router, database mongoifc.Database) {
	taskCollection := database.Collection(domain.TaskCollection)
	addTaskRoutes(r, taskCollection)
}

func addTaskRoutes(r chi.Router, collection mongoifc.Collection) {
	repo := mongo.NewTaskRepository(context.TODO(), collection)
	u := usecase.NewTaskUsecase(repo)
	c := controllers.NewTaskController(u)
	taskRouter := chi.NewRouter()
	{
		taskRouter.Post("/", c.CreateTask)
		taskRouter.Get("/", c.FetchTasks)
		taskRouter.Get("/{taskID}", c.FetchTaskByID)
	}
	r.Mount("/tasks", taskRouter)
}
