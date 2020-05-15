package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/session", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			latency(150*time.Millisecond, 600*time.Millisecond)
		case http.MethodGet:
			latency(350*time.Millisecond, 800*time.Millisecond)
		case http.MethodPut:
			latency(200*time.Millisecond, 300*time.Millisecond)
		}
		status(w)
	})
	router.HandleFunc("/api/v2/session", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			latency(800*time.Millisecond, 1200*time.Millisecond)
		case http.MethodGet:
			latency(100*time.Millisecond, 200*time.Millisecond)
		case http.MethodPut:
			latency(50*time.Millisecond, 100*time.Millisecond)
		}
		status(w)
	})
	router.HandleFunc("/api/v1/issue", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			latency(450*time.Millisecond, 2500*time.Millisecond)
		case http.MethodGet:
			latency(1000*time.Millisecond, 1200*time.Millisecond)
		case http.MethodPut:
			latency(10*time.Millisecond, 20000*time.Millisecond)
		}
		status(w)
	})
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Print(err)
	}
}

func latency(from, through time.Duration) {
	d := (through - from).Nanoseconds() / (rand.Int63n(100) + 1)
	time.Sleep(from + time.Duration(d))
}

func status(w http.ResponseWriter) {
	seeds := []int{200, 200, 200, 200, 404, 404, 500}
	w.WriteHeader(seeds[rand.Intn(len(seeds))])
}
