package routers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hunderaweke/spher/config"
	"github.com/hunderaweke/spher/domain"
	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	router, err := setupRouters(config)
	log.Fatal(http.ListenAndServe(config.Server.Port, router))
}

func setupRouters(config config.Config) (chi.Router, error) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Mount("/debug", middleware.Profiler())
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Database.Url))
	client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	database := client.Database(config.Database.Name)
	taskCollection := database.Collection(domain.TaskCollection)
	addTaskRoutes(router, mongoifc.WrapCollection(taskCollection))
	return router, nil
}
