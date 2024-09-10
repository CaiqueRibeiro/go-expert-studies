package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int     `gorm:"primaryKey"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	gorm.Model         // permite que o GORM adicione campos CreatedAt, UpdatedAt e DeletedAt automaticamente
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create one
	db.Create(&Product{
		Name:  "Laptop",
		Price: 3940.99,
	})

	// create in batch
	products := []Product{
		{Name: "Mouse", Price: 9.99},
		{Name: "Keyboard", Price: 49.99},
		{Name: "Monitor", Price: 399.99},
	}

	db.Create(products)

	// soft delete
	db.Delete(&Product{ID: 1}) // graças ao gorm.Model, o campo não será deletado, mas sim atualizado o valor de deleted_at
}
