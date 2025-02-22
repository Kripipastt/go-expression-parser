package api

import (
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/api/handlers/v1"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/calculate", handlers.CalcHandler)
	router.Get("/expressions", handlers.GetExpressionsHandler)
	router.Get("/expressions/{id}", handlers.GetOneExpressionHandler)
	return router
}
