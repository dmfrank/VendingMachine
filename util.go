package main

var AllowedPayment map[float64]bool = map[float64]bool{
	0.5: true,
	1:   true,
	2:   true,
	5:   true,
	10:  true,
	20:  true,
	50:  false,
	100: false,
}

var AllowedChange []float64 = []float64{5, 2, 1, 0.5}

var products map[int][]*Product = map[int][]*Product{
	10: []*Product{
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
		{Name: "orange", PriceUnit: 1.5},
	},
	11: []*Product{
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
		{Name: "apple", PriceUnit: 1},
	},
}

var coins map[float64][]*Coin = map[float64][]*Coin{
	0.5: []*Coin{
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
		{Nominal: 0.5, Recognized: true},
	},
	1: []*Coin{
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
		{Nominal: 1, Recognized: true},
	},
	2: []*Coin{
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
		{Nominal: 2, Recognized: true},
	},
	5: []*Coin{
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
		{Nominal: 5, Recognized: true},
	},
}

func LoadProducts() map[int]*Bar {
	m := make(map[int]*Bar)
	for k, v := range products {
		b := &Bar{
			Id:        k,
			PriceUnit: v[0].PriceUnit,
			IsEmpty:   IsEmptyBar,
			Items:     v,
		}
		m[k] = b
	}
	return m
}
func LoadCoins() map[float64][]*Coin {
	return coins
}
