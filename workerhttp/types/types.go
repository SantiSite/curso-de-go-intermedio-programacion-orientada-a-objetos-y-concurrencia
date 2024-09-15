package types

import (
	"time"
)

// A Job is
type Job struct {
	// A Name es el nombre que se le da al trabajo
	Name string
	// A Delay es el tiempo extra de demora del Job (simulación)
	Delay time.Duration
	// A Number número de la seria de Fibonacci que se va a calcular
	Number int
}

// WorkerPool .
type WorkerPool chan chan Job
