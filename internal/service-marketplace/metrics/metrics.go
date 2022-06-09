package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	Requests            prometheus.Counter
	RequestErrors       prometheus.Counter
	CreateProduct       prometheus.Counter
	CreateProductErrors prometheus.Counter
	GetProduct          prometheus.Counter
	GetProductErrors    prometheus.Counter
	AddReview           prometheus.Counter
	AddReviewErrors     prometheus.Counter
	GetReviews          prometheus.Counter
	GetReviewsErrors    prometheus.Counter
	UpdateCart          prometheus.Counter
	UpdateCartErrors    prometheus.Counter
	GetCart             prometheus.Counter
	GetCartErrors       prometheus.Counter
}

func (m *metrics) RequestsInc() {
	m.Requests.Inc()
}
func (m *metrics) RequestErrorsInc() {
	m.RequestErrors.Inc()
}
func (m *metrics) CreateProductInc() {
	m.CreateProduct.Inc()
}
func (m *metrics) CreateProductErrorsInc() {
	m.CreateProductErrors.Inc()
}
func (m *metrics) GetProductInc() {
	m.GetProduct.Inc()
}
func (m *metrics) GetProductErrorsInc() {
	m.GetProductErrors.Inc()
}
func (m *metrics) AddReviewInc() {
	m.AddReview.Inc()
}
func (m *metrics) AddReviewErrorsInc() {
	m.AddReviewErrors.Inc()
}
func (m *metrics) GetReviewsInc() {
	m.GetReviews.Inc()
}
func (m *metrics) GetReviewsErrorsInc() {
	m.GetReviewsErrors.Inc()
}
func (m *metrics) UpdateCartInc() {
	m.UpdateCart.Inc()
}
func (m *metrics) UpdateCartErrorsInc() {
	m.UpdateCartErrors.Inc()
}
func (m *metrics) GetCartInc() {
	m.GetCart.Inc()
}
func (m *metrics) GetCartErrorsInc() {
	m.GetCartErrors.Inc()
}

func New() *metrics {
	r := &metrics{
		Requests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "Requests",
		}),
		RequestErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "RequestErrors",
		}),
		CreateProduct: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CreateProduct",
		}),
		CreateProductErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "CreateProductErrors",
		}),
		GetProduct: promauto.NewCounter(prometheus.CounterOpts{
			Name: "GetProduct",
		}),
		GetProductErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "GetProductErrors",
		}),
		AddReview: promauto.NewCounter(prometheus.CounterOpts{
			Name: "AddReview",
		}),
		AddReviewErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "AddReviewErrors",
		}),
		GetReviews: promauto.NewCounter(prometheus.CounterOpts{
			Name: "GetReviews",
		}),
		GetReviewsErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "GetReviewsErrors",
		}),
		UpdateCart: promauto.NewCounter(prometheus.CounterOpts{
			Name: "UpdateCart",
		}),
		UpdateCartErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "UpdateCartErrors",
		}),
		GetCart: promauto.NewCounter(prometheus.CounterOpts{
			Name: "GetCart",
		}),
		GetCartErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "GetCartErrors",
		}),
	}

	return r
}
