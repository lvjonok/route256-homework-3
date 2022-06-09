package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	Requests              prometheus.Counter
	RequestsErrors        prometheus.Counter
	RegisterProduct       prometheus.Counter
	RegisterProductErrors prometheus.Counter
	CheckProducts         prometheus.Counter
	CheckProductsErrors   prometheus.Counter
	BookProducts          prometheus.Counter
	BookProductsErrors    prometheus.Counter
	UnbookProducts        prometheus.Counter
	UnbookProductsErrors  prometheus.Counter
}

func (m *metrics) RequestsInc() {
	m.Requests.Inc()
}
func (m *metrics) RequestsErrorsInc() {
	m.RequestsErrors.Inc()
}
func (m *metrics) RegisterProductInc() {
	m.RegisterProduct.Inc()
}
func (m *metrics) RegisterProductErrorsInc() {
	m.RegisterProductErrors.Inc()
}
func (m *metrics) CheckProductsInc() {
	m.CheckProducts.Inc()
}
func (m *metrics) CheckProductsErrorsInc() {
	m.CheckProductsErrors.Inc()
}
func (m *metrics) BookProductsInc() {
	m.BookProducts.Inc()
}
func (m *metrics) BookProductsErrorsInc() {
	m.BookProductsErrors.Inc()
}
func (m *metrics) UnbookProductsInc() {
	m.UnbookProducts.Inc()
}
func (m *metrics) UnbookProductsErrorsInc() {
	m.UnbookProductsErrors.Inc()
}

func New() *metrics {
	m := metrics{
		Requests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "Requests",
		}),
		RequestsErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "RequestsErrors",
		}),
		RegisterProduct: promauto.NewCounter(prometheus.CounterOpts{
			Name: "RegisterProduct",
		}),
		RegisterProductErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "RegisterProductErrors",
		}),
		CheckProducts: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CheckProducts",
		}),
		CheckProductsErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CheckProductsErrors",
		}),
		BookProducts: promauto.NewCounter(prometheus.CounterOpts{
			Name: "BookProducts",
		}),
		BookProductsErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "BookProductsErrors",
		}),
		UnbookProducts: promauto.NewCounter(prometheus.CounterOpts{
			Name: "UnbookProducts",
		}),
		UnbookProductsErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "UnbookProductsErrors",
		}),
	}

	return &m
}
