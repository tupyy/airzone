package main

import "github.com/prometheus/client_golang/prometheus"

var (
	labels = []string{
		"room",
	}

	tempCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature",
			Help: "Room temperature",
		},
		labels,
	)

	humidityCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "humidity",
			Help: "Room humidity",
		},
		labels,
	)
)

func UpdateTemperatureMetric(room string, temp float64) {
	labels := prometheus.Labels{
		"room": room,
	}
	tempCounter.With(labels).Set(temp)
}

func UpdateHumidityMetric(room string, humidity int) {
	labels := prometheus.Labels{
		"room": room,
	}
	humidityCounter.With(labels).Set(float64(humidity))
}
