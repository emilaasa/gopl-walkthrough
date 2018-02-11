package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type WorkRequest struct {
	Name  string
	Delay time.Duration
}

var WorkQueue = make(chan WorkRequest, 100)

// Collector responsible for running sanity checks on the
// incoming work requests, and alert the client if their work request does not fit
// within whatever acceptable boundaries you define. We also do not want our
// collector to hold an open network connection for any longer than it has to
func Collector(w http.ResponseWriter, r *http.Request) {
	// Only allow POST because this is a RPC
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}
	if delay < 1 || delay > 10 {
		http.Error(w, "Delay must be between 1 and 10 seconds: "+err.Error(), http.StatusBadRequest)
		return
	}
	name, err := r.FormValue("name")
	if err != nil {
		http.Error(w, "Provide a name: "+err.Error(), http.StatusBadRequest)
		return
	}
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	work := WorkRequest{Name: name, Delay: delay}
	WorkQueue <- work
	fmt.Println("Work request queued")
	w.WriteHeader(http.StatusCreated)
	return
}

func NewWorker(id int, workerQueue chan chan workRequest) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
	return worker
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves onto the worker queue
			w.WorkerQueue <- w.Work
			select {
			case work := <-w.Work:
				fmt.Printf("worker%d: Received work request, delaying for %s seconds\n", w.ID, work.Delay)
				time.Sleep(work.Delay)
				fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

			case <-w.QuitChan:
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func StartDispatcher(nworkers int) {
	WorkerQueue := make(chan chan WorkRequest, nworkers)

	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Recieved work request")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
