package exampleapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product represents a product in the store
type Product struct {
	Title          string  `json:"title,omitempty"`
	Price          float64 `json:"price,omitempty"`
	InventoryCount int     `json:"inventory_count,omitempty"`
}

var store []Product

func main() {
	test1 := Product{Title: "test1", Price: 5.99, InventoryCount: 1}
	test2 := Product{Title: "test2", Price: 10.99, InventoryCount: 0}
	store = append(store, test1)
	store = append(store, test2)

	router := mux.NewRouter()
	router.HandleFunc("/products/{show_empty}", GetProducts).Methods("GET")
	router.HandleFunc("/product/{title}", GetProduct).Methods("GET")
	router.HandleFunc("/product/purchase/{title}", PurchaseProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetProducts returns all products in the store
func GetProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	showEmpty := params["show_empty"]
	_, ok := params["test"]
	if !ok {
		fmt.Println("test")
	}
	if showEmpty == "true" {
		//w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(store)
	} else {
		var nonEmptyProducts []Product
		for _, v := range store {
			if v.InventoryCount > 0 {
				nonEmptyProducts = append(nonEmptyProducts, v)
			}
		}
		json.NewEncoder(w).Encode(nonEmptyProducts)
	}
}

// GetProduct returns the product with the specified title
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]
	for _, v := range store {
		if v.Title == title {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("Not Found")
}

// GetTest returns nothing
func GetTest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Test")
}

// PurchaseProduct decrements the product with the specified title by 1
func PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]
	for i, v := range store {
		if v.Title == title {
			if v.InventoryCount <= 0 {
				json.NewEncoder(w).Encode("Not enough inventory")
				return
			}
			store[i].InventoryCount--
			json.NewEncoder(w).Encode(v)
		}
	}
}
