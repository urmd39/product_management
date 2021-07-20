package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urmd39/product_management/entities"
	"github.com/urmd39/product_management/model"
)

func main() {

	/// Create routers
	router := mux.NewRouter()

	router.HandleFunc("/api/product", getProduct).Methods("GET")
	router.HandleFunc("/api/product", createProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", getProductById).Methods("GET")
	router.HandleFunc("/api/product/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/api/product/{id}", deleteProduct).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Server start at port 8000")
	http.ListenAndServe(":8000", nil)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listProduct := model.GetListProduct()
	json.NewEncoder(w).Encode(listProduct)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	fmt.Println(id)
	pd := model.GetProductById(id)
	json.NewEncoder(w).Encode(pd)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var pd entities.Product
	json.Unmarshal(body, &pd)

	model.AddProduct(pd)
	json.NewEncoder(w).Encode(pd)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	model.Delete_Product(id)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	body, _ := ioutil.ReadAll(r.Body)
	var pd entities.Product
	json.Unmarshal(body, &pd)
	model.Update_Product(id, pd)
	json.NewEncoder(w).Encode(pd)
}
