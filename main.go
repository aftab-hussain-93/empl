package main

import (
	"github.com/aftab-hussain-93/empl/internal/repository"
	routes "github.com/aftab-hussain-93/empl/internal/routes"
	"github.com/aftab-hussain-93/empl/internal/service"
	"github.com/aftab-hussain-93/empl/pkg/postgres"
	http_server "github.com/aftab-hussain-93/empl/pkg/server"
)

func main() {
	db := postgres.New()
	defer db.Close()
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	http_server.RunHTTPServer(routes.CreateHandler(service))
}
