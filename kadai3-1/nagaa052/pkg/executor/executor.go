package executor

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

type Payload interface {
	Execute() error
}

type Job struct {
	Payload
}

type Executor struct {
	Timeout time.Duration
	Jobs    []Job
}

func New(maxWorkers int, timeout time.Duration) *Executor {
	return &Executor{
		Timeout: timeout,
		Jobs:    make([]Job, 0),
	}
}

func (ex *Executor) AddJob(job Job) {
	ex.Jobs = append(ex.Jobs, job)
}

func (ex *Executor) Start() error {
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, job := range ex.Jobs {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				return job.Execute()
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
