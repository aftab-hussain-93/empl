package main

import (
	"github.com/aftab-hussain-93/empl/handlers"
	"github.com/aftab-hussain-93/empl/repository"
	"github.com/aftab-hussain-93/empl/service"
)

func main() {
	db, closer := NewDB()
	defer closer()
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	RunHTTPServer(handlers.CreateHandler(service))
}
