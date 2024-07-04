package asyncjob

import (
	"context"
	"food-delivery/common"
	"sync"
)

type group struct {
	isConcurrent bool // if true jobs run concurrent, otherwise sequence
	jobs         []Job
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           &sync.WaitGroup{},
	}
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				defer common.AppRecover()
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])

			continue
		}

		job := g.jobs[i]
		errChan <- g.runJob(ctx, job)
		g.wg.Done()
	}
	g.wg.Wait()

	var err error
	for i := 0; i < len(g.jobs); i++ {

		if v := <-errChan; v != nil {
			err = v
		}
	}


	return err
}

func (g *group) runJob(ctx context.Context, job Job) error {
	if err := job.Execute(ctx); err != nil {
		for {

			if job.State() == StateRetryFailed {
				return err
			}
			if err := job.Retry(ctx); err != nil {
				return nil
			}
		}
	}
	return nil
}
