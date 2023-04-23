package domain

import "context"

type Task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type TaskRepository interface {
	Store(ctx context.Context, task Task) error
	Update(ctx context.Context, id string, task Task) error
	Fetch(ctx context.Context) ([]Task, error)
	Delete(ctx context.Context, id string) error
}

type TaskService interface {
	Store(ctx context.Context, task Task) error
	Update(ctx context.Context, id string, task Task) error
	Fetch(ctx context.Context) ([]Task, error)
	Delete(ctx context.Context, id string) error
}
