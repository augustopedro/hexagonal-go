package db

import (
	"database/sql"

	"github.com/augustopedro/hexagonal-go/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
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

func (ProductDb *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	ProductDb.db.QueryRow("SELECT id FROM products WHERE id = ?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := ProductDb.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := ProductDb.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (ProductDb *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := ProductDb.db.Prepare(
		`INSERT INTO products(id, name, price, status)
		 VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ProductDb *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := ProductDb.db.Exec(
		"UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())

	if err != nil {
		return nil, err
	}
	return product, nil
}
