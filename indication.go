package main

type Info struct {
	System struct {
		Payment struct {
			Stack []float64 `json:"coins"`
			Sum   float64   `json:"amount"`
		} `json:"payment"`
		Status string `json:"status"`
	} `json:"system"`
	User struct {
		Product struct {
			Name      string  `json:"name"`
			PriceUnit float64 `json:"price"`
		} `json:"product"`
		Change struct {
			Stack []float64 `json:"coins"`
			Sum   float64   `json:"amount"`
		} `json:"change"`
	} `json:"user_benefit"`
}

func (i *Info) DisplayInfo(v *vmachine) {
	i.System.Payment.Stack, i.System.Payment.Sum = func() ([]float64, float64) {
		stack := make([]float64, 0)
		for _, coin := range v.cstack {
			stack = append(stack, coin.Nominal)
		}
		return stack, v.cstacksum
	}()
	i.User.Product.Name = v.pstack.Name
	i.User.Product.PriceUnit = v.pstack.PriceUnit
	i.User.Change.Stack, i.User.Change.Sum = func() ([]float64, float64) {
		coins := make([]float64, 0)
		var sum float64
		for _, el := range v.chgstack {
			sum += el.Nominal
			coins = append(coins, el.Nominal)
		}
		return coins, sum
	}()
	i.System.Status = v.status
}
