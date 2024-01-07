package metrics

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

func MetricStart(ListNameMetrics []string) []Metric {
	var metricList []Metric
	for _, NameMetrics := range ListNameMetrics {
		metricList = append(metricList, Metric{Name: NameMetrics})
	}
	return metricList
}
