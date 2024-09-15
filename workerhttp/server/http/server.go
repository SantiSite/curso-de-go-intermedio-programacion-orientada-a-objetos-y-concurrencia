package http

import (
	"github.com/SantiSite/workerhttp/types"
	"log"
	"net/http"

	"github.com/SantiSite/workerhttp/constants"
)

func Server(jobQueue chan types.Job) {
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(constants.Port, nil))
}
