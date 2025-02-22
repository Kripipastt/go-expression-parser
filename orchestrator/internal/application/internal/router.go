package internal

import (
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/internal/handlers/v1"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/task", handlers.GetTaskHandler)
	router.Post("/task", handlers.PostTaskResultHandler)
	return router
}
