// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product gRPC client CLI support package
//
// Command:
// $ goa gen goa-demo/design

package client

import (
	"encoding/json"
	"fmt"
	productpb "goa-demo/gen/grpc/product/pb"
	product "goa-demo/gen/product"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetProductPayload builds the payload for the product getProduct
// endpoint from CLI flags.
func BuildGetProductPayload(productGetProductMessage string, productGetProductView string) (*product.GetProductPayload, error) {
	var err error
	var message productpb.GetProductRequest
	{
		if productGetProductMessage != "" {
			err = json.Unmarshal([]byte(productGetProductMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 287285371201450806\n   }'")
			}
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
	v := &product.GetProductPayload{
		ID: int(message.Id),
	}
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
func BuildCreateProductPayload(productCreateProductMessage string) (*product.Product, error) {
	var err error
	var message productpb.CreateProductRequest
	{
		if productCreateProductMessage != "" {
			err = json.Unmarshal([]byte(productCreateProductMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"currency_unit\": \"Modi aliquid non.\",\n      \"description\": \"Eos voluptate in.\",\n      \"name\": \"Blue\\'s Cuvee\",\n      \"purchase_price\": 0.2030375670316061,\n      \"selling_price\": 0.24227437102930335,\n      \"updated_time\": \"Dolorem iure qui quasi.\"\n   }'")
			}
		}
	}
	v := &product.Product{
		Name:          message.Name,
		CurrencyUnit:  message.CurrencyUnit,
		Description:   message.Description,
		UpdatedTime:   message.UpdatedTime,
		PurchasePrice: message.PurchasePrice,
		SellingPrice:  message.SellingPrice,
	}

	return v, nil
}

// BuildUpdateProductPayload builds the payload for the product updateProduct
// endpoint from CLI flags.
func BuildUpdateProductPayload(productUpdateProductMessage string) (*product.UpdateProductPayload, error) {
	var err error
	var message productpb.UpdateProductRequest
	{
		if productUpdateProductMessage != "" {
			err = json.Unmarshal([]byte(productUpdateProductMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 2439857911910992395,\n      \"product\": {\n         \"currency_unit\": \"Facilis quisquam suscipit magnam magnam quidem.\",\n         \"description\": \"Aut sed quae dolores.\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"purchase_price\": 0.6121581005022516,\n         \"selling_price\": 0.00734776034275089,\n         \"updated_time\": \"Harum incidunt minus.\"\n      }\n   }'")
			}
		}
	}
	v := &product.UpdateProductPayload{
		ID: int(message.Id),
	}
	if message.Product != nil {
		v.Product = protobufProductpbProduct2ToProductProduct(message.Product)
	}

	return v, nil
}

// BuildRemoveProductPayload builds the payload for the product removeProduct
// endpoint from CLI flags.
func BuildRemoveProductPayload(productRemoveProductMessage string) (*product.RemoveProductPayload, error) {
	var err error
	var message productpb.RemoveProductRequest
	{
		if productRemoveProductMessage != "" {
			err = json.Unmarshal([]byte(productRemoveProductMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 8160317928609962447\n   }'")
			}
		}
	}
	v := &product.RemoveProductPayload{
		ID: int(message.Id),
	}

	return v, nil
}
