package apis

import (
	"fmt"
	"net/http"
)

// type Post struct {
// 	ID    string `json:"id"`
// 	Title string `json:"title"`
// }

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// listProduct := model.GetListProduct()
	// json.NewEncoder(w).Encode(listProduct)
	fmt.Fprintln(w, "Hello")
}
