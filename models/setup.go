package models

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// "gorm.io/driver/postgres"
	// _ "gorm.io/driver/postgres"
	// "gorm.io/gorm"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// var DB *gorm.DB
// var DB *sql.DB

// func ConnectDB() {
// 	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=klinik sslmode=disable password=postgres")
// 	// db, err := gorm.Open(postgres.New(postgres.Config{
// 	// 	DSN:                  "host=localhost user=postgres password=postgres dbname=klinik port=5432 sslmode=disable",
// 	// 	PreferSimpleProtocol: true,
// 	// }), &gorm.Config{})

// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "postgres", "klinik")
// 	db, err := sql.Open("postgres", dbinfo)

// 	if err != nil {
// 		panic("Failed to connect to DB")
// 	}

// 	// db.AutoMigrate(&Nota{}, &Patient{}, &Visit{})

// 	DB = db
// }

// SetupDB open connection to DB
// and return instance of DB
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "postgres", "klinik")
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic("Failed to connect to DB")
	}

	return db
}
