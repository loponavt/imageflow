package usecase

import (
	"github.com/google/uuid"
	"imageflow/internal/model"
)

type Repository interface {
	Save(task *model.ImageTask) error
	Get(id string) (*model.ImageTask, error)
}

type ImageUseCase struct {
	repo Repository
}

func NewImageUseCase(r Repository) *ImageUseCase {
	return &ImageUseCase{repo: r}
}

func (u *ImageUseCase) Submit(filename string) (string, error) {
	id := uuid.New().String()
	task := &model.ImageTask{
		ID:       id,
		Filename: filename,
		Status:   "pending",
	}
	return id, u.repo.Save(task)
}

func (u *ImageUseCase) GetStatus(id string) (*model.ImageTask, error) {
	return u.repo.Get(id)
}
