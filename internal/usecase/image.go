package usecase

import (
	"errors"
	"github.com/google/uuid"
	"imageflow/internal/model"
	"imageflow/internal/port"
	"imageflow/internal/worker"
)

var ErrUnknownTaskType = errors.New("unknown task type")

type ImageUseCase struct {
	repo port.Repository
	pool *worker.Pool
}

func NewImageUseCase(r port.Repository, pool *worker.Pool) *ImageUseCase {
	return &ImageUseCase{
		repo: r,
		pool: pool,
	}
}

func (uc *ImageUseCase) Submit(filename, taskType string) (string, error) {
	task := &model.ImageTask{
		ID:       uuid.NewString(),
		Filename: filename,
		Type:     taskType,
		Status:   "pending",
	}
	if err := uc.repo.Save(task); err != nil {
		return "", err
	}

	var job worker.Job
	switch taskType {
	case "resize":
		job = worker.NewResizeJob(task.ID, filename, uc.repo)
	case "blur":
		job = worker.NewBlurJob(task.ID, filename, uc.repo)
	default:
		return "", ErrUnknownTaskType
	}

	uc.pool.Submit(job)

	return task.ID, nil
}

func (u *ImageUseCase) GetStatus(id string) (*model.ImageTask, error) {
	return u.repo.Get(id)
}
