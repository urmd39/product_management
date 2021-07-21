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
	sql := `insert into product (name, currency_unit_id, description, updated_time, 
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

	pd_old := GetProductById(id)
	if pd_old.Purchase_price != pd.Purchase_price || pd_old.Selling_price != pd.Selling_price {
		Update_History(id)
	}
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

func FilterProductWithCU(f string) []entities.Product {
	db := Connected()
	defer db.Close()
	db.Exec(`set search_path='product'`)

	listProduct := make([]entities.Product, 0)
	sql := `SELECT p.* FROM product p, currency_unit cu
			WHERE p.currency_unit_id = cu.id
			AND cu.unit = $1`
	rows, err := db.Query(sql, f)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var pd entities.Product
		query_err := rows.Scan(&pd.Id, &pd.Name, &pd.Currency_unit_id, &pd.Description,
			&pd.Updated_time, &pd.Purchase_price, &pd.Selling_price)
		if query_err != nil {
			panic(query_err)
		}
		listProduct = append(listProduct, pd)
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
	rows, err := db.Query(sql, id, pd.Updated_time, pd.Purchase_price, pd.Selling_price)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
}
