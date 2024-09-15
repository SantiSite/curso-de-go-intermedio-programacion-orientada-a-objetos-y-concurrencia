package workers

import (
	"fmt"
	"time"

	"github.com/SantiSite/workerhttp/functions"
	"github.com/SantiSite/workerhttp/types"
)

// Worker is
type Worker struct {
	// A Id es un identificador único para cada trabajador
	Id         int
	JobQueue   chan types.Job
	WorkerPool types.WorkerPool
	QuiteChan  chan bool
}

func NewWorker(id int, workerPool types.WorkerPool) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan types.Job),
		WorkerPool: workerPool,
		QuiteChan:  make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with Id %d started. Job: %+v\n", w.Id, job)
				fib := functions.Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Job with the number %d. Worker with Id %d finished with result %d.\n", job.Number, w.Id, fib)
				//w.Stop() No me queda claro donde ejecuto el w.Stop()
			case <-w.QuiteChan:
				fmt.Printf("Worker with Id %d stopped\n", w.Id)
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuiteChan <- true
	}()
}
