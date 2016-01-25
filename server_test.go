package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type UserAnswer struct {
	System struct {
		Payment struct {
			Coins  []float64 `json:"coins"`
			Amount float64   `json:"amount"`
		} `json:"payment"`
		Status string `json:"status"`
	} `json:"system"`
	User struct {
		Product struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"product"`
		Change struct {
			Coins  []float64 `json:"coins"`
			Amount float64   `json:"amount`
		} `json:"change"`
	} `json:"user_benefit"`
}

func doGetInformation(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:5000/", nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		t.Error(fmt.Printf("Error when set connection. Error: %s\n", err.Error()))
	}

	if resp.StatusCode != 200 {
		t.Error(fmt.Printf("Error when sending message to user. StatusCode: %d\n", resp.StatusCode))
	}
}

func doIllegalPayment(t *testing.T) {
	params, _ := json.Marshal(&Payment{
		Payment: 100,
	})
	req, err := http.NewRequest("POST", "http://localhost:5000/", bytes.NewBuffer(params))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		t.Error(fmt.Printf("Error when set connection. Error: %s\n", err.Error()))
	}

	if resp.StatusCode == 200 {
		a := new(UserAnswer)
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, a)
		if err != nil {
			log.Println(err)
		}
		if a.System.Payment.Amount != 0 && a.User.Change.Amount == 0 {
			t.Error(fmt.Printf("Illegal payment error."))
		}
	} else {
		t.Error(fmt.Printf("Payment error. StatusCode: %d\n", resp.StatusCode))
	}
}

func doBuyProduct(t *testing.T) {
	for _, coin := range AllowedChange {
		params, _ := json.Marshal(&Payment{
			Payment: coin,
			BarId:   10,
		})

		req, err := http.NewRequest("POST", "http://localhost:5000/", bytes.NewBuffer(params))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		defer resp.Body.Close()
		if err != nil {
			t.Error(fmt.Printf("Error when set connection. Error: %d\n", err.Error()))
			break
		}
		m := new(UserAnswer)
		if resp.StatusCode == 200 {
			data, _ := ioutil.ReadAll(resp.Body)
			err := json.Unmarshal(data, m)
			if err != nil {
				log.Println(err)
				break
			}
			/*fmt.Printf("\n------------- action begin ----------\n")
			fmt.Printf("You paid with Coins - %v, Amounts - %v\n", m.System.Payment.Coins, m.System.Payment.Amount)
			fmt.Printf("------------- Status - %v\n", m.System.Status)
			fmt.Printf("Your choose product Name - %v, Price - %v\n", m.User.Product.Name, m.User.Product.Price)
			fmt.Printf("-------------")
			fmt.Printf("Your change with Coins- %v, Amount - %v\n", m.User.Change.Coins, m.User.Change.Amount)
			fmt.Printf("------------- end ----------\n")*/

		} else {
			t.Error(fmt.Printf("Payment error. StatusCode %d\n", resp.StatusCode))
			break
		}
	}

}

func TestServer(t *testing.T) {
	go func() {
		vm := NewVMachine()
		mux := http.NewServeMux()
		mux.HandleFunc("/", handler(vm))
		log.Fatal(http.ListenAndServe(":5000", mux))
	}()
	// #0 Case
	doGetInformation(t)

	// #1 Case
	doIllegalPayment(t)

	// #2 Case
	doBuyProduct(t)

	// #3 Case
	//doOrderWithoutPayment(t)
}
