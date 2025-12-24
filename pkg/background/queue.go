package background

import (
	"context"
	"log/slog"

	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type queue struct {
	tasks       chan Task
	rateLimiter *rate.Limiter
}

type Worker struct {
	queues map[TaskType]queue
	logger *slog.Logger
	db     *gorm.DB
}

func NewWorker(logger *slog.Logger, db *gorm.DB) *Worker {
	return &Worker{
		queues: map[TaskType]queue{
			//TODO
			"taskTypeUpdateMapDataAddress": {
				tasks:       make(chan Task, 100),
				rateLimiter: rate.NewLimiter(1, 10),
			},
		},
		logger: logger,
		db:     db,
	}
}

func (w *Worker) Run() {
	ctx := context.Background()
	for k, v := range w.queues {
		go func() {
			w.logger.Info("starting worker", "queue", k)
			for {
				if v.rateLimiter != nil {
					if err := v.rateLimiter.Wait(ctx); err != nil {
						w.logger.Error("rate limiter error", "err", err)
						return
					}
				}
				task := <-v.tasks
				if err := task.Run(w.db); err != nil {
					w.logger.Error("task run error", "err", err)
					continue
				}
			}
		}()
	}
}

// will block is wait queue is full
func (w *Worker) Submit(t Task) {
	w.queues[t.TaskType()].tasks <- t
}
