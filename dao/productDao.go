package dao

import (
	"awesomeProject/model"
	"database/sql"
	"log"
)

const insertStatement = `
	insert into products (name) values ($1) returning id;`
const updateStatement = `
	update products set name = $1 where id = $2;`
const deleteStatementById = `
	delete from products where id = $1;`
const deleteStatementByName = `
	delete from products where lower(name) like lower($1);`
const selectStatementById = `
	select * from products where id = $1;`
const selectAllStatement = `
	select * from products`
const selectStatementByName = `
	select * from products where lower(name) like lower($1)`

var db *sql.DB

func SetDB(d *sql.DB) {
	db = d
}

func CreateProduct(product model.Product) int {
	id := 0
	err := db.QueryRow(insertStatement, product.Name).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func UpdateProduct(product model.Product) bool {
	row := db.QueryRow(updateStatement, product.Name, product.Id)
	var p model.Product
	_ = row.Scan(&p.Id, &p.Name)
	return product.Name != p.Name
}

func DeleteById(id int) bool {
	result, _ := db.Exec(deleteStatementById, id)
	count, _ := result.RowsAffected()
	if count == 0 {
		return false
	}
	return true
}

func DeleteByName(name string) bool {
	result, _ := db.Exec(deleteStatementByName, name)
	count, _ := result.RowsAffected()
	if count == 0 {
		return false
	}
	return true
}

func SelectById(id int) model.Product {
	var p model.Product
	row := db.QueryRow(selectStatementById, id)
	_ = row.Scan(&p.Id, &p.Name)

	return p
}

func SelectByName(name string) model.Product {
	var p model.Product
	row := db.QueryRow(selectStatementByName, name)
	_ = row.Scan(&p.Id, &p.Name)
	return p
}

func SelectAll() []model.Product {
	var products []model.Product
	var p model.Product
	rows, err := db.Query(selectAllStatement)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		_ = rows.Scan(&p.Id, &p.Name)
		products = append(products, p)
	}
	return products
}
