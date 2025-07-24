package worker

import (
	"context"
	"log/slog"
)

type Pool struct {
	workerCount int
	jobQueue    chan Job
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewPool(workers int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		workerCount: workers,
		jobQueue:    make(chan Job),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.workerCount; i++ {
		go func(id int) {
			slog.Info("Worker started", "id", id)
			for {
				select {
				case job := <-p.jobQueue:
					if err := job.Process(); err != nil {
						slog.Error("Job failed", "err", err)
					}
				case <-p.ctx.Done():
					slog.Info("Worker stopped", "id", id)
					return
				}
			}
		}(i)
	}
}

func (p *Pool) Submit(job Job) {
	p.jobQueue <- job
}

func (p *Pool) Stop() {
	p.cancel()
	close(p.jobQueue)
}
