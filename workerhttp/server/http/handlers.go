package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SantiSite/workerhttp/types"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan types.Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
	}

	job := types.Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	jobQueue <- job
	w.WriteHeader(http.StatusNoContent)
}
