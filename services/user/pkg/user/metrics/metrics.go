package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type User struct {
	opsProcessed prometheus.Counter
	usersSaved   prometheus.Counter
}

func New() *User {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "royalafg_user_processed_ops_total",
		Help: "The total number of processed events",
	})

	usersSaved := promauto.NewCounter(prometheus.CounterOpts{
		Name: "royalafg_user_users_saved",
		Help: "The total number of users saved to the database",
	})
	return &User{
		opsProcessed: opsProcessed,
		usersSaved:   usersSaved,
	}
}

func (m *User) SavedUser() {
	m.usersSaved.Inc()
	m.Operation()
}

func (m *User) Operation() {
	m.opsProcessed.Inc()
}
