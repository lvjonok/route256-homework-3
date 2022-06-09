package service

type Metrics interface {
	RequestsInc()
	RequestErrorsInc()
	CreateOrderInc()
	CreateOrderErrorsInc()
	CheckStatusInc()
	CheckStatusErrorsInc()
	UpdateStatusInc()
	UpdateStatusErrorsInc()
}
