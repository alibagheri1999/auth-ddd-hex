package postgres

import (
	"DDD-HEX/internal/domain"
	"database/sql"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) Save(product domain.Product) error {
	stmt, err := r.DB.Prepare("INSERT INTO products (id, name, description, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Description, product.Price)
	return err
}

func (r *ProductRepository) FindByID(id string) (domain.Product, error) {
	var product domain.Product
	query := "SELECT id, name, description, price FROM products WHERE id = $1"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return product, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	return product, err
}

func (r *ProductRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	query := "SELECT id, name, description, price FROM products"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return products, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
