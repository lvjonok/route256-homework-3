package service

type Metrics interface {
	RequestsInc()
	RequestsErrorsInc()
	RegisterProductInc()
	RegisterProductErrorsInc()
	CheckProductsInc()
	CheckProductsErrorsInc()
	BookProductsInc()
	BookProductsErrorsInc()
	UnbookProductsInc()
	UnbookProductsErrorsInc()
}
