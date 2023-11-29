package entity

import (
	"github.com/prometheus/client_golang/prometheus"
)

type HitStats struct {
	Ok                  prometheus.Counter
	InternalServerError prometheus.Counter
	NotFoundError       prometheus.Counter
	UrlMetric           prometheus.CounterVec
	Timing              prometheus.CounterVec
}
