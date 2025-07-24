package worker

import (
	"imageflow/internal/port"
	"log/slog"
	"time"
)

type BlurJob struct {
	ID       string
	Filename string
	Repo     port.Repository
}

func NewBlurJob(id, filename string, repo port.Repository) *BlurJob {
	return &BlurJob{ID: id, Filename: filename, Repo: repo}
}

func (j *BlurJob) Process() error {
	slog.Info("BlurJob started", "id", j.ID)

	// TODO: логика
	time.Sleep(1 * time.Second)

	if err := j.Repo.UpdateStatus(j.ID, "done"); err != nil {
		slog.Error("BlurJob update failed", "id", j.ID, "err", err)
		return err
	}
	slog.Info("BlurJob finished", "id", j.ID)
	return nil
}
