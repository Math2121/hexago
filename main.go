package main

import (
	"database/sql"

	"github.com/codeedu/go-hexagonal/adapters/database"
	"github.com/codeedu/go-hexagonal/application"
)

func main() {
	Db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdpter := database.NewProductDb(Db)
	productService := application.NewProductService(productDbAdpter)

	product, _ := productService.Create("product",30)

	productService.Enable(product)

}