package main

import (
	"log"

	"github.com/aftab-hussain-93/empl/handlers"
	"github.com/aftab-hussain-93/empl/repository"
	"github.com/aftab-hussain-93/empl/service"
)

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	RunHTTPServer(handlers.CreateHandler(service))
}
