package metrics

import (
	"fmt"
	"runtime"
)

var ListNameMetrics = []string{
	"Alloc",
	"BuckHashSys",
	"Frees",
	"GCCPUFraction",
	"GCSys",
	"HeapAlloc",
	"HeapIdle",
	"HeapInuse",
	"HeapObjects",
	"HeapReleased",
	"HeapSys",
	"LastGC",
	"Lookups",
	"MCacheInuse",
	"MCacheSys",
	"MSpanInuse",
	"MSpanSys",
	"Mallocs",
	"NextGC",
	"NumForcedGC",
	"NumGC",
	"OtherSys",
	"PauseTotalNs",
	"StackInuse",
	"StackSys",
	"Sys",
	"TotalAlloc",
}

type Metric struct {
	Name    string  // имя
	Gauge   float64 // значение
	Counter int64   // количество запросов
}

func (m Metric) Update() Metric {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	switch m.Name {
	case "Alloc":
		m.Gauge = float64(memStats.Alloc)
		m.Counter += 1
	case "BuckHashSys":
		m.Gauge = float64(memStats.BuckHashSys)
		m.Counter += 1
	case "Frees":
		m.Gauge = float64(memStats.Frees)
		m.Counter += 1
	case "GCCPUFraction":
		m.Gauge = memStats.GCCPUFraction
		m.Counter += 1
	case "GCSys":
		m.Gauge = float64(memStats.GCSys)
		m.Counter += 1
	case "HeapAlloc":
		m.Gauge = float64(memStats.HeapAlloc)
		m.Counter += 1
	case "HeapIdle":
		m.Gauge = float64(memStats.HeapIdle)
		m.Counter += 1
	case "HeapInuse":
		m.Gauge = float64(memStats.HeapInuse)
		m.Counter += 1
	case "HeapObjects":
		m.Gauge = float64(memStats.HeapObjects)
		m.Counter += 1
	case "HeapReleased":
		m.Gauge = float64(memStats.HeapReleased)
		m.Counter += 1
	case "HeapSys":
		m.Gauge = float64(memStats.HeapSys)
		m.Counter += 1
	case "LastGC":
		m.Gauge = float64(memStats.LastGC)
		m.Counter += 1
	case "Lookups":
		m.Gauge = float64(memStats.Lookups)
		m.Counter += 1
	case "MCacheInuse":
		m.Gauge = float64(memStats.MCacheInuse)
		m.Counter += 1
	case "MCacheSys":
		m.Gauge = float64(memStats.MCacheSys)
		m.Counter += 1
	case "MSpanInuse":
		m.Gauge = float64(memStats.MSpanInuse)
		m.Counter += 1
	case "MSpanSys":
		m.Gauge = float64(memStats.MSpanSys)
		m.Counter += 1
	case "Mallocs":
		m.Gauge = float64(memStats.Mallocs)
		m.Counter += 1
	case "NextGC":
		m.Gauge = float64(memStats.NextGC)
		m.Counter += 1
	case "NumForcedGC":
		m.Gauge = float64(memStats.NumForcedGC)
		m.Counter += 1
	case "NumGC":
		m.Gauge = float64(memStats.NumGC)
		m.Counter += 1
	case "OtherSys":
		m.Gauge = float64(memStats.OtherSys)
		m.Counter += 1
	case "PauseTotalNs":
		m.Gauge = float64(memStats.PauseTotalNs)
		m.Counter += 1
	case "StackInuse":
		m.Gauge = float64(memStats.StackInuse)
		m.Counter += 1
	case "StackSys":
		m.Gauge = float64(memStats.StackSys)
		m.Counter += 1
	case "Sys":
		m.Gauge = float64(memStats.Sys)
		m.Counter += 1
	case "TotalAlloc":
		m.Gauge = float64(memStats.TotalAlloc)
		m.Counter += 1
	default:
		fmt.Printf("Параметр %s не обрабатывается.\n", m)
	}
	return m
}

func MetricStart(ListNameMetrics []string) []Metric {
	var metricList []Metric
	for _, NameMetrics := range ListNameMetrics {
		metricList = append(metricList, Metric{Name: NameMetrics})
	}
	return metricList
}
