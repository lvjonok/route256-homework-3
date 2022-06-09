package models

import (
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
)

type Product struct {
	ID   types.ID
	Name string
	Desc string
}

type Review struct {
	ID        types.ID
	ProductID types.ID // foreign key of Product
	Text      string
}

type Cart struct {
	UserID   types.ID // id of user
	Products []types.ProductUnit
}
