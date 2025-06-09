package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Post("/tasks", h.CreateTask)
	r.Get("/tasks/{id}", ValidateTaskID(h.GetTask))
	r.Delete("/tasks/{id}", ValidateTaskID(h.DeleteTask))

	return r
}
