// Code generated by goa v3.4.3, DO NOT EDIT.
//
// product HTTP client CLI support package
//
// Command:
// $ goa gen goa-demo/design

package cli

import (
	"flag"
	"fmt"
	productc "goa-demo/gen/http/product/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `product (get-list-product|get-product|filter-product|create-product|update-product|remove-product)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` product get-list-product` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		productFlags = flag.NewFlagSet("product", flag.ContinueOnError)

		productGetListProductFlags = flag.NewFlagSet("get-list-product", flag.ExitOnError)

		productGetProductFlags    = flag.NewFlagSet("get-product", flag.ExitOnError)
		productGetProductIDFlag   = productGetProductFlags.String("id", "REQUIRED", "ID of product to show")
		productGetProductViewFlag = productGetProductFlags.String("view", "", "")

		productFilterProductFlags            = flag.NewFlagSet("filter-product", flag.ExitOnError)
		productFilterProductCurrencyUnitFlag = productFilterProductFlags.String("currency-unit", "REQUIRED", "")

		productCreateProductFlags    = flag.NewFlagSet("create-product", flag.ExitOnError)
		productCreateProductBodyFlag = productCreateProductFlags.String("body", "REQUIRED", "")

		productUpdateProductFlags    = flag.NewFlagSet("update-product", flag.ExitOnError)
		productUpdateProductBodyFlag = productUpdateProductFlags.String("body", "REQUIRED", "")
		productUpdateProductIDFlag   = productUpdateProductFlags.String("id", "REQUIRED", "ID of Product need update")

		productRemoveProductFlags  = flag.NewFlagSet("remove-product", flag.ExitOnError)
		productRemoveProductIDFlag = productRemoveProductFlags.String("id", "REQUIRED", "ID of product to remove")
	)
	productFlags.Usage = productUsage
	productGetListProductFlags.Usage = productGetListProductUsage
	productGetProductFlags.Usage = productGetProductUsage
	productFilterProductFlags.Usage = productFilterProductUsage
	productCreateProductFlags.Usage = productCreateProductUsage
	productUpdateProductFlags.Usage = productUpdateProductUsage
	productRemoveProductFlags.Usage = productRemoveProductUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "product":
			svcf = productFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "product":
			switch epn {
			case "get-list-product":
				epf = productGetListProductFlags

			case "get-product":
				epf = productGetProductFlags

			case "filter-product":
				epf = productFilterProductFlags

			case "create-product":
				epf = productCreateProductFlags

			case "update-product":
				epf = productUpdateProductFlags

			case "remove-product":
				epf = productRemoveProductFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "product":
			c := productc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-list-product":
				endpoint = c.GetListProduct()
				data = nil
			case "get-product":
				endpoint = c.GetProduct()
				data, err = productc.BuildGetProductPayload(*productGetProductIDFlag, *productGetProductViewFlag)
			case "filter-product":
				endpoint = c.FilterProduct()
				data, err = productc.BuildFilterProductPayload(*productFilterProductCurrencyUnitFlag)
			case "create-product":
				endpoint = c.CreateProduct()
				data, err = productc.BuildCreateProductPayload(*productCreateProductBodyFlag)
			case "update-product":
				endpoint = c.UpdateProduct()
				data, err = productc.BuildUpdateProductPayload(*productUpdateProductBodyFlag, *productUpdateProductIDFlag)
			case "remove-product":
				endpoint = c.RemoveProduct()
				data, err = productc.BuildRemoveProductPayload(*productRemoveProductIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// productUsage displays the usage of the product command and its subcommands.
func productUsage() {
	fmt.Fprintf(os.Stderr, `The product management service.
Usage:
    %s [globalflags] product COMMAND [flags]

COMMAND:
    get-list-product: GetListProduct implements getListProduct.
    get-product: GetProduct implements getProduct.
    filter-product: FilterProduct implements filterProduct.
    create-product: CreateProduct implements createProduct.
    update-product: UpdateProduct implements updateProduct.
    remove-product: Remove product from storage

Additional help:
    %s product COMMAND --help
`, os.Args[0], os.Args[0])
}
func productGetListProductUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] product get-list-product

GetListProduct implements getListProduct.

Example:
    `+os.Args[0]+` product get-list-product
`, os.Args[0])
}

func productGetProductUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] product get-product -id INT -view STRING

GetProduct implements getProduct.
    -id INT: ID of product to show
    -view STRING: 

Example:
    `+os.Args[0]+` product get-product --id 1105576211304954661 --view "tiny"
`, os.Args[0])
}

func productFilterProductUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] product filter-product -currency-unit STRING

FilterProduct implements filterProduct.
    -currency-unit STRING: 

Example:
    `+os.Args[0]+` product filter-product --currency-unit "Laborum sint fuga eveniet."
`, os.Args[0])
}

func productCreateProductUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] product create-product -body JSON

CreateProduct implements createProduct.
    -body JSON: 

Example:
    `+os.Args[0]+` product create-product --body '{
      "currency_unit": "Minima debitis et et sed rerum.",
      "description": "Voluptas rerum tempore reiciendis.",
      "name": "Blue\'s Cuvee",
      "purchase_price": 0.869811431951743,
      "selling_price": 0.08936100869612645,
      "updated_time": "Quis et."
   }'
`, os.Args[0])
}

func productUpdateProductUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] product update-product -body JSON -id INT

UpdateProduct implements updateProduct.
    -body JSON: 
    -id INT: ID of Product need update

Example:
    `+os.Args[0]+` product update-product --body '{
      "product": {
         "currency_unit": "Possimus earum.",
         "description": "Ut repudiandae.",
         "name": "Blue\'s Cuvee",
         "purchase_price": 0.9193737983768083,
         "selling_price": 0.8555095810048335,
         "updated_time": "In non nihil provident qui vel iste."
      }
   }' --id 6591465912836711761
`, os.Args[0])
}

func productRemoveProductUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] product remove-product -id INT

Remove product from storage
    -id INT: ID of product to remove

Example:
    `+os.Args[0]+` product remove-product --id 3507015756099224420
`, os.Args[0])
}
