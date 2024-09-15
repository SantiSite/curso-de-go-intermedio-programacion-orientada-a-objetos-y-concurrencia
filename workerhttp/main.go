// Package main es el paquete principal, el cual, es el punto de entrada para el programa
package main

import (
	"github.com/SantiSite/workerhttp/constants"
	"github.com/SantiSite/workerhttp/dispatchers"
	serverHTTP "github.com/SantiSite/workerhttp/server/http"
	"github.com/SantiSite/workerhttp/types"
)

// main is the entry point of the programme
func main() {
	jobQueue := make(chan types.Job, constants.MaxQueueSise)
	d := dispatchers.NewDispatcher(jobQueue, constants.MaxWorkers)
	d.Run()

	serverHTTP.Server(jobQueue)
}
