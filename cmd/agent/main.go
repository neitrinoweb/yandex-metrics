package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	HostAddress    string = "localhost"
	HostPort       int    = 8080
	PollInterval   int    = 2
	reportInterval int    = 10

	// Paths
	UpdatePath        string = "/update/"
	GaugeMetricPath   string = "/gauge/"
	CounterMetricPath string = "/Counter/"
)

type gaugeMetric struct {
	Name  string
	Value float64
}

type counterMetirc struct {
	Name  string
	Value int
}

func main() {

	fmt.Println("Init")
	// Check server hostname

	hostURL := url.URL{
		Scheme: "http",
		Host:   strings.Join([]string{HostAddress, strconv.Itoa(HostPort)}, ":"),
	}

	// Init metrics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// PollCount += 1

	for {
		gaugeMetrics := []gaugeMetric{
			{Name: "Alloc", Value: float64(m.Alloc)},
			{Name: "BuckHashSys", Value: float64(m.BuckHashSys)},
			{Name: "Frees", Value: float64(m.Frees)},
			{Name: "GCCPUFraction", Value: float64(m.GCCPUFraction)},
			{Name: "GCSys", Value: float64(m.GCSys)},
			{Name: "HeapAlloc", Value: float64(m.HeapAlloc)},
			{Name: "HeapIdle", Value: float64(m.HeapIdle)},
			{Name: "HeapObjects", Value: float64(m.HeapObjects)},
			{Name: "HeapReleased", Value: float64(m.HeapReleased)},
			{Name: "HeapSys", Value: float64(m.HeapSys)},
			{Name: "LastGC", Value: float64(m.LastGC)},
			{Name: "Lookups", Value: float64(m.Lookups)},
			{Name: "MCacheInuse", Value: float64(m.MCacheInuse)},
			{Name: "MCacheSys", Value: float64(m.MCacheSys)},
			{Name: "MSpanInuse", Value: float64(m.MSpanInuse)},
			{Name: "MSpanSys", Value: float64(m.MSpanSys)},
			{Name: "Mallocs", Value: float64(m.Mallocs)},
			{Name: "NextGC", Value: float64(m.NextGC)},
			{Name: "NumForcedGC", Value: float64(m.NumForcedGC)},
			{Name: "NumGC", Value: float64(m.NumGC)},
			{Name: "OtherSys", Value: float64(m.OtherSys)},
			{Name: "PauseTotalNs", Value: float64(m.PauseTotalNs)},
			{Name: "StackInuse", Value: float64(m.StackInuse)},
			{Name: "StackSys", Value: float64(m.StackSys)},
			{Name: "Sys", Value: float64(m.Sys)},
			{Name: "TotalAlloc", Value: float64(m.TotalAlloc)},
		}

		for _, metric := range gaugeMetrics {
			hostURL.Path, _ = url.JoinPath("/update/", "gauge", "/", metric.Name, "/", strconv.FormatFloat(metric.Value, 'g', -1, 64))

			fmt.Println("Request sent:", hostURL.String())
			resp, err := http.Post(hostURL.String(), "", nil)
			if err != nil {
				log.Println("Error: ", err)
				return
			}
			defer resp.Body.Close()
		}

		time.Sleep(5 * time.Second)
	}

}
