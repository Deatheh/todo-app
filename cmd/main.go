package main

import (
	"log"

	"github.com/Deatheh/todo-app"
	"github.com/Deatheh/todo-app/pkg/handler"
	"github.com/Deatheh/todo-app/pkg/repository"
	"github.com/Deatheh/todo-app/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("errore initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("&s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
