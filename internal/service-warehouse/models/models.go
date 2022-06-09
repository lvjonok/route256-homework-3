package models

import types "gitlab.ozon.dev/lvjonok/homework-3/core/models"

type Entry struct {
	ID        types.ID
	ProductID types.ID // references product id from marketplace
	Quantity  int
}
