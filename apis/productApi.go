package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/urmd39/product_management/control"
	"github.com/urmd39/product_management/entities"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listProduct := control.GetListProduct()
	json.NewEncoder(w).Encode(listProduct)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	id := ctx.Value("product").(int)

	pd := control.GetProductById(id)
	json.NewEncoder(w).Encode(pd)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var pd entities.Product
	json.Unmarshal(body, &pd)

	control.AddProduct(pd)
	json.NewEncoder(w).Encode(pd)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	id := ctx.Value("product").(int)

	control.Delete_Product(id)
	fmt.Fprintf(w, "Product with id = %d has been deleted!!!", id)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	id := ctx.Value("product").(int)

	body, _ := ioutil.ReadAll(r.Body)
	var pd entities.Product
	json.Unmarshal(body, &pd)
	control.Update_Product(id, pd)
	json.NewEncoder(w).Encode(pd)
}

func FilterProductWithCU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	val := r.FormValue("currency_unit")
	listProduct := control.FilterProductWithCU(strings.ToUpper(val))
	json.NewEncoder(w).Encode(listProduct)
}

func ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		// product := GetProductById(id)

		ctx := context.WithValue(r.Context(), "product", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
