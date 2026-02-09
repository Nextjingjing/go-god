# บทที่ 9 GORM

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO
- ORM คืออะไร?
- Database
- Transaction

## GORM คืออะไร?
GORM เป็น package ที่ช่วยทำ ORM ไว้คุยกับ Database

## การติดตั้ง
```bash
go mod init ...
go get -u gorm.io/gorm
```

## การเตรียมการ Database
ผมได้เตรียม `docker-compose.yml` ที่จะสร้าง container `PostgreSQl` และ `pgadmin`
```bash
docker compose up -d  
```
โดยผมได้ตั้งรหัสผ่านดังนี้
```
POSTGRES_USER: user123
POSTGRES_PASSWORD: password123
POSTGRES_DB: my_database
```
และ pgadmin คือ
```
PGADMIN_DEFAULT_EMAIL: admin@example.com
PGADMIN_DEFAULT_PASSWORD: adminpassword
```
และสุดท้ายผมไม่ได้ทำ volume ไว้นะเดี๋ยวจะบอกว่าทำไม?

## การลง Driver Postgres
[Docment](https://gorm.io/docs/connecting_to_the_database.html) ได้รวบรวม Driver ต่างๆ เอาไว้แล้ว ให้คุณเลือก Database ที่คุณใช้ แต่ในตัวอย่างนี้ผมเลือกใช้เป็น Postgres

```bash
go get gorm.io/driver/postgres
```

## การใช้งาน GORM

### 1. การเชื่อมต่อกับ Database
```go
package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=user123 password=password123 dbname=my_database port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
```

`user=user123 password=password123 dbname=my_database port=5432 ` มาจากการตั้งค่าใน `docker-compose.yml` นะจ๊ะ

### 2. การสร้าง Model ที่ไว้คุยกับ Database
```go
type Product struct {
	gorm.Model // adds fields ID, CreatedAt, UpdatedAt, DeletedAt
	Code       string
	Price      uint
}
}
```
- [เอกสาร](https://gorm.io/docs/models.html) เอกสารนี้ช่วยในการทำ Model

### 3. การ Migrate
```go
package main

import (
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
}

```
- ใส่หลังเชื่อมต่อ Database ด้วยโค้ดนี้
```go
// Migrate the schema
// ใช้ Model ของข้อ 2
err = db.AutoMigrate(&Product{})
if err != nil {
	panic("failed to migrate database")
}
```

### ปล. ขอเตือนไว้ก่อนไฟล์ `main.go` ที่ผมได้สร้างมานั้น มีข้อควรระวัง!
- รันได้ครั้งเดียว
- ถ้าจะรันครั้งที่สองต้องทำดังนี้
```bash
docker compose down
docker compose up -d
```

- ลำบากหน่อย แต่ code ที่ผมเขียนมาต้องการให้เข้าใจง่ายๆ ครับ เลยลำบากหน่อย

### 4. การใช้ Context เพื่อกำหนด Lifecycle ของ Database Operation

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db = db.WithContext(ctx)
```

- จากตัวอย่างนี้ คือ สร้าง `Deadline` ให้ Database Operation

### 5. การ Create
```go
// ใช้ Model ของข้อ 2
db.Create(&Product{Code: "D42", Price: 100})
```

### 6. การ Read
```go
// ใช้ Model ของข้อ 2
var product Product
db.First(&product, 1)                 // find product with integer primary key
db.First(&product, "code = ?", "D42") // find product with code D42
```

### 7. การ Update
```go
// ใช้ Model ของข้อ 2
// Update - update product's price to 200
db.Model(&product).Update("Price", 200)
// Update - update multiple fields
db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
```

### 8. การ Delete
```go
// ใช้ Model ของข้อ 2
// Delete - delete product
db.Delete(&product, 1)
```

### 9. Transaction
- โดยปกติแล้วทุกๆ Database Operation จะเป็น Transactin โดย Default ของ GORM อยู่แล้ว แต่ไม่ได้ความว่าจะรวมเป็น Transaction เดียวกัน

```
Database Operation()
Error
Database Operation()
```
แบบนี้จะไม่ถูกมองเป็น Transaction เดียวกันแล้วไม่ Atomicity เลย

- การรวม Transaction ให้เป็นหนึ่งเดียวกันสามารถทำได้โดย

```go
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
```

### 10. เอกสารที่มีประโยชน์
- [การเขียน RAW SQL](https://gorm.io/docs/sql_builder.html)
- ความสัมพันธ์ของ Model
  - [Belong to](https://gorm.io/docs/belongs_to.html)
  - [Has One](https://gorm.io/docs/has_one.html)
  - [Has Many](https://gorm.io/docs/has_many.html)
  - [Many To Many](https://gorm.io/docs/many_to_many.html)
- [การแปลง Serializer](https://gorm.io/docs/serializer.html)