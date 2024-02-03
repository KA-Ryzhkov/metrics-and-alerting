package main

import (
	"flag"
	"fmt"
	"github.com/KA-Ryzhkov/metrics-and-alerting/cmd/agent/metrics"
	"github.com/KA-Ryzhkov/metrics-and-alerting/cmd/server/flags"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"strconv"
)

var MemStorage []metrics.Metric

func updateHandle(res http.ResponseWriter, req *http.Request) {
	t, n, v := chi.URLParam(req, "type"), chi.URLParam(req, "name"), chi.URLParam(req, "value")
	flag := true
	for i, value := range MemStorage {
		if n == value.Name {

			switch t {
			case "gauge":
				num, err := strconv.ParseFloat(v, 64)
				if err != nil {
					res.WriteHeader(http.StatusNotFound)
					fmt.Println("неправильно значение метрики")
				}
				res.WriteHeader(http.StatusOK)
				res.Write([]byte(v))
				MemStorage[i].Gauge = num
				flag = false

			case "counter":
				num, err := strconv.ParseInt(v, 0, 64)
				if err != nil {
					res.WriteHeader(http.StatusNotFound)
					fmt.Println("неправильно значение метрики")
				}
				res.WriteHeader(http.StatusOK)
				res.Write([]byte(v))
				MemStorage[i].Counter = num
				flag = false

			default:
				flag = false
				res.WriteHeader(http.StatusNotFound)
				fmt.Println("неправильно значение типа")
			}
		}

	}
	if flag {
		res.WriteHeader(http.StatusNotFound)
		fmt.Println("нет такой метрики")
	}
}

func printHandle(res http.ResponseWriter, req *http.Request) {
	t, n := chi.URLParam(req, "type"), chi.URLParam(req, "name")
	text := ""
	for _, value := range MemStorage {
		if value.Name == n {

			switch t {
			case "gauge":
				text = fmt.Sprintf("%f", value.Gauge)
			case "counter":
				text = fmt.Sprintf("%d", value.Counter)
			default:
				res.WriteHeader(http.StatusNotFound)
				fmt.Println("неправильно значение типа")
				return
			}
		}
	}
	res.Write([]byte(text))
	res.WriteHeader(http.StatusOK)
}

func allMetricsHandle(res http.ResponseWriter, req *http.Request) {
	text := ""
	for _, value := range MemStorage {
		text += fmt.Sprintf("Name: %s\t\tGauge: %f\t\tCounter %d\n", value.Name, value.Gauge, value.Counter)

	}
	res.Write([]byte(text))
	res.WriteHeader(http.StatusOK)
}

func main() {
	// Get parameters from environment variable
	address := os.Getenv("ADDRESS")

	// Get parameters from command line argument (flag)
	addr := new(flags.NetAddress)
	flag.Var(addr, "a", "Net address host:port")
	flag.Parse()

	MemStorage = metrics.MetricStart(metrics.ListNameMetrics)
	r := chi.NewRouter()

	r.Get("/update/{type}/{name}/{value}/", updateHandle)
	r.Get("/value/{type}/{name}/", printHandle)
	r.Get("/", allMetricsHandle)

	if address != "" {
		fmt.Println("Server start, with parameters: http://" + address)
		log.Fatal(http.ListenAndServe(address, r))
	} else {
		fmt.Println("Server start, with parameters: http://" + addr.String())
		log.Fatal(http.ListenAndServe(addr.String(), r))
	}

}
