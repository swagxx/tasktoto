package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testask/internal/domain"
	"testask/internal/storage"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskService struct {
	storage *storage.MemoryStorage
}

func NewTaskService(storage *storage.MemoryStorage) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) CreateTask(ctx context.Context) (*domain.Task, error) {
	task := &domain.Task{
		ID:        generateID(),
		Status:    domain.StatusPending,
		CreatedAt: time.Now(),
	}

	if err := s.storage.Create(task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	go s.processTask(task.ID)

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*domain.TaskResponse, error) {
	task, err := s.storage.Get(id)
	if err != nil {
		return nil, err
	}

	response := &domain.TaskResponse{
		ID:        task.ID,
		Status:    task.Status,
		Result:    task.Result,
		Error:     task.Error,
		CreatedAt: task.CreatedAt,
	}

	if task.CompletedAt != nil && task.StartedAt != nil {
		response.Duration = task.CompletedAt.Sub(*task.StartedAt).String()
	} else if task.StartedAt != nil {
		response.Duration = time.Since(*task.StartedAt).String()
	}

	return response, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.storage.Delete(id)
}

func (s *TaskService) processTask(id string) {
	task, err := s.storage.Get(id)
	if err != nil {
		return
	}

	task.Status = domain.StatusRunning
	now := time.Now()
	task.StartedAt = &now
	if err := s.storage.Update(task); err != nil {
		log.Printf("failed to update completed task %s: %v", id, err)
	}

	time.Sleep(time.Duration(180+time.Now().UnixNano()%120) * time.Second)

	task.Status = domain.StatusCompleted
	task.Result = "Task completed successfully"
	now = time.Now()
	task.CompletedAt = &now
	if err := s.storage.Update(task); err != nil {
		log.Printf("failed to update completed task %s: %v", id, err)
	}

}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
