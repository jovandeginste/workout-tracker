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

// queues will only be written in init function
var queues map[TaskType]queue = make(map[TaskType]queue)

type Worker struct {
	logger *slog.Logger
	db     *gorm.DB
}

type RegisterOpts struct {
	BufferSize  *int // if nil, a default size 100 will be used
	RateLimiter *rate.Limiter
}

func RegisterQueue(taskType TaskType, opts RegisterOpts) {
	q := queue{
		rateLimiter: opts.RateLimiter,
	}
	if opts.BufferSize != nil {
		if *opts.BufferSize == 0 {
			q.tasks = make(chan Task)
		} else {
			q.tasks = make(chan Task, *opts.BufferSize)
		}
	} else {
		q.tasks = make(chan Task, 100)
	}
	queues[taskType] = q
}

func NewWorker(logger *slog.Logger, db *gorm.DB) *Worker {
	return &Worker{
		logger: logger,
		db:     db,
	}
}

func (w *Worker) Run() {
	ctx := context.Background()
	for k, v := range queues {
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
	queues[t.TaskType()].tasks <- t
}
