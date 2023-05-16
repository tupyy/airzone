package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(tempCounter)
	prometheus.MustRegister(humidityCounter)
	prometheus.MustRegister(zoneStateCounter)
	prometheus.MustRegister(modeCounter)
}

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

	zoneStateCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "state",
			Help: "Zone state",
		},
		labels,
	)

	modeCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "mode",
			Help: "Mode",
		},
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

func UpdateZoneStateMetric(room string, state int) {
	labels := prometheus.Labels{
		"room": room,
	}
	zoneStateCounter.With(labels).Set(float64(state))
}

func UpdateModeMetric(mode int) {
	modeCounter.Set(float64(mode))
}
