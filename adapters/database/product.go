package database

import (
	"database/sql"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}
func (pd *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := pd.db.Prepare(" SELECT id, name, price,status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (pd *ProductDB) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := pd.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err!= nil {
        return nil, err
    }
	return product, nil
}

func (pd *ProductDB) update(product application.ProductInterface) (application.ProductInterface, error){
	stmt, err := pd.db.Prepare("UPDATE products SET name =?, price =?, status =? WHERE id =?")
    if err!= nil {
        return nil, err
    }
    _, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

    if err!= nil {
        return nil, err
    }
    err = stmt.Close()
    if err!= nil {
        return nil, err
    }
    return product, nil
}

func (pd *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	pd.db.QueryRow("SELECT COUNT(*) FROM products WHERE id =?", product.GetID()).Scan(&rows)

	if rows == 0 {
       _, err := pd.create(product)
	   if err!= nil {
           return nil, err
       }
    }else{
		_, err := pd.update(product)
		if err!= nil {
            return nil, err
        }

	}
	return product, nil
}