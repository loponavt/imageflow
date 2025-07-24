package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"imageflow/internal/model"
)

type Repo struct {
	db *gorm.DB
}

func NewPostgresRepo(host, port, user, password, dbname string) (*Repo, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repo{db: db}, nil
}

func (r *Repo) Save(task *model.ImageTask) error {
	return r.db.Create(task).Error
}

func (r *Repo) Get(id string) (*model.ImageTask, error) {
	var task model.ImageTask
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repo) UpdateStatus(id, status string) error {
	return r.db.Model(&model.ImageTask{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}
