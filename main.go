package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urmd39/product_management/entities"
	"github.com/urmd39/product_management/model"
)

func main() {

	/// Create routers
	router := mux.NewRouter()

	// router.HandleFunc("/api/product", GetProduct).Methods("GET")
	router.HandleFunc("/api/product", CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", GetProductById).Methods("GET")
	router.HandleFunc("/api/product/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product/{id}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/product", FilterProductWithCU).
		Queries("currency_unit", "{currency_unit}").Methods("GET")
	router.HandleFunc("/api/product", GetProduct).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Server start at port 8000")
	http.ListenAndServe(":8000", nil)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listProduct := model.GetListProduct()
	json.NewEncoder(w).Encode(listProduct)
	fmt.Println("Get All")
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	fmt.Println(id)
	pd := model.GetProductById(id)
	json.NewEncoder(w).Encode(pd)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var pd entities.Product
	json.Unmarshal(body, &pd)

	model.AddProduct(pd)
	json.NewEncoder(w).Encode(pd)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	model.Delete_Product(id)
	fmt.Fprintf(w, "Product with id = %d has been deleted!!!/n", id)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	body, _ := ioutil.ReadAll(r.Body)
	var pd entities.Product
	json.Unmarshal(body, &pd)
	model.Update_Product(id, pd)
	json.NewEncoder(w).Encode(pd)
}

func FilterProductWithCU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filter := r.FormValue("currency_unit")
	listProduct := model.FilterProductWithCU(strings.ToUpper(filter))
	// fmt.Printf("%T", filter)
	json.NewEncoder(w).Encode(listProduct)
}
