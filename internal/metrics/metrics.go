package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

var (
	labels = []string{
		"room",
	}

	tempCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "airzone",
			Name:      "temperature",
			Help:      "Room temperature",
		},
		labels,
	)

	humidityCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "airzone",
			Name:      "humidity",
			Help:      "Room humidity",
		},
		labels,
	)

	zoneStateCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "airzone",
			Name:      "state",
			Help:      "Zone state",
		},
		labels,
	)

	modeCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Subsystem: "airzone",
			Name:      "mode",
			Help:      "Mode",
		},
	)
)

func init() {
	prometheus.MustRegister(tempCounter)
	prometheus.MustRegister(humidityCounter)
	prometheus.MustRegister(zoneStateCounter)
	prometheus.MustRegister(modeCounter)
}

func UpdateTemperatureMetric(room string, temp float64) {
	labels := prometheus.Labels{
		"room": room,
	}
	tempCounter.With(labels).Set(temp)
	zap.S().Debugw("temperature metrics updated", "value", temp)
}

func UpdateHumidityMetric(room string, humidity int) {
	labels := prometheus.Labels{
		"room": room,
	}
	humidityCounter.With(labels).Set(float64(humidity))
	zap.S().Debugw("humidity metrics updated", "value", humidity)
}

func UpdateZoneStateMetric(room string, state int) {
	labels := prometheus.Labels{
		"room": room,
	}
	zoneStateCounter.With(labels).Set(float64(state))
	zap.S().Debugw("state metrics updated", "value", state)
}

func UpdateModeMetric(mode int) {
	modeCounter.Set(float64(mode))
	zap.S().Debugw("mode metrics updated", "value", mode)
}
