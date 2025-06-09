package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"testask/internal/domain"
	"testask/internal/service"
)

type Handler struct {
	taskService *service.TaskService
}

func NewHandler(taskService *service.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.taskService.CreateTask(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := domain.TaskResponse{
		ID:        task.ID,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	response, err := h.taskService.GetTask(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.taskService.DeleteTask(r.Context(), id); err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
