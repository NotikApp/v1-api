package main

import (
	"log"
	"os"

	"github.com/gavrylenkoIvan/gonotes/pkg/handlers"
	"github.com/gavrylenkoIvan/gonotes/pkg/repository"
	"github.com/gavrylenkoIvan/gonotes/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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

	if err := initConfig(); err != nil {
		logger.Fatal(err.Error())
	}

	dbPass := os.Getenv("DB_PASSWORD")
	config := repository.Config{
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.name"),
		Host:     viper.GetString("db.host"),
		User:     viper.GetString("db.user"),
		Password: dbPass,
		SSL:      viper.GetString("db.ssl"),
	}
	db, err := repository.InitDB(config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handlers.NewHandler(service)
	server := handler.InitRoutes()
	server.Run(":8080")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
