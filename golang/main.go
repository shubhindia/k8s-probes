package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pongCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
		ConstLabels: map[string]string{
			"handler": "ping",
		},
	},
)

func pong(w http.ResponseWriter, r *http.Request) {
	pongCounter.Inc()
	w.Header().Set("Content-Type", "application/json")
	hostname, _ := os.Hostname()
	resp := map[string]string{
		"message": fmt.Sprintf("pong from %s", hostname),
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hostname, _ := os.Hostname()
	resp := map[string]string{
		"message": fmt.Sprintf("Hello from %s", hostname),
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func handleRequests() {
	httpPort := 8080
	prometheus.MustRegister(pongCounter)
	http.HandleFunc("/healthz", homePage)
	http.HandleFunc("/ping", pong)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("listening on %v\n", httpPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	handleRequests()
}
