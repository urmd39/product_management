package design

import . "goa.design/goa/v3/dsl"

var _ = API("product", func() {
	Title("Product Management Service")
	Description("HTTP service for managing your product")

	Server("product", func() {
		Description("product management services")
		Services("product")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000/product")
			URI("grpc://localhost:8080/product")
		})
		Host("goa.design", func() {
			Description("public host")
			URI("https://goa.design/product")
			URI("grpcs://goa.design/product")
		})
	})
})

var NotFound = Type("NotFound", func() {
	Description(`NotFound is the type returned when attempting to show or delete product that does not exist.`)
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("Product 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", Int, "ID of missing product")
	Required("message", "id")
})
