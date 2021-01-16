package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type Service struct {
	httpHistogram *prometheus.HistogramVec
}

func NewMetricsService() (*Service, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "request_duration_seconds",
		Namespace: "http",
		Help:      "The latency of the http response",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	ret := &Service{
		httpHistogram: http,
	}

	if err := prometheus.Register(ret.httpHistogram); err != nil {
		switch e := err.(type) {
		case prometheus.AlreadyRegisteredError:
			ret.httpHistogram = e.ExistingCollector.(*prometheus.HistogramVec)
			break
		default:
			return nil, err
		}
	}

	return ret, nil
}

func (s *Service) SaveHTTP(h *HTTP) {
	s.httpHistogram.WithLabelValues(h.Handler, h.Method, strconv.Itoa(h.StatusCode)).Observe(h.Duration.Seconds())
}
