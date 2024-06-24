package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/codeedu/go-hexagonal/application"
)

type ProductDB struct {
	db *sql.DB
}

func (pd *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt,err := pd.db.Prepare(" SELECT id, name, price,status FROM products WHERE id = ?")
	if err!= nil {
        return nil, err
    }
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err!= nil {
        return nil, err
    }
	return &product, nil

}