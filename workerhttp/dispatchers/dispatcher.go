package dispatchers

import (
	"github.com/SantiSite/workerhttp/types"
	"github.com/SantiSite/workerhttp/workers"
)

type Dispatcher struct {
	WorkerPool types.WorkerPool
	MaxWorkers int
	JobQueue   chan types.Job
}

func NewDispatcher(jobQueue chan types.Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(types.WorkerPool, maxWorkers),
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := workers.NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}
