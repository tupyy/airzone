package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
	"github.com/tupyy/airzone/internal/metrics"
	"github.com/tupyy/airzone/internal/worker"
	"go.uber.org/zap"
)

const (
	allZones = 0
)

var (
	httpPort    int = 8080
	reqInterval int = 10
)

func work(url string, systemID, zoneID int) func() error {
	return func() error {
		hvac, err := hvac.GetData(url, systemID, zoneID)
		if err != nil {
			return err
		}
		for _, zone := range hvac.Zones {
			metrics.UpdateTemperatureMetric(zone.Name, zone.RoomTemp)
			metrics.UpdateHumidityMetric(zone.Name, int(zone.Humidity))
		}
		return nil
	}
}

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Start the metric server",
	Long: `This command starts a http server having single endpoint /metrics.
These metrics can be scraped by prometheus.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}
		defer logger.Sync()

		undo := zap.ReplaceGlobals(logger)
		defer undo()

		t := time.NewTicker(time.Duration(reqInterval) * time.Second)

		ctx, cancel := context.WithCancel(context.Background())

		zap.S().Infof("start worker.ticking every %d seconds\n", reqInterval)

		go worker.Worker(ctx, t.C, work(common.Host, common.SystemID, common.AllZones))

		http.Handle("/metrics", promhttp.Handler())
		go http.ListenAndServe(":8080", nil)

		zap.S().Infof("metrics server started. Endpoint /metrics on localhost:%d", httpPort)

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, os.Kill)

		<-done

		cancel()
	},
}

func init() {
	RootCmd.AddCommand(metricsCmd)

	metricsCmd.Flags().IntVarP(&httpPort, "port", "", 8080, "http port of the metric server")
	metricsCmd.Flags().IntVarP(&reqInterval, "ticker", "", 10, "interval of request to airzone server")
}
