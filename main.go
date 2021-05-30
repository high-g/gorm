package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate schema
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{Code: "D42", Price: 100})

	// read
	var product Product
	// db.First(&product, 1) // find product with primary key
	db.First(&product, "code = ?", "D42") // find product with code 
	fmt.Println(product)


	// update
	// db.Model(&product).Update("Price", 200)
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//db.Delete(&product, 1)
}
