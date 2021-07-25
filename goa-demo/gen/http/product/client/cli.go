// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product HTTP client CLI support package
//
// Command:
// $ goa gen goa-demo/design

package client

import (
	"encoding/json"
	"fmt"
	product "goa-demo/gen/product"
	"strconv"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetProductPayload builds the payload for the product getProduct
// endpoint from CLI flags.
func BuildGetProductPayload(productGetProductID string, productGetProductView string) (*product.GetProductPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(productGetProductID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	var view *string
	{
		if productGetProductView != "" {
			view = &productGetProductView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &product.GetProductPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildFilterProductPayload builds the payload for the product filterProduct
// endpoint from CLI flags.
func BuildFilterProductPayload(productFilterProductCurrencyUnit string) (*product.FilterProductPayload, error) {
	var currencyUnit string
	{
		currencyUnit = productFilterProductCurrencyUnit
	}
	v := &product.FilterProductPayload{}
	v.CurrencyUnit = currencyUnit

	return v, nil
}

// BuildCreateProductPayload builds the payload for the product createProduct
// endpoint from CLI flags.
func BuildCreateProductPayload(productCreateProductBody string) (*product.Product, error) {
	var err error
	var body CreateProductRequestBody
	{
		err = json.Unmarshal([]byte(productCreateProductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"currency_unit\": \"Minima debitis et et sed rerum.\",\n      \"description\": \"Voluptas rerum tempore reiciendis.\",\n      \"name\": \"Blue\\'s Cuvee\",\n      \"purchase_price\": 0.869811431951743,\n      \"selling_price\": 0.08936100869612645,\n      \"updated_time\": \"Quis et.\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &product.Product{
		Name:          body.Name,
		CurrencyUnit:  body.CurrencyUnit,
		Description:   body.Description,
		UpdatedTime:   body.UpdatedTime,
		PurchasePrice: body.PurchasePrice,
		SellingPrice:  body.SellingPrice,
	}

	return v, nil
}

// BuildUpdateProductPayload builds the payload for the product updateProduct
// endpoint from CLI flags.
func BuildUpdateProductPayload(productUpdateProductBody string, productUpdateProductID string) (*product.UpdateProductPayload, error) {
	var err error
	var body UpdateProductRequestBody
	{
		err = json.Unmarshal([]byte(productUpdateProductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"product\": {\n         \"currency_unit\": \"Possimus earum.\",\n         \"description\": \"Ut repudiandae.\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"purchase_price\": 0.9193737983768083,\n         \"selling_price\": 0.8555095810048335,\n         \"updated_time\": \"In non nihil provident qui vel iste.\"\n      }\n   }'")
		}
		if body.Product == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("product", "body"))
		}
		if body.Product != nil {
			if err2 := ValidateProductRequestBody(body.Product); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(productUpdateProductID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &product.UpdateProductPayload{}
	if body.Product != nil {
		v.Product = marshalProductRequestBodyToProductProduct(body.Product)
	}
	v.ID = id

	return v, nil
}

// BuildRemoveProductPayload builds the payload for the product removeProduct
// endpoint from CLI flags.
func BuildRemoveProductPayload(productRemoveProductID string) (*product.RemoveProductPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(productRemoveProductID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &product.RemoveProductPayload{}
	v.ID = id

	return v, nil
}