package productapi

import (
	"context"
	"fmt"
	"goa-demo/control"
	product "goa-demo/gen/product"
	"log"
	"strings"
)

// product service example implementation.
// The example methods log the requests and return zero values.
type productsrvc struct {
	logger *log.Logger
}

// NewProduct returns the product service implementation.
func NewProduct(logger *log.Logger) product.Service {
	return &productsrvc{logger}
}

// GetListProduct implements getListProduct.
func (s *productsrvc) GetListProduct(ctx context.Context) (res product.StoredProductCollection, view string, err error) {
	view = "default"
	s.logger.Print("product.getListProduct")
	res = control.GetListProduct()
	return res, view, err
}

// GetProduct implements getProduct.
func (s *productsrvc) GetProduct(ctx context.Context, p *product.GetProductPayload) (res *product.StoredProduct, view string, err error) {
	// res = &product.StoredProduct{}
	s.logger.Print("product.getProduct")
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}
	res = control.GetProductById(p.ID)
	return res, view, err
}

// FilterProduct implements filterProduct.
func (s *productsrvc) FilterProduct(ctx context.Context, p *product.FilterProductPayload) (res product.StoredProductCollection, view string, err error) {
	view = "default"
	s.logger.Print("product.filterProduct")
	cu := strings.ToUpper(p.CurrencyUnit)
	fmt.Println(cu)
	res = control.FilterProductWithCU(cu)
	return res, view, err
}

// CreateProduct implements createProduct.
func (s *productsrvc) CreateProduct(ctx context.Context, p *product.Product) (res string, err error) {
	s.logger.Print("product.createProduct")
	res = control.AddProduct(p)
	return res, err
}

// UpdateProduct implements updateProduct.
func (s *productsrvc) UpdateProduct(ctx context.Context, p *product.UpdateProductPayload) (res *product.StoredProduct, view string, err error) {
	view = "default"
	s.logger.Print("product.updateProduct")
	res = control.Update_Product(p.ID, *p.Product)
	return res, view, err
}

// Remove product from storage
func (s *productsrvc) RemoveProduct(ctx context.Context, p *product.RemoveProductPayload) (err error) {
	s.logger.Print("product.removeProduct")
	control.Delete_Product(p.ID)
	return
}