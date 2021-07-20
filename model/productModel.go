package model

import (
	"fmt"
	"log"

	"github.com/urmd39/product_management/entities"
)

func GetListProduct() []entities.Product {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	sql := "SELECT * FROM product"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	listProduct := make([]entities.Product, 0)
	for rows.Next() {
		var pd entities.Product
		query_err := rows.Scan(&pd.Id, &pd.Name, &pd.Currency_unit_id, &pd.Description,
			&pd.Updated_time, &pd.Purchase_price, &pd.Selling_price)
		if query_err != nil {
			panic(query_err)
		}
		listProduct = append(listProduct, pd)
	}
	// fmt.Println(listProduct)
	return listProduct
}

func AddProduct(pd entities.Product) bool {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	check := false
	sql := `insert into product.product (name1, currency_unit_id, description, updated_time, 
		purchase_price, selling_price) values ($1, $2, $3, $4, $5, $6)`
	rows, err := db.Query(sql, pd.Name, pd.Currency_unit_id, pd.Description, pd.Updated_time,
		pd.Purchase_price, pd.Selling_price)
	if err != nil {
		log.Println(err)
		return check
	}
	check = true
	defer rows.Close()
	return check
}

func GetProductById(id int) entities.Product {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	sql := `SELECT * FROM product WHERE id = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var pd entities.Product
	for rows.Next() {
		pd = entities.Product{}
		query_err := rows.Scan(&pd.Id, &pd.Name, &pd.Currency_unit_id, &pd.Description,
			&pd.Updated_time, &pd.Purchase_price, &pd.Selling_price)
		if query_err != nil {
			panic(query_err)
		}
	}
	// fmt.Println(pd)
	return pd
}

func Update_Product(id int, pd entities.Product) {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	// pd_old := GetProductById(id)
	sql := `UPDATE product
			SET name = $1, currency_unit_id = $2, description = $3, updated_time = $4, 
			purchase_price = $5, selling_price = $6
			WHERE id = $7`
	rows, err := db.Query(sql, pd.Name, pd.Currency_unit_id, pd.Description, pd.Updated_time,
		pd.Purchase_price, pd.Selling_price, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println(pd)
}

func Delete_Product(id int) {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	// pd_old := GetProductById(id)
	sql := `DELETE FROM product
			WHERE id = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("Deleted")
}
