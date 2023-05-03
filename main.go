package main

import (
	"database/sql"

	db2 "github.com/augustopedro/hexagonal-go/adapter/db"
	"github.com/augustopedro/hexagonal-go/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Produto exemplo", 30)

	productService.Enable(product)
}
