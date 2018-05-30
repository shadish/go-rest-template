package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Customer struct {
	ID             string   `json:"id,omitempty"`
	FirstName      string   `json:"firstName,omitempty"`
	LastName       string   `json:"lastName,omitempty"`
	ShipAddress    *Address `json:"shipAddress,omitempty"`
	ServiceAddress *Address `json:"serviceAddress,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var customers []Customer

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range customers {
		if item.ID == params["id"] {
			log.Printf("yup, ", item)
			json.NewEncoder(w).Encode(item)
			return
		}
		log.Printf("no customer found")
		json.NewEncoder(w).Encode(&Customer{})
	}
}

func main() {
	router := mux.NewRouter()

	customers = append(customers, Customer{ID: "33", FirstName: "John", LastName: "Doe", ShipAddress: &Address{City: "City X", State: "State X"}, ServiceAddress: &Address{City: "City Y", State: "State Y"}})

	router.HandleFunc("/customer/{id}", GetCustomer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
