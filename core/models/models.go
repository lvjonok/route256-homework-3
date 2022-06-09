package models

type ID uint64

func Int2ID(i uint64) *ID {
	r := ID(i)
	return &r
}

type ProductUnit struct {
	ProductID ID
	Quantity  int
}
