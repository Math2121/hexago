package database_test

import (
	"database/sql"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/database"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp(){
	Db, _ = sql.Open("sqlite3", ":memory:")
	CreateTable(Db)
	createProduct(Db)
}

func CreateTable(db *sql.DB){
	table := "create table products(id varchar(255), name varchar(255), price float, status string);"
	stmt, err := db.Prepare(table)

	if err!= nil {
        panic(err)
    }
	stmt.Exec()
}

func createProduct(db *sql.DB){
	insert := `INSERT INTO products (id, name, price, status) VALUES (
		'1',
        'teste',
        100,
        'enabled'
    );`

	stmt, err := db.Prepare(insert)
	if err!= nil {
        panic(err)
    }
	stmt.Exec()
}

func TestProduct_Get(t *testing.T){
	setUp()
	defer Db.Close()
	productDb := database.NewProductDb(Db)
	product, err := productDb.Get("1")

	require.Nil(t,err)

	require.Equal(t, "teste", product.GetName())

}

func TestProduct_Save(t *testing.T){
	setUp()
    defer Db.Close()
    productDb := database.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "teste"
	product.Price = 100

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, "teste", productResult.GetName())

}

func TestProduct_Update(t *testing.T){
	setUp()
    defer Db.Close()
    productDb := database.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "teste"
	product.Price = 100

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, "teste", productResult.GetName())
	
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Name = "teste2"
	product.Price = 200

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "teste2", productResult.GetName())

}