// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product views
//
// Command:
// $ goa gen goa-demo/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredProductCollection is the viewed result type that is projected based on
// a view.
type StoredProductCollection struct {
	// Type to project
	Projected StoredProductCollectionView
	// View to render
	View string
}

// StoredProduct is the viewed result type that is projected based on a view.
type StoredProduct struct {
	// Type to project
	Projected *StoredProductView
	// View to render
	View string
}

// StoredProductCollectionView is a type that runs validations on a projected
// type.
type StoredProductCollectionView []*StoredProductView

// StoredProductView is a type that runs validations on a projected type.
type StoredProductView struct {
	// ID is the unique id of the bottle.
	ID *int
	// Name of product
	Name          *string
	CurrencyUnit  *string
	Description   *string
	UpdatedTime   *string
	PurchasePrice *float64
	SellingPrice  *float64
}

var (
	// StoredProductCollectionMap is a map of attribute names in result type
	// StoredProductCollection indexed by view name.
	StoredProductCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"currency_unit",
			"description",
			"updated_time",
			"purchase_price",
			"selling_price",
		},
		"tiny": []string{
			"id",
			"name",
			"currency_unit",
			"selling_price",
		},
	}
	// StoredProductMap is a map of attribute names in result type StoredProduct
	// indexed by view name.
	StoredProductMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"currency_unit",
			"description",
			"updated_time",
			"purchase_price",
			"selling_price",
		},
		"tiny": []string{
			"id",
			"name",
			"currency_unit",
			"selling_price",
		},
	}
)

// ValidateStoredProductCollection runs the validations defined on the viewed
// result type StoredProductCollection.
func ValidateStoredProductCollection(result StoredProductCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredProductCollectionView(result.Projected)
	case "tiny":
		err = ValidateStoredProductCollectionViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStoredProduct runs the validations defined on the viewed result type
// StoredProduct.
func ValidateStoredProduct(result *StoredProduct) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredProductView(result.Projected)
	case "tiny":
		err = ValidateStoredProductViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStoredProductCollectionView runs the validations defined on
// StoredProductCollectionView using the "default" view.
func ValidateStoredProductCollectionView(result StoredProductCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredProductView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredProductCollectionViewTiny runs the validations defined on
// StoredProductCollectionView using the "tiny" view.
func ValidateStoredProductCollectionViewTiny(result StoredProductCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredProductViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredProductView runs the validations defined on StoredProductView
// using the "default" view.
func ValidateStoredProductView(result *StoredProductView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.CurrencyUnit == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("currency_unit", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.UpdatedTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_time", "result"))
	}
	if result.PurchasePrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("purchase_price", "result"))
	}
	if result.SellingPrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("selling_price", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	return
}

// ValidateStoredProductViewTiny runs the validations defined on
// StoredProductView using the "tiny" view.
func ValidateStoredProductViewTiny(result *StoredProductView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.CurrencyUnit == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("currency_unit", "result"))
	}
	if result.SellingPrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("selling_price", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	return
}
