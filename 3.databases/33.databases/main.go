package main

import "github.com/google/uuid"
import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "fmt"

type Product struct {
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	product := NewProduct("Product 1", 9.99)

	fmt.Println(product)

	err = createProduct(db, product)

	if err != nil {
		panic(err)
	}
}

func createProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}