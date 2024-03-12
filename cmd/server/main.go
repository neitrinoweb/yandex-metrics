package main

import (
	"net/http"
	"yandex-metrics/internal/server/handlers"
	"yandex-metrics/internal/server/storage"
)

func main() {
	storageMemory := storage.InitInMemoryMetricStorage()

	// Init data (storage pre-flight checks)
	storageMemory.UpdateCounter("init_counter_metric", 0)
	storageMemory.UpdateGauge("init_gauge_metric", 0.0)

	// Endpoints
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, handlers.UpdateMetric(storageMemory))

	// Init server
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
