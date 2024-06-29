package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Order struct {
	ID     string 
	UserID string 
	Item   string 
}

var orders = []Order{
	{ID: "1", UserID: "1", Item: "Book"},
	{ID: "2", UserID: "2", Item: "Pen"},
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(orders)
  }

func getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, order := range orders {
		if order.ID == id {
			w.Header().Set("Content-Type", "application/json")
 			json.NewEncoder(w).Encode(order)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", getOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", getOrder).Methods("GET")

	log.Println("Order Service is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
