package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			fmt.Println("Increasing")
			playerCount.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	playerCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "multiserver_playercount",
		Help: "The total number of players",
	})
)


func init() {
	fmt.Println("Starting HTTP prometheus endpoint on port 2112...")
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":2112", nil)
}