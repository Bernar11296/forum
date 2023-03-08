package main

import (
	"fmt"
	"forum/internal/handler"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"log"
)

func main() {
	configDB := repository.NewConfigDB()
	db, err := repository.InitDB(configDB)
	if err != nil {
		log.Fatalf("failed to initialize db : %s", err.Error())
	}
	if err := repository.CreateTables(db); err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	services := service.NewService(*repo)
	handler := handler.NewHandler(services)

	server := new(server.Server)
	fmt.Printf("Starting server at port %s\nhttp://localhost%s/\n", ":8080", ":8080")
	if err := server.Run(":8080", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
