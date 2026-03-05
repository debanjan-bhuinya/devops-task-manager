package main

import (
	"log"
	"net/http"
	"os"

	"github.com/debanjan-bhuinya/devops-task-manager/internal/database"
	"github.com/debanjan-bhuinya/devops-task-manager/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handler.HealthHandler)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", handler.CreateUser)
			r.Get("/", handler.GetUsers)
		})
	})

	log.Printf("Server running on :%s\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
