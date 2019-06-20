package executor

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

type Payload interface {
	Execute(context.Context) error
}

type Job struct {
	Payload
}

type Executor struct {
	Timeout time.Duration
	Jobs    []*Job
}

func New(maxWorkers int, timeout time.Duration) *Executor {
	return &Executor{
		Timeout: timeout,
		Jobs:    make([]*Job, 0),
	}
}

func (ex *Executor) AddPayload(payload Payload) {
	ex.Jobs = append(ex.Jobs, &Job{payload})
}

func (ex *Executor) Start() error {
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, job := range ex.Jobs {
		job := job
		eg.Go(func() error {
			return job.Execute(ctx)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
