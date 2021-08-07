package app

import (
	"baking/domain"
	"baking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9999", router))
}