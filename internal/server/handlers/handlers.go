package handlers

import (
	// "fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"yandex-metrics/internal/server/storage"
)

func UpdateMetric(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// TODO: Check for Content-Type
		// if r.Header["Content-Type"]

		// Check for valid URL params
		url := r.URL.RequestURI()
		segments := strings.Split(strings.TrimPrefix(url, "/"), "/")

		log.Printf("Checking request \"%s\" for valid data", url)
		// Check if there's no more then 4 segments in url path
		if len(segments) != 4 || segments[3] == "" {
			log.Printf("URL \"%s\" is not valid", url)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		// Split URL request into params for metric update
		metricType := segments[1]
		metricName := segments[2]
		metricValue := segments[3]

		metricNamePattern := regexp.MustCompile(`^[a-zA-Z_:][a-zA-Z0-9_:]*$`)
		if !metricNamePattern.MatchString(segments[2]) {
			log.Printf("\"%s\" metric name is nov valid", segments[2])
			http.Error(w, "Invalid metric name", http.StatusBadRequest)
			return
		}

		// Update metric
		switch metricType {
		case "gauge":
			mv, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, "Unsupported value for gauge metric", http.StatusBadRequest)
				return
			}
			err = storage.UpdateGauge(metricName, mv)
			if err != nil {
				http.Error(w, "Failed to update gauge metric", http.StatusBadRequest)
				return
			}
			log.Printf("Updated metric \"%s\" sucessfully, new value: \"%f\"", segments[2], mv)
			return

		case "counter":
			mv, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(w, "Unsupported value for counter metric", http.StatusBadRequest)
			}

			err = storage.UpdateCounter(metricName, mv)
			if err != nil {
				http.Error(w, "Failed to update counter metric", http.StatusBadRequest)
				return
			}
			log.Printf("Updated metric \"%s\" sucessfully, new value: \"%d\"", segments[2], mv)
			return

		default:
			http.Error(w, "Bad Request, metric type is not supported", http.StatusBadRequest)
			log.Printf("Metric type \"%s\" is not supported", segments[2])
			return
		}
	}
}
