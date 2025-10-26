package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	registry          *prometheus.Registry
	errors            *prometheus.CounterVec
	httpRequestsTotal *prometheus.CounterVec
	httpStatusCodes   *prometheus.CounterVec
}

func New() *Metric {
	r := prometheus.NewRegistry()
	m := &Metric{
		registry: r,
		errors: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "app",
				Subsystem: "errors",
				Name:      "total",
				Help:      "Total number of errors",
			},
			[]string{"service", "action"},
		),
		httpRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "http",
				Subsystem: "requests",
				Name:      "total",
				Help:      "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		httpStatusCodes: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "http",
				Subsystem: "status_codes",
				Name:      "total",
				Help:      "Total number of HTTP status codes",
			},
			[]string{"status"},
		),
	}

	r.MustRegister(
		m.errors,
		m.httpRequestsTotal,
		m.httpStatusCodes,
	)

	return m
}

func (m *Metric) RecordError(service, action string) {
	fmt.Println("got in error")
	m.errors.WithLabelValues(service, action).Inc()
}

func (m *Metric) RecordHTTPRequest(method, path, status string) {
	fmt.Println("got in HTTPRequest")
	m.httpRequestsTotal.WithLabelValues(method, path, status).Inc()
	m.httpStatusCodes.WithLabelValues(status).Inc()
}

func (m *Metric) Registry() *prometheus.Registry {
	fmt.Println("got in 2")
	return m.registry
}
