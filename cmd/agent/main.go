package main

import (
	"flag"
	"fmt"
	"github.com/KA-Ryzhkov/metrics-and-alerting/cmd/agent/flags"
	"github.com/KA-Ryzhkov/metrics-and-alerting/cmd/agent/metrics"
	"net/http"
	"os"
	"strconv"
	"time"
)

func dataRequest(m []metrics.Metric) []metrics.Metric {
	for i, v := range m {
		m[i] = v.Update()
	}
	return m
}

func generateUrl(m metrics.Metric, gauge bool, addr string) string {
	if gauge {
		url := fmt.Sprintf("http://%s/update/gauge/%s/%f/", addr, m.Name, m.Gauge)
		return url
	} else {
		url := fmt.Sprintf("http://%s/update/counter/%s/%d/", addr, m.Name, m.Counter)
		return url
	}
}

func sendPost(url string) error {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "text/plain")
	req.ContentLength = 0
	req.Header.Add("Content-Length", "0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
	return nil
}

func main() {
	// Get parameters from environment variable
	addrEnv := os.Getenv("ADDRESS")
	reportEnv, err := strconv.Atoi(os.Getenv("REPORT_INTERVAL"))
	if err != nil {
		reportEnv = 0
	}
	pollEnv, err := strconv.Atoi(os.Getenv("POLL_INTERVAL"))
	if err != nil {
		pollEnv = 0
	}

	// Get parameters from command line argument (flag)
	addrFlag := new(flags.NetAddress)
	reportFlag := new(flags.ReportInterval)
	pollFlag := new(flags.PollInterval)
	flag.Var(addrFlag, "a", "Net address host:port")
	flag.Var(reportFlag, "r", "Report interval, integer")
	flag.Var(pollFlag, "p", "Poll interval, integer")
	flag.Parse()

	addr := parameterPriorityAddr(addrEnv, addrFlag.String())
	report := parameterPriorityTime(reportEnv, reportFlag.TimeInterval, 10)
	poll := parameterPriorityTime(pollEnv, pollFlag.TimeInterval, 2)

	fmt.Printf("Start agent, with parameters: %s, report: %s , poll: %s", addr, report, poll)

	metricsList := metrics.MetricStart(metrics.ListNameMetrics)

	metricsChan := make(chan []metrics.Metric)
	go func() {
		for {
			metricsChan <- dataRequest(metricsList)
			time.Sleep(poll)
		}
	}()

	for {
		for _, v := range <-metricsChan {
			url := generateUrl(v, true, addr)
			fmt.Println("URL gauge:", url)
			err := sendPost(url)
			if err != nil {
				fmt.Println("Send gauge err:", err)
			}

			url = generateUrl(v, false, addr)
			fmt.Println("URL counter:", url)
			err = sendPost(url)
			if err != nil {
				fmt.Println("Send counter err:", err)
			}
		}

		time.Sleep(report)
	}

}

func parameterPriorityAddr(env string, s string) string {
	if env != "" {
		return env
	}
	return s
}

func parameterPriorityTime(env int, flag time.Duration, t int) time.Duration {
	if env > 0 {
		return time.Duration(env)
	} else if flag > 0 {
		return flag
	} else {
		return time.Duration(t) * time.Second // Default value
	}
}
