// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product endpoints
//
// Command:
// $ goa gen goa-demo/design

package product

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "product" service endpoints.
type Endpoints struct {
	GetListProduct goa.Endpoint
	GetProduct     goa.Endpoint
	FilterProduct  goa.Endpoint
	CreateProduct  goa.Endpoint
	UpdateProduct  goa.Endpoint
	RemoveProduct  goa.Endpoint
}

// NewEndpoints wraps the methods of the "product" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetListProduct: NewGetListProductEndpoint(s),
		GetProduct:     NewGetProductEndpoint(s),
		FilterProduct:  NewFilterProductEndpoint(s),
		CreateProduct:  NewCreateProductEndpoint(s),
		UpdateProduct:  NewUpdateProductEndpoint(s),
		RemoveProduct:  NewRemoveProductEndpoint(s),
	}
}

// Use applies the given middleware to all the "product" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetListProduct = m(e.GetListProduct)
	e.GetProduct = m(e.GetProduct)
	e.FilterProduct = m(e.FilterProduct)
	e.CreateProduct = m(e.CreateProduct)
	e.UpdateProduct = m(e.UpdateProduct)
	e.RemoveProduct = m(e.RemoveProduct)
}

// NewGetListProductEndpoint returns an endpoint function that calls the method
// "getListProduct" of service "product".
func NewGetListProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, view, err := s.GetListProduct(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredProductCollection(res, view)
		return vres, nil
	}
}

// NewGetProductEndpoint returns an endpoint function that calls the method
// "getProduct" of service "product".
func NewGetProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetProductPayload)
		res, view, err := s.GetProduct(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredProduct(res, view)
		return vres, nil
	}
}

// NewFilterProductEndpoint returns an endpoint function that calls the method
// "filterProduct" of service "product".
func NewFilterProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FilterProductPayload)
		res, view, err := s.FilterProduct(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredProductCollection(res, view)
		return vres, nil
	}
}

// NewCreateProductEndpoint returns an endpoint function that calls the method
// "createProduct" of service "product".
func NewCreateProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Product)
		return s.CreateProduct(ctx, p)
	}
}

// NewUpdateProductEndpoint returns an endpoint function that calls the method
// "updateProduct" of service "product".
func NewUpdateProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateProductPayload)
		res, view, err := s.UpdateProduct(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredProduct(res, view)
		return vres, nil
	}
}

// NewRemoveProductEndpoint returns an endpoint function that calls the method
// "removeProduct" of service "product".
func NewRemoveProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemoveProductPayload)
		return nil, s.RemoveProduct(ctx, p)
	}
}