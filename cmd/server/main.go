package main

import (
	"log"
	"net/http"
	"testask/internal/api"
	"testask/internal/service"
	"testask/internal/storage"
)

func main() {

	taskStorage := storage.NewMemoryStorage()

	taskService := service.NewTaskService(taskStorage)

	apiHandler := api.NewHandler(taskService)
	router := api.NewRouter(apiHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
