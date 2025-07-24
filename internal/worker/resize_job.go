package worker

import (
	"imageflow/internal/port"
	"log/slog"
	"time"
)

type ResizeJob struct {
	ID       string
	Filename string
	Repo     port.Repository
}

func NewResizeJob(id, filename string, repo port.Repository) *ResizeJob {
	return &ResizeJob{ID: id, Filename: filename, Repo: repo}
}

func (j *ResizeJob) Process() error {
	slog.Info("ResizeJob started", "id", j.ID, "filename", j.Filename)

	// TODO: логика
	time.Sleep(1 * time.Second)

	// Обновляем статус в репозитории
	if err := j.Repo.UpdateStatus(j.ID, "done"); err != nil {
		slog.Error("failed to update task status", "id", j.ID, "error", err)
		return err
	}
	slog.Info("ResizeJob completed", "id", j.ID)
	return nil
}
