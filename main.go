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
	db.Create(&Product{Code: "D43", Price: 400})

	// read
	var product Product
	// db.First(&product, 1) // find product with primary key
	db.First(&product, "code = ?", "D43") // find product with code 
	//result := db.Find(&product)
	//fmt.Println(result)
	fmt.Println(product.Code)


	// update
	// db.Model(&product).Update("Price", 200)
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//db.Delete(&product, 1)
}
