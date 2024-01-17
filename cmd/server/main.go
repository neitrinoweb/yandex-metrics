package main

import (
	"net/http"
)

// TODO: Move packages to internal
// TODO: move server init to server.go
var storage *InMemoryMetricStorage

func main() {
	// TODO: Init storage
	storage = InitInMemoryMetricStorage()

	// Test storage init data
	storage.UpdateCounter("init_counter_metric", 0)
	storage.UpdateGauge("init_gauge_metric", 0.0)

	// Init mux
	mux := http.NewServeMux()

	// Metric update endpoints
	mux.HandleFunc(`/update/`, updateMetricHandler)

	// Init server
	err := http.ListenAndServe(`localhost:8080`, mux)
	if err != nil {
		panic(err)
	}
}
