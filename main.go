package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ohadvaknin/go-wolt-delivery/models"
	"github.com/ohadvaknin/go-wolt-delivery/storage"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var order models.Delivery
	err := json.NewDecoder(r.Body).Decode(&order)
	if (err != nil) {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	order = storage.CreateOrder(order)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(order)
}	

func main() {
	N := 2
	storage.InitializeServer(N)
	http.HandleFunc("/delivery/create", HandleRequest)
	log.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
