package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urmd39/product_management/apis"
)

// var db *sql.DB
// var err error

func main() {
	// infoDB := "host=localhost user=postgres dbname=product_management password=model.kn2412" +
	// 	" sslmode=disable"
	// db, err = sql.Open("postgres", infoDB)
	// if err != nil {
	// 	panic(err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// fmt.Println(ctr.GetProductById(1))

	/// Create routers
	router := mux.NewRouter()

	// http.HandleFunc("/api/product", apis.GetProduct)
	router.HandleFunc("/api/product", apis.GetProduct).Methods("GET")
	// router.HandleFunc("/api/product", createPost).Methods("POST")
	// router.HandleFunc("/api/product/{id}", getPost).Methods("GET")
	// router.HandleFunc("/api/product/{id}", updatePost).Methods("PUT")
	// router.HandleFunc("/api/product/{id}", deletePost).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Server start at port 8000")
	http.ListenAndServe(":8000", nil)
}
