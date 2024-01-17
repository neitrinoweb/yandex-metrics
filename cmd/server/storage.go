package main

import (
	"log"
)

// TODO:
// - [ ] Универсальный интерфейс для MetricStorage

// *********************************************
// *                                           *
// *        IN-MEMORY METRIC STORAGE           *
// *                                           *
// *********************************************

type InMemoryMetricStorage struct {
	Counters map[string]int64
	Gauges   map[string]float64
}

func (storage *InMemoryMetricStorage) UpdateGauge(name string, value float64) error {
	storage.Gauges[name] = value
	log.Printf("Metric \"%s\" of type 'Gauge' is updated with value %f", name, value)
	return nil
}

func (storage *InMemoryMetricStorage) UpdateCounter(name string, value int64) error {
	if _, ok := storage.Counters[name]; !ok {
		storage.Counters[name] = 0
	}
	storage.Counters[name] += value
	log.Printf("Metric \"%s\" of type 'Counter' is updated with value %d", name, value)
	return nil
}

func InitInMemoryMetricStorage() *InMemoryMetricStorage {
	return &InMemoryMetricStorage{
		Counters: make(map[string]int64),
		Gauges:   make(map[string]float64),
	}
}

// TODO:
// *********************************************
// *                                           *
// *        DATABASE METRIC STORAGE            *
// *                                           *
// *********************************************

type DatabaseStorage struct{}
