package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	systemID    int = 1
	zoneID      int = 0
	url         string
	reqInterval int = 10
)

func worker(ctx context.Context, ticker <-chan time.Time, fn func()) {
	defer func() { log.Printf("worker stopped") }()

	for {
		select {
		case <-ticker:
			fn()
		case <-ctx.Done():
			return
		}
	}
}

func work(url string, systemID, zoneID int) func() {
	return func() {
		hvac, err := DoHVAC(url, systemID, zoneID)
		if err != nil {
			zap.S().Error(err)
		}
		for _, zone := range hvac.Zones {
			UpdateTemperatureMetric(zone.Name, zone.RoomTemp)
			UpdateHumidityMetric(zone.Name, int(zone.Humidity))
		}
	}
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	flag.IntVar(&systemID, "system-id", 1, "system id. default to 0")
	flag.IntVar(&zoneID, "zone-id", 0, "Default to 0 (all zones)")
	flag.StringVar(&url, "url", "airzone:3000", "airzone local url. Example: 192.168.1.1:3000")
	flag.IntVar(&reqInterval, "tick", 10, "fetching data interval.Defaults to 10 seconds")
	flag.Parse()

	prometheus.MustRegister(tempCounter)
	prometheus.MustRegister(humidityCounter)

	t := time.NewTicker(time.Duration(reqInterval) * time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	zap.S().Infof("start worker.ticking every %d seconds\n", reqInterval)

	go worker(ctx, t.C, work(url, systemID, zoneID))

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)

	<-done

	cancel()
}
