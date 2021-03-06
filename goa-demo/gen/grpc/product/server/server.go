// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product gRPC server
//
// Command:
// $ goa gen goa-demo/design

package server

import (
	"context"
	productpb "goa-demo/gen/grpc/product/pb"
	product "goa-demo/gen/product"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the productpb.ProductServer interface.
type Server struct {
	GetListProductH goagrpc.UnaryHandler
	GetProductH     goagrpc.UnaryHandler
	FilterProductH  goagrpc.UnaryHandler
	CreateProductH  goagrpc.UnaryHandler
	UpdateProductH  goagrpc.UnaryHandler
	RemoveProductH  goagrpc.UnaryHandler
	productpb.UnimplementedProductServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the product service endpoints.
func New(e *product.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetListProductH: NewGetListProductHandler(e.GetListProduct, uh),
		GetProductH:     NewGetProductHandler(e.GetProduct, uh),
		FilterProductH:  NewFilterProductHandler(e.FilterProduct, uh),
		CreateProductH:  NewCreateProductHandler(e.CreateProduct, uh),
		UpdateProductH:  NewUpdateProductHandler(e.UpdateProduct, uh),
		RemoveProductH:  NewRemoveProductHandler(e.RemoveProduct, uh),
	}
}

// NewGetListProductHandler creates a gRPC handler which serves the "product"
// service "getListProduct" endpoint.
func NewGetListProductHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeGetListProductResponse)
	}
	return h
}

// GetListProduct implements the "GetListProduct" method in
// productpb.ProductServer interface.
func (s *Server) GetListProduct(ctx context.Context, message *productpb.GetListProductRequest) (*productpb.StoredProductCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getListProduct")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.GetListProductH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.StoredProductCollection), nil
}

// NewGetProductHandler creates a gRPC handler which serves the "product"
// service "getProduct" endpoint.
func NewGetProductHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetProductRequest, EncodeGetProductResponse)
	}
	return h
}

// GetProduct implements the "GetProduct" method in productpb.ProductServer
// interface.
func (s *Server) GetProduct(ctx context.Context, message *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getProduct")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.GetProductH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				er := err.(*product.NotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewGetProductNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.GetProductResponse), nil
}

// NewFilterProductHandler creates a gRPC handler which serves the "product"
// service "filterProduct" endpoint.
func NewFilterProductHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeFilterProductRequest, EncodeFilterProductResponse)
	}
	return h
}

// FilterProduct implements the "FilterProduct" method in
// productpb.ProductServer interface.
func (s *Server) FilterProduct(ctx context.Context, message *productpb.FilterProductRequest) (*productpb.StoredProductCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "filterProduct")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.FilterProductH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				er := err.(*product.NotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewFilterProductNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.StoredProductCollection), nil
}

// NewCreateProductHandler creates a gRPC handler which serves the "product"
// service "createProduct" endpoint.
func NewCreateProductHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateProductRequest, EncodeCreateProductResponse)
	}
	return h
}

// CreateProduct implements the "CreateProduct" method in
// productpb.ProductServer interface.
func (s *Server) CreateProduct(ctx context.Context, message *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "createProduct")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.CreateProductH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.CreateProductResponse), nil
}

// NewUpdateProductHandler creates a gRPC handler which serves the "product"
// service "updateProduct" endpoint.
func NewUpdateProductHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateProductRequest, EncodeUpdateProductResponse)
	}
	return h
}

// UpdateProduct implements the "UpdateProduct" method in
// productpb.ProductServer interface.
func (s *Server) UpdateProduct(ctx context.Context, message *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "updateProduct")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.UpdateProductH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.UpdateProductResponse), nil
}

// NewRemoveProductHandler creates a gRPC handler which serves the "product"
// service "removeProduct" endpoint.
func NewRemoveProductHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveProductRequest, EncodeRemoveProductResponse)
	}
	return h
}

// RemoveProduct implements the "RemoveProduct" method in
// productpb.ProductServer interface.
func (s *Server) RemoveProduct(ctx context.Context, message *productpb.RemoveProductRequest) (*productpb.RemoveProductResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "removeProduct")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.RemoveProductH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.RemoveProductResponse), nil
}
