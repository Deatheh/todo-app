package main

import (
	"log"

	"github.com/Deatheh/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("&s", err.Error())
	}

}
