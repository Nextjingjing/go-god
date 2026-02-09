package main

import (
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model // adds fields ID, CreatedAt, UpdatedAt, DeletedAt
	Code       string
	Price      uint
}

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=user123 password=password123 dbname=my_database port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to migrate database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db = db.WithContext(ctx)

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

	// transaction example
	db.Transaction(func(tx *gorm.DB) error {
		// Operation 1
		err := tx.Create(&Product{Code: "T1000", Price: 500})
		if err.Error != nil {
			return err.Error // rollback the transaction
		}

		// Operation 2
		err = tx.Model(&Product{}).Where("code = ?", "T1000").Update("Price", 600)
		if err.Error != nil {
			return err.Error // rollback the transaction
		}

		return nil // commit the transaction
	})
}
