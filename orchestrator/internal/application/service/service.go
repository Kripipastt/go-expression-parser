package service

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// PingHandler
// @Summary Ping
// @Description Ping for healthcheck
// @Tags Other
// @Produce json
// @Success 	200 {string} string
// @Router		/ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/ping", PingHandler)
	return router
}
