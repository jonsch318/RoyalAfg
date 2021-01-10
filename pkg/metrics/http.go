package metrics

import (
	"time"
)

//HTTP is the metrics for an HTTP app
type HTTP struct {
	Handler    string
	Method     string
	StatusCode int
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   time.Duration
}

//NewHTTP creates a new HTTP app metrics
func NewHTTP(handler, method string) *HTTP {
	return &HTTP{
		Handler: handler,
		Method:  method,
	}
}

func (h *HTTP) Start() {
	h.StartedAt = time.Now()
}

func (h *HTTP) End(status int) {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt)
	h.StatusCode = status
}
