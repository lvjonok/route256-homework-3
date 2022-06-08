package service

type Metrics interface {
	RequestsInc()
	RequestErrorsInc()
	CreateProductInc()
	CreateProductErrorsInc()
	GetProductInc()
	GetProductErrorsInc()
	AddReviewInc()
	AddReviewErrorsInc()
	GetReviewsInc()
	GetReviewsErrorsInc()
	UpdateCartInc()
	UpdateCartErrorsInc()
	GetCartInc()
	GetCartErrorsInc()
}
