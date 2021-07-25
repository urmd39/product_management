package control

import (
	"fmt"
	"goa-demo/gen/product"
	"log"
)

func GetListProduct() product.StoredProductCollection {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	fmt.Println("Schemas path: Product")
	sql := `SELECT p.id, p.name, cu.unit, p.description, p.updated_time, p.purchase_price, 
		p.selling_price
		FROM product p, currency_unit cu
		WHERE cu.id = p.currency_unit_id`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	listProduct := make([]*product.StoredProduct, 0)
	for rows.Next() {
		var pd product.StoredProduct
		query_err := rows.Scan(&pd.ID, &pd.Name, &pd.CurrencyUnit, &pd.Description,
			&pd.UpdatedTime, &pd.PurchasePrice, &pd.SellingPrice)
		if query_err != nil {
			panic(query_err)
		}
		listProduct = append(listProduct, &pd)
	}
	fmt.Println(listProduct)
	return listProduct
}

func AddProduct(pd *product.Product) (str string) {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)
	fmt.Println("Add Product")
	fmt.Println(pd)
	check := false
	sql := `INSERT INTO product (name, currency_unit_id, description, updated_time, purchase_price,
		selling_price) VALUES ($1, (SELECT cu.id FROM currency_unit cu
			WHERE cu.unit = $2), $3, $4, $5, $6)`
	rows, err := db.Query(sql, pd.Name, pd.CurrencyUnit, pd.Description, pd.UpdatedTime,
		pd.PurchasePrice, pd.SellingPrice)
	if err != nil {
		log.Println(err)
	} else {
		check = true
	}
	defer rows.Close()
	if check {
		str = "Product Created"
	} else {
		str = err.Error()
	}
	return str
}

func GetProductById(id int) *product.StoredProduct {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	sql := `SELECT p.id, p.name, cu.unit, p.description, p.updated_time, p.purchase_price, 
	p.selling_price
	FROM product p, currency_unit cu
	WHERE cu.id = p.currency_unit_id
	AND p.id = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var pd product.StoredProduct
	if rows.Next() {
		query_err := rows.Scan(&pd.ID, &pd.Name, &pd.CurrencyUnit, &pd.Description,
			&pd.UpdatedTime, &pd.PurchasePrice, &pd.SellingPrice)
		if query_err != nil {
			panic(query_err)
		}
	} else {
		pd.ID = id
		pd.Name = "Error Not Found"
	}
	return &pd
}

func Update_Product(id int, pd product.Product) *product.StoredProduct {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	pd_old := GetProductById(id)
	fmt.Println(pd_old)
	if pd_old.PurchasePrice != pd.PurchasePrice || pd_old.SellingPrice != pd.SellingPrice {
		Update_History(id)
	}
	sql := `UPDATE product
			SET name = $1, currency_unit_id = (SELECT cu.id
				FROM currency_unit cu
				WHERE cu.unit = $2), description = $3, updated_time = $4,
			purchase_price = $5, selling_price = $6
			WHERE id = $7`
	rows, err := db.Query(sql, pd.Name, pd.CurrencyUnit, pd.Description, pd.UpdatedTime,
		pd.PurchasePrice, pd.SellingPrice, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return GetProductById(id)

}

func Delete_Product(id int) {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	sql := `DELETE FROM product
			WHERE id = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func FilterProductWithCU(f string) product.StoredProductCollection {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	listProduct := make([]*product.StoredProduct, 0)
	sql := `SELECT p.id, p.name, cu.unit, p.description, p.updated_time, p.purchase_price,
			p.selling_price
			FROM product p, currency_unit cu
			WHERE cu.id = p.currency_unit_id
			AND cu.unit = $1`
	rows, err := db.Query(sql, f)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var pd product.StoredProduct
		query_err := rows.Scan(&pd.ID, &pd.Name, &pd.CurrencyUnit, &pd.Description,
			&pd.UpdatedTime, &pd.PurchasePrice, &pd.SellingPrice)
		if query_err != nil {
			panic(query_err)
		}
		listProduct = append(listProduct, &pd)
	}

	defer rows.Close()
	return listProduct
}

func Update_History(id int) {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	pd := GetProductById(id)
	sql := `INSERT INTO product_cost_hist (id_product, updated_time, purchase_price, selling_price)
			VALUES ($1, $2, $3, $4)`
	rows, err := db.Query(sql, id, pd.UpdatedTime, pd.PurchasePrice, pd.SellingPrice)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
}
