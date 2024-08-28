package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/hunderaweke/spher/api/controllers"
	"github.com/hunderaweke/spher/domain"
	"github.com/hunderaweke/spher/repositories/mongo"
	"github.com/hunderaweke/spher/usecases"
	"github.com/sv-tools/mongoifc"
)

func SetupTaskRoutes(r chi.Router, database mongoifc.Database) {
	taskCollection := database.Collection(domain.TaskCollection)
	addTaskRoutes(r, taskCollection)
}

func addTaskRoutes(r chi.Router, collection mongoifc.Collection) {
	repo := mongo.NewTaskRepository(context.TODO(), collection)
	u := usecases.NewTaskUsecase(repo)
	c := controllers.NewTaskController(u)
	taskRouter := chi.NewRouter()
	{
		taskRouter.Post("/", c.CreateTask)
		taskRouter.Get("/", c.FetchTasks)
		taskRouter.Put("/{taskID}", c.UpdateTask)
		taskRouter.Get("/{taskID}", c.FetchTaskByID)
		taskRouter.Delete("/{taskID}", c.DeleteTask)
	}
	r.Mount("/tasks", taskRouter)
}
