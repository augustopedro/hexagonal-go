package db

import (
	"database/sql"

	"github.com/augustopedro/hexagonal-go/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func (productDb *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := productDb.db.Prepare("SELECT id, name, price, status FROM products WHERE id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
