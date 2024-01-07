package main

import (
	"fmt"
	"github.com/KA-Ryzhkov/metrics-and-alerting/cmd/agent/metrics"
	"net/http"
	"time"
)

func dataRequest(m []metrics.Metric) []metrics.Metric {
	for i, v := range m {
		m[i] = v.Update()
	}
	return m
}

func generateUrl(m metrics.Metric, gauge bool) string {
	if gauge {
		url := fmt.Sprintf("http://localhost:8080/update/gauge/%s/%f/", m.Name, m.Gauge)
		return url
	} else {
		url := fmt.Sprintf("http://localhost:8080/update/counter/%s/%d/", m.Name, m.Counter)
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
	m := metrics.MetricStart(metrics.ListNameMetrics)
	for {
		m = dataRequest(m)

		for _, v := range m {
			url := generateUrl(v, true)
			fmt.Println("URL1:", url)
			err := sendPost(url)
			if err != nil {
				fmt.Println("Send gauge err:", err)
			}

			url = generateUrl(v, false)
			fmt.Println("URL2:", url)
			err = sendPost(url)
			if err != nil {
				fmt.Println("Send gauge err:", err)
			}
		}

		time.Sleep(2 * time.Second)
	}

}
