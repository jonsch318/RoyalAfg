package mw

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/metrics"
)

type MetricsMW struct {
	metrics metrics.UseCase
}

func Metrics(metrics metrics.UseCase) *MetricsMW {
	return &MetricsMW{
		metrics: metrics,
	}
}

func (h *MetricsMW) MW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		requestMetric := metrics.NewHTTP(r.URL.Path, r.Method)
		requestMetric.Start()

		next.ServeHTTP(rw, r)

		requestMetric.End()
	})
}
