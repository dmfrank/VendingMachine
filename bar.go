package main

type Bar struct {
	Id        int
	PriceUnit float64
	IsEmpty   func([]*Product) bool
	Items     []*Product
}

var IsEmptyBar = func(s []*Product) bool {
	if len(s) > 0 {
		return false
	}
	return true
}
