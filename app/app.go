package app

import (
	"baking/domain"
	"baking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9999", router))
}
