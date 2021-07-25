package design

import (
	. "goa.design/goa/v3/dsl"
)

/// Services
var _ = Service("product", func() {
	Description("The product management service.")

	Method("getListProduct", func() {
		HTTP(func() {
			GET("/product")
			Response(StatusOK)
		})
		Result(CollectionOf(ProductStored))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("getProduct", func() {
		Result(ProductStored)
		Error("not_found", NotFound, "Product not found!!")
		Payload(func() {
			Field(1, "id", Int, "ID of product to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		HTTP(func() {
			GET("/product/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("filterProduct", func() {
		Result(CollectionOf(ProductStored))
		Error("not_found", NotFound, "Product not found!!")
		Payload(func() {
			Field(1, "currency_unit", String, "Currency Unit")
			Required("currency_unit")
		})
		HTTP(func() {
			GET("/product/filter")
			Param("currency_unit")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("currency_unit")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("createProduct", func() {
		Result(String)
		Payload(Product)
		HTTP(func() {
			POST("/product")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("updateProduct", func() {
		Result(ProductStored)
		Error("not_found", NotFound, "Product not found!!")
		Payload(func() {
			Field(1, "id", Int, "ID of Product need update")
			Field(2, "product", Product, "Product Updated")
			Required("id", "product")
		})
		HTTP(func() {
			PUT("/product/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("removeProduct", func() {
		Description("Remove product from storage")
		Payload(func() {
			Field(1, "id", Int, "ID of product to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Product not found")
		HTTP(func() {
			DELETE("/product/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

/////////////////////////////////////////////////////////////////////////////////////
/// Types
var Product = Type("Product", func() {
	Description("Product describes")
	Attribute("name", String, "Name of product", func() {
		MaxLength(100)
		Example("Blue's Cuvee")
		Meta("rpc:tag", "1")
	})
	Field(2, "currency_unit", String)
	Field(3, "description", String)
	Field(4, "updated_time", String)
	Field(5, "purchase_price", Float64)
	Field(6, "selling_price", Float64)
	Required("name", "currency_unit", "description", "updated_time", "purchase_price",
		"selling_price")
})

var ProductStored = ResultType("application/goa-demo.product", func() {
	Description("StoredProduct describes")
	TypeName("StoredProduct")
	Reference(Product)
	Attributes(func() {
		Attribute("id", Int, "ID is the unique id of the bottle.", func() {
			Meta("rpc:tag", "1")
		})
		Field(2, "name")
		Field(3, "currency_unit")
		Field(4, "description")
		Field(5, "updated_time")
		Field(6, "purchase_price")
		Field(7, "selling_price")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("currency_unit")
		Attribute("description")
		Attribute("updated_time")
		Attribute("purchase_price")
		Attribute("selling_price")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("currency_unit")
		Attribute("selling_price")
	})

	Required("id", "name", "currency_unit", "description", "updated_time", "purchase_price",
		"selling_price")
})

var List_StoredProduct = CollectionOf(ProductStored)

var Currency_Unit = Type("Currency_Unit", func() {
	Description("Currency Unit")
	Attribute("unit", String, "Currency Unit", func() {
		MaxLength(10)
		Example("USD")
		Meta("rpc:tag", "1")
	})
	Required("unit")
})

var Stored_Currency_Unit = ResultType("application/goa-demo.currency_unit", func() {
	Description("StoredCurrencyUnit describes")
	TypeName("StoredCurrencyUnit")
	Reference(Currency_Unit)

	Attributes(func() {
		Attribute("id", Int, "ID is the unique id of currency unit.", func() {
			Meta("rpc:tag", "1")
		})
		Field(2, "unit")
	})

	View("default", func() {
		Attribute("id")
		Attribute("unit")
	})

	View("tiny", func() {
		Attribute("unit")
	})
	Required("id", "unit")
})
