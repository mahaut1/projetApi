package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux" // implémentation du routeur
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// ... d'autres champs selon vos besoins
}

var items []Item

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items", addItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	items = append(items, newItem)
	json.NewEncoder(w).Encode(items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedItem Item
	_ = json.NewDecoder(r.Body).Decode(&updatedItem)

	for index, item := range items {
		if item.ID == id {
			// Update the item found in the slice
			items[index] = updatedItem
			json.NewEncoder(w).Encode(items[index])
			return
		}
	}

	// If the ID is not found, return an error or appropriate response
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Item not found"))
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, item := range items {
		if item.ID == id {
			// Remove the item from the slice
			items = append(items[:index], items[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Item deleted"))
			return
		}
	}

	// If the ID is not found, return an error or appropriate response
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Item not found"))
}
