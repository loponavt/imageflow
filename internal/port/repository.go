package port

import "imageflow/internal/model"

type Repository interface {
	Save(task *model.ImageTask) error
	Get(id string) (*model.ImageTask, error)
	UpdateStatus(id, taskType string) error
}
