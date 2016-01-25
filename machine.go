package main

import (
	_ "log"
	"sync"
)

type vmachine struct {
	mu        sync.Mutex
	cstore    map[float64][]*Coin
	pstore    map[int]*Bar
	pstack    *Product
	barid     int
	cstack    []*Coin
	chgstack  []*Coin
	cstacksum float64
	acoins    map[float64]bool
	achange   []float64
	info      *Info
	status    string
}

func NewVMachine() *vmachine {
	v := new(vmachine)

	v.cstore = LoadCoins()
	v.pstore = LoadProducts()
	v.pstack = new(Product)
	v.cstack = make([]*Coin, 0)
	v.chgstack = make([]*Coin, 0)
	v.cstacksum = 0
	v.acoins = AllowedPayment
	v.achange = AllowedChange
	v.info = new(Info)
	v.status = "Idle"
	v.info.DisplayInfo(v)
	return v
}

func (v *vmachine) Insert(c *Coin) {
	v.recognizeCoin(c)
	if c.Recognized {
		v.cstack = append(v.cstack, c)
		v.cstore[c.Nominal] = append(v.cstore[c.Nominal], c)
		v.cstacksum += c.Nominal
	} else {
		v.chgstack = append(v.chgstack, c)
		v.illegalPayment()
		v.returnPayment(c.Recognized)
	}
}

func (v *vmachine) SelectProduct(barid int) {
	if b, ok := v.pstore[barid]; ok {
		v.barid = barid
		if !v.pstore[barid].IsEmpty(v.pstore[barid].Items) {
			v.pstack = b.Items[len(b.Items)-1]
			v.calculator()
		} else {
			v.insufficientProduct()
			v.returnPayment(true)
		}
	}
}

func (v *vmachine) recognizeCoin(c *Coin) {
	if v.acoins[c.Nominal] {
		c.Recognized = true
	}
}

func (v *vmachine) returnChange() {
	v.status = "Sold"
	v.info.DisplayInfo(v)
	for i := range v.cstack {
		_ = i
		v.popCoinStack()
	}
	for i := range v.chgstack {
		_ = i
		v.popChangeStack()
	}
	v.popProductStore()
	v.popProductStack()
	v.cstacksum = 0
	v.status = "Idle"
}

func (v *vmachine) popCoinStack() {
	v.cstack = append(v.cstack[:0], v.cstack[1:]...)
}

func (v *vmachine) popProductStack() {
	v.pstack = &Product{}
}

func (v *vmachine) popProductStore() {
	v.pstore[v.barid].Items = append(v.pstore[v.barid].Items[:0], v.pstore[v.barid].Items[1:]...)
}

func (v *vmachine) popChangeStack() {
	v.chgstack = append(v.chgstack[:0], v.chgstack[1:]...)
}

func (v *vmachine) popCoinStore(d float64) {
	v.cstore[d] = append(v.cstore[d][:0], v.cstore[d][1:]...)
}

func (v *vmachine) returnPayment(r bool) {
	if !r {
		v.info.DisplayInfo(v)
	} else {
		for _, d := range v.cstack {
			i := d.Nominal
			v.chgstack = append(v.chgstack, v.cstore[i][0])
			v.popCoinStore(i)
			v.popCoinStack()
		}
		v.cstacksum = 0
		v.popProductStack()
	}

	for i := range v.chgstack {
		_ = i
		v.popChangeStack()
	}
}

func (v *vmachine) calculator() {
	if v.pstack.PriceUnit <= v.cstacksum {
		change := v.cstacksum - v.pstack.PriceUnit
		for _, d := range v.achange {
			if change == 0 {
				v.returnChange()
				break
			} else if change < d {
				continue
			}
			qty := int(change / d)
			if len(v.cstore[d])-qty < 0 {
				if d == v.achange[len(v.achange)-1] {
					v.chgstack = v.cstack
					v.insufficientChange()
					v.returnPayment(true)
					break
				}
				continue
			} else {
				for j := 0; j < qty; j++ {
					v.chgstack = append(v.chgstack, v.cstore[d][0])
					v.popCoinStore(d)
				}
			}
			change -= float64(qty) * d
			if change == 0 {
				v.returnChange()
				break
			}
			if change != 0 && d == v.achange[len(v.achange)-1] {
				v.insufficientChange()
				v.returnPayment(true)
				break
			}
		}
	} else {
		v.insufficientCredit()
	}
}

func (v *vmachine) insufficientCredit() {
	v.status = "Insufficient credit"
	v.info.DisplayInfo(v)
}

func (v *vmachine) insufficientChange() {
	v.status = "Sorry, you can not buy, not enough change"
	v.info.DisplayInfo(v)
}

func (v *vmachine) insufficientProduct() {
	v.status = "Sorry, you can not buy this, not enough product. Change other."
	v.info.DisplayInfo(v)
}

func (v *vmachine) illegalPayment() {
	v.status = "Sorry, your payment is illegal. Try to use another coin or note."
	v.info.DisplayInfo(v)
}
