package main

import (
	http_server "github.com/aftab-hussain-93/empl/http"
	"github.com/aftab-hussain-93/empl/repository"
	routes "github.com/aftab-hussain-93/empl/routes"
	"github.com/aftab-hussain-93/empl/service"
)

func main() {
	db, closer := NewDB()
	defer closer()
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	http_server.RunHTTPServer(routes.CreateHandler(service))
}
