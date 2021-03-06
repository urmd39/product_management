// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product HTTP client types
//
// Command:
// $ goa gen goa-demo/design

package client

import (
	product "goa-demo/gen/product"
	productviews "goa-demo/gen/product/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateProductRequestBody is the type of the "product" service
// "createProduct" endpoint HTTP request body.
type CreateProductRequestBody struct {
	// Name of product
	Name          string  `form:"name" json:"name" xml:"name"`
	CurrencyUnit  string  `form:"currency_unit" json:"currency_unit" xml:"currency_unit"`
	Description   string  `form:"description" json:"description" xml:"description"`
	UpdatedTime   string  `form:"updated_time" json:"updated_time" xml:"updated_time"`
	PurchasePrice float64 `form:"purchase_price" json:"purchase_price" xml:"purchase_price"`
	SellingPrice  float64 `form:"selling_price" json:"selling_price" xml:"selling_price"`
}

// UpdateProductRequestBody is the type of the "product" service
// "updateProduct" endpoint HTTP request body.
type UpdateProductRequestBody struct {
	// Product Updated
	Product *ProductRequestBody `form:"product" json:"product" xml:"product"`
}

// GetListProductResponseBody is the type of the "product" service
// "getListProduct" endpoint HTTP response body.
type GetListProductResponseBody []*StoredProductResponse

// GetProductResponseBody is the type of the "product" service "getProduct"
// endpoint HTTP response body.
type GetProductResponseBody struct {
	// ID is the unique id of the bottle.
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of product
	Name          *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	CurrencyUnit  *string  `form:"currency_unit,omitempty" json:"currency_unit,omitempty" xml:"currency_unit,omitempty"`
	Description   *string  `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	UpdatedTime   *string  `form:"updated_time,omitempty" json:"updated_time,omitempty" xml:"updated_time,omitempty"`
	PurchasePrice *float64 `form:"purchase_price,omitempty" json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
	SellingPrice  *float64 `form:"selling_price,omitempty" json:"selling_price,omitempty" xml:"selling_price,omitempty"`
}

// FilterProductResponseBody is the type of the "product" service
// "filterProduct" endpoint HTTP response body.
type FilterProductResponseBody []*StoredProductResponse

// UpdateProductResponseBody is the type of the "product" service
// "updateProduct" endpoint HTTP response body.
type UpdateProductResponseBody struct {
	// ID is the unique id of the bottle.
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of product
	Name          *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	CurrencyUnit  *string  `form:"currency_unit,omitempty" json:"currency_unit,omitempty" xml:"currency_unit,omitempty"`
	Description   *string  `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	UpdatedTime   *string  `form:"updated_time,omitempty" json:"updated_time,omitempty" xml:"updated_time,omitempty"`
	PurchasePrice *float64 `form:"purchase_price,omitempty" json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
	SellingPrice  *float64 `form:"selling_price,omitempty" json:"selling_price,omitempty" xml:"selling_price,omitempty"`
}

// GetProductNotFoundResponseBody is the type of the "product" service
// "getProduct" endpoint HTTP response body for the "not_found" error.
type GetProductNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing product
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// FilterProductNotFoundResponseBody is the type of the "product" service
// "filterProduct" endpoint HTTP response body for the "not_found" error.
type FilterProductNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing product
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredProductResponse is used to define fields on response body types.
type StoredProductResponse struct {
	// ID is the unique id of the bottle.
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of product
	Name          *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	CurrencyUnit  *string  `form:"currency_unit,omitempty" json:"currency_unit,omitempty" xml:"currency_unit,omitempty"`
	Description   *string  `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	UpdatedTime   *string  `form:"updated_time,omitempty" json:"updated_time,omitempty" xml:"updated_time,omitempty"`
	PurchasePrice *float64 `form:"purchase_price,omitempty" json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
	SellingPrice  *float64 `form:"selling_price,omitempty" json:"selling_price,omitempty" xml:"selling_price,omitempty"`
}

// ProductRequestBody is used to define fields on request body types.
type ProductRequestBody struct {
	// Name of product
	Name          string  `form:"name" json:"name" xml:"name"`
	CurrencyUnit  string  `form:"currency_unit" json:"currency_unit" xml:"currency_unit"`
	Description   string  `form:"description" json:"description" xml:"description"`
	UpdatedTime   string  `form:"updated_time" json:"updated_time" xml:"updated_time"`
	PurchasePrice float64 `form:"purchase_price" json:"purchase_price" xml:"purchase_price"`
	SellingPrice  float64 `form:"selling_price" json:"selling_price" xml:"selling_price"`
}

// NewCreateProductRequestBody builds the HTTP request body from the payload of
// the "createProduct" endpoint of the "product" service.
func NewCreateProductRequestBody(p *product.Product) *CreateProductRequestBody {
	body := &CreateProductRequestBody{
		Name:          p.Name,
		CurrencyUnit:  p.CurrencyUnit,
		Description:   p.Description,
		UpdatedTime:   p.UpdatedTime,
		PurchasePrice: p.PurchasePrice,
		SellingPrice:  p.SellingPrice,
	}
	return body
}

// NewUpdateProductRequestBody builds the HTTP request body from the payload of
// the "updateProduct" endpoint of the "product" service.
func NewUpdateProductRequestBody(p *product.UpdateProductPayload) *UpdateProductRequestBody {
	body := &UpdateProductRequestBody{}
	if p.Product != nil {
		body.Product = marshalProductProductToProductRequestBody(p.Product)
	}
	return body
}

// NewGetListProductStoredProductCollectionOK builds a "product" service
// "getListProduct" endpoint result from a HTTP "OK" response.
func NewGetListProductStoredProductCollectionOK(body GetListProductResponseBody) productviews.StoredProductCollectionView {
	v := make([]*productviews.StoredProductView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredProductResponseToProductviewsStoredProductView(val)
	}

	return v
}

// NewGetProductStoredProductOK builds a "product" service "getProduct"
// endpoint result from a HTTP "OK" response.
func NewGetProductStoredProductOK(body *GetProductResponseBody) *productviews.StoredProductView {
	v := &productviews.StoredProductView{
		ID:            body.ID,
		Name:          body.Name,
		CurrencyUnit:  body.CurrencyUnit,
		Description:   body.Description,
		UpdatedTime:   body.UpdatedTime,
		PurchasePrice: body.PurchasePrice,
		SellingPrice:  body.SellingPrice,
	}

	return v
}

// NewGetProductNotFound builds a product service getProduct endpoint not_found
// error.
func NewGetProductNotFound(body *GetProductNotFoundResponseBody) *product.NotFound {
	v := &product.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// NewFilterProductStoredProductCollectionOK builds a "product" service
// "filterProduct" endpoint result from a HTTP "OK" response.
func NewFilterProductStoredProductCollectionOK(body FilterProductResponseBody) productviews.StoredProductCollectionView {
	v := make([]*productviews.StoredProductView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredProductResponseToProductviewsStoredProductView(val)
	}

	return v
}

// NewFilterProductNotFound builds a product service filterProduct endpoint
// not_found error.
func NewFilterProductNotFound(body *FilterProductNotFoundResponseBody) *product.NotFound {
	v := &product.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// NewUpdateProductStoredProductOK builds a "product" service "updateProduct"
// endpoint result from a HTTP "OK" response.
func NewUpdateProductStoredProductOK(body *UpdateProductResponseBody) *productviews.StoredProductView {
	v := &productviews.StoredProductView{
		ID:            body.ID,
		Name:          body.Name,
		CurrencyUnit:  body.CurrencyUnit,
		Description:   body.Description,
		UpdatedTime:   body.UpdatedTime,
		PurchasePrice: body.PurchasePrice,
		SellingPrice:  body.SellingPrice,
	}

	return v
}

// ValidateGetProductNotFoundResponseBody runs the validations defined on
// getProduct_not_found_response_body
func ValidateGetProductNotFoundResponseBody(body *GetProductNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateFilterProductNotFoundResponseBody runs the validations defined on
// filterProduct_not_found_response_body
func ValidateFilterProductNotFoundResponseBody(body *FilterProductNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateStoredProductResponse runs the validations defined on
// StoredProductResponse
func ValidateStoredProductResponse(body *StoredProductResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CurrencyUnit == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("currency_unit", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.UpdatedTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_time", "body"))
	}
	if body.PurchasePrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("purchase_price", "body"))
	}
	if body.SellingPrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("selling_price", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateProductRequestBody runs the validations defined on ProductRequestBody
func ValidateProductRequestBody(body *ProductRequestBody) (err error) {
	if utf8.RuneCountInString(body.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
	}
	return
}
