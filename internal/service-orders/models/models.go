package models

import types "gitlab.ozon.dev/lvjonok/homework-3/core/models"

type Order struct {
	OrderID  types.ID
	UserID   types.ID
	Products []types.ProductUnit
	Status   string
}
