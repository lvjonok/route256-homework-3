package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	Requests           prometheus.Counter
	RequestErrors      prometheus.Counter
	CreateOrder        prometheus.Counter
	CreateOrderErrors  prometheus.Counter
	CheckStatus        prometheus.Counter
	CheckStatusErrors  prometheus.Counter
	UpdateStatus       prometheus.Counter
	UpdateStatusErrors prometheus.Counter
}

func (m *metrics) RequestsInc() {
	m.Requests.Inc()
}
func (m *metrics) RequestErrorsInc() {
	m.RequestErrors.Inc()
}
func (m *metrics) CreateOrderInc() {
	m.CreateOrder.Inc()
}
func (m *metrics) CreateOrderErrorsInc() {
	m.CreateOrderErrors.Inc()
}
func (m *metrics) CheckStatusInc() {
	m.CheckStatus.Inc()
}
func (m *metrics) CheckStatusErrorsInc() {
	m.CheckStatusErrors.Inc()
}
func (m *metrics) UpdateStatusInc() {
	m.UpdateStatus.Inc()
}
func (m *metrics) UpdateStatusErrorsInc() {
	m.UpdateStatusErrors.Inc()
}

func New() *metrics {
	m := metrics{
		Requests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "Requests",
		}),
		RequestErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "RequestErrors",
		}),
		CreateOrder: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CreateOrder",
		}),
		CreateOrderErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CreateOrderErrors",
		}),
		CheckStatus: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CheckStatus",
		}),
		CheckStatusErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CheckStatusErrors",
		}),
		UpdateStatus: promauto.NewCounter(prometheus.CounterOpts{
			Name: "UpdateStatus",
		}),
		UpdateStatusErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "UpdateStatusErrors",
		}),
	}

	return &m
}
