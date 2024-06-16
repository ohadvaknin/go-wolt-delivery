package main

import (
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	N := 20
	
	http.HandleFunc("/delivery/create", HandleRequest)
	log.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
