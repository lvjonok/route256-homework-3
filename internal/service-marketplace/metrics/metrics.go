package metrics

import (
	"log"

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
}

func (m *metrics) RequestsInc() {
	log.Printf("requests inc")
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

func New() *metrics {
	r := &metrics{
		Requests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "processed_requests",
		}),
		RequestErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "processed_requests_errors",
		}),
		CreateProduct: promauto.NewCounter(prometheus.CounterOpts{
			Name: "create_product",
		}),
		CreateProductErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "create_product_errors",
		}),
		GetProduct: promauto.NewCounter(prometheus.CounterOpts{
			Name: "get_product",
		}),
		GetProductErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "get_product_errors",
		}),
		AddReview: promauto.NewCounter(prometheus.CounterOpts{
			Name: "add_review",
		}),
		AddReviewErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "add_review_errors",
		}),
		GetReviews: promauto.NewCounter(prometheus.CounterOpts{
			Name: "get_reviews",
		}),
		GetReviewsErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "get_reviews_errors",
		}),
		UpdateCart: promauto.NewCounter(prometheus.CounterOpts{
			Name: "update_cart",
		}),
		UpdateCartErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "update_cart_errors",
		}),
	}

	return r
}
