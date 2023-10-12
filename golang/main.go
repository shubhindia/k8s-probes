package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Rate metric
var requestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of HTTP requests",
}, []string{"method"})

// Errors metric
var errorsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_errors_total",
	Help: "Total number of HTTP errors",
}, []string{"method"})

// Duration metric
var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_request_duration_seconds",
	Help:    "HTTP request duration in seconds",
	Buckets: prometheus.ExponentialBuckets(0.05, 1.5, 10),
}, []string{"method"})

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
	startTime := time.Now()
	time.Sleep(5 * time.Second)
	defer func() {
		duration := time.Since(startTime).Seconds()
		method := r.Method
		requestsTotal.WithLabelValues(method).Inc()
		requestDuration.WithLabelValues(method).Observe(duration)
	}()
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
	prometheus.MustRegister(pongCounter, requestsTotal, errorsTotal, requestDuration)
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
