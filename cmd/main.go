package main

import (
	"log"
	"os"

	"github.com/gavrylenkoIvan/gonotes/pkg/handlers"
	"github.com/gavrylenkoIvan/gonotes/pkg/repository"
	"github.com/gavrylenkoIvan/gonotes/pkg/service"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err.Error())
	}

	db, err := repository.InitDB(os.Getenv("psql_url"))
	if err != nil {
		logger.Fatal(err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handlers.NewHandler(service)
	server := handler.InitRoutes()

	server.Run(":" + os.Getenv("PORT"))
}
