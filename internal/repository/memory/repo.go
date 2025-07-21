package memory

import (
	"sync"

	"imageflow/internal/model"
)

type InMemoryRepo struct {
	mu    sync.RWMutex
	tasks map[string]*model.ImageTask
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{tasks: make(map[string]*model.ImageTask)}
}

func (r *InMemoryRepo) Save(task *model.ImageTask) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepo) Get(id string) (*model.ImageTask, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.tasks[id], nil
}
