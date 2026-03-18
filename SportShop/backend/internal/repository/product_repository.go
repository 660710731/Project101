package repository

import (
	"database/sql"
	"sportshop/backend/internal/model"
)

type ProductRepository struct{
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
func (r *ProductRepository) AllProducts() ([]model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, stock, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	var product model.Product
	
	query := `SELECT id, name, price FROM products WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No product found with the given ID
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	query := `INSERT INTO products (name, description, price, stock, category_id, image_url) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ImageURL)
	return err
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, stock = $4, category_id = $5, image_url = $6 WHERE id = $7`
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ImageURL, product.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

