package app

import (
	"baking/service"
	"encoding/json"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers)getAllCustomers(w http.ResponseWriter, r *http.Request)  {

	customer, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(customer); err != nil {
		panic(err)
	}
}