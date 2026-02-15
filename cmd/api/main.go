package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VaLTrexx/go-job-queue/internal/job"
	"github.com/VaLTrexx/go-job-queue/internal/redis"
	"github.com/VaLTrexx/go-job-queue/internal/worker"
)

func main() {
	client := redis.NewClient()

	queue := redis.NewRedisQueue(client, "job_queue")
	store := redis.NewRedisStore(client)

	w1 := worker.Worker{ID: 1, Queue: queue, Store: store}
	w2 := worker.Worker{ID: 2, Queue: queue, Store: store}

	go w1.Start()
	go w2.Start()

	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		j := job.NewJob("email")

		store.Save(j)
		queue.Enqueue(j)

		json.NewEncoder(w).Encode(map[string]string{
			"job_id": j.ID,
			"status": string(j.Status),
		})
	})

	http.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		j, ok, err := store.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(j)
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
