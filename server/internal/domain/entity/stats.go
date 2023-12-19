package entity

import (
	"github.com/prometheus/client_golang/prometheus"
)

//HitStats entity
type HitStats struct {
	Ok                  prometheus.Counter
	InternalServerError prometheus.Counter
	NotFoundError       prometheus.Counter
	URLMetric           prometheus.CounterVec
	Timing              prometheus.CounterVec
}
