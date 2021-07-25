package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/urmd39/product_management/apis"
)

func main() {
	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Homepage"))
	})

	r.Route("/product", func(r chi.Router) {
		r.With(paginate).Get("/", apis.GetProducts)
		r.Post("/", apis.CreateProduct)
		r.Get("/filter", apis.FilterProductWithCU)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(apis.ProductCtx)
			r.Get("/", apis.GetProductById)
			r.Put("/", apis.UpdateProduct)
			r.Delete("/", apis.DeleteProduct)
		})
	})

	// Mount the admin sub-router
	// r.Mount("/admin", adminRouter())

	http.ListenAndServe(":3000", r)
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
