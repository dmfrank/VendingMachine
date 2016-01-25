package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payment struct {
	BarId   int     `json:"pid"`
	Payment float64 `json:"payment"`
}

func handler(vm *vmachine) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			vm.mu.Lock()
			defer vm.mu.Unlock()
			vm.info.DisplayInfo(vm)
			data, err := json.Marshal(vm.info)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		} else if req.Method == "POST" {
			decoder := json.NewDecoder(req.Body)
			var p Payment

			err := decoder.Decode(&p)
			if err != nil {
				log.Println(err)
			}
			var Coin = new(Coin)
			Coin.Nominal = p.Payment
			vm.mu.Lock()
			defer vm.mu.Unlock()

			vm.Insert(Coin)
			vm.info.DisplayInfo(vm)

			if len(vm.chgstack) == 0 || p.BarId > 0 {
				vm.SelectProduct(p.BarId)
			}

			data, err := json.Marshal(vm.info)

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}
}

func main() {
	vm := NewVMachine()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler(vm))
	log.Print("Start vending machine Http server...")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
