package models

import types "gitlab.ozon.dev/lvjonok/homework-3/core/models"

type OrderStatus string

var (
	// we initialized order in database and ready to process
	// we do not compensate this transaction
	Created OrderStatus = "created"

	// checked that there were enough items in warehouse
	// we do not compensate this transaction
	Checked OrderStatus = "checked"

	// successfully booked required amount of items in the warehouse
	//
	Booked OrderStatus = "booked"
)

type Order struct {
	OrderID  types.ID
	UserID   types.ID
	Products []types.ProductUnit
	Status   string

	SagaStatus OrderStatus
}
