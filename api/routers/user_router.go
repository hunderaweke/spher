package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/hunderaweke/spher/api/controllers"
	"github.com/hunderaweke/spher/repositories/mongo"
	"github.com/hunderaweke/spher/usecases"
	"github.com/sv-tools/mongoifc"
)

func addUserRouters(router chi.Router, db mongoifc.Database) {
	repo := mongo.NewMongoUserRepository(db, context.Background())
	u := usecases.NewUserUsecases(repo)
	c := controllers.NewUserControllers(u)
	userRouter := chi.NewRouter()
	{
		userRouter.Post("/register", c.Register)
	}
	router.Mount("/users", userRouter)
}
