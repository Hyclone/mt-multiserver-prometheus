package main

import (
	"log"
	"net/http"
	"time"

	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			playerCount.Set(float64(len(proxy.Players())))
			time.Sleep(5 * time.Second)
		}
	}()
}

var (
	playerCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "multiserver_playercount",
		Help: "The total number of players",
	})
)


func init() {
	log.Println("[mt-multiserver-prometheus] Starting HTTP Prometheus endpoint on port 2112...")
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":2112", nil)
}